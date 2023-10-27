package router

import (
	"github.com/gin-gonic/gin"
	"poc-push/controllers"
)

func Routes(app *gin.Engine) {
	app.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main website",
		})
	})
	app.StaticFile("/service-worker.js", "./public/service-worker.js")

	app.DELETE("/subscription", controllers.RemoveSubscription)
	app.POST("/subscription", controllers.PostSubscription)
	app.GET("/broadcast", controllers.BroadcastNotification)
}
