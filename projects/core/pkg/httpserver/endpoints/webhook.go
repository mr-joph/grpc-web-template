package endpoints

import "github.com/gin-gonic/gin"

type WebhookRequest struct {
}

type WebhookResponse struct {
}

func postWebhook(context *gin.Context) {
	res := gin.H{"message": "webhook not implemented yet"}
	context.JSON(501, res)
}

func init() {
	Router.POST("/webhook", postWebhook)
}
