package entities

import (
	"gorm.io/gorm"
	"poc-push/config"
)

type Subscription struct {
	gorm.Model
	Endpoint       string `json:"endpoint" binding:"required" gorm:"unique"`
	ExpirationTime int    `json:"expirationTime"`
	Auth           string `json:"auth"`
	P256dh         string `json:"p256dh"`
}

func CreateSubscription(subscription Subscription) (Subscription, error) {
	result := config.DB.Create(&subscription)
	if result.Error != nil {
		return Subscription{}, result.Error
	}

	return subscription, nil
}

func DeleteSubscriptionByEndpoint(endpoint string) (bool, error) {
	var subscription Subscription
	result := config.DB.Where("endpoint = ?", endpoint).Delete(&subscription)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func GetAllSubscriptions() ([]Subscription, error) {
	var subscriptions []Subscription
	result := config.DB.Find(&subscriptions)
	if result.Error != nil {
		return nil, result.Error
	}

	return subscriptions, nil
}
