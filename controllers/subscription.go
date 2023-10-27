package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"poc-push/entities"
)

type Keys struct {
	ID     int
	Auth   string `json:"auth"`
	P256dh string `json:"p256dh"`
}

type Subscription struct {
	Endpoint       string `json:"endpoint" binding:"required" gorm:"unique"`
	ExpirationTime int    `json:"expirationTime"`
	Keys           Keys   `json:"keys"`
}

func PostSubscription(c *gin.Context) {
	var subscription Subscription

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new subscription
	newSubscription, err := entities.CreateSubscription(entities.Subscription{
		Endpoint:       subscription.Endpoint,
		ExpirationTime: subscription.ExpirationTime,
		Auth:           subscription.Keys.Auth,
		P256dh:         subscription.Keys.P256dh,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	// Send 201 - resource created
	c.JSON(http.StatusCreated, newSubscription)
}

func RemoveSubscription(c *gin.Context) {
	endpoint := c.Query("endpoint")
	if endpoint == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	successful, err := entities.DeleteSubscriptionByEndpoint(endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if successful {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusInternalServerError)
	}
}

func BroadcastNotification(c *gin.Context) {
	notification := map[string]string{
		"title": "Hey, this is a push notification!",
		//"body":  fmt.Sprintf("this body is writen at %s", time.Now().Format("2006-01-02 15:04:05")),
		"body": "https://7567-2804-1b1-2107-ac9b-85b1-1052-ee4c-8c59.ngrok-free.app/index",
	}

	message, err := json.Marshal(notification)
	if err != nil {
		log.Fatal(err)
	}

	subscriptions, err := entities.GetAllSubscriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, subscription := range subscriptions {
		sub := webpush.Subscription{
			Endpoint: subscription.Endpoint,
			Keys: webpush.Keys{
				Auth:   subscription.Auth,
				P256dh: subscription.P256dh,
			},
		}

		// Enviando a notificação
		resp, err := webpush.SendNotification(message, &sub, &webpush.Options{
			Subscriber:      "email@email.com",
			VAPIDPublicKey:  os.Getenv("VAPID_PUB_KEY"),
			VAPIDPrivateKey: os.Getenv("VAPID_PRIVATE_KEY"),
			TTL:             30,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.StatusCode, body)
	}

	c.Status(http.StatusOK)
}
