package endpoints

import "github.com/gin-gonic/gin"

func getIndex(context *gin.Context) {
	res := gin.H{"message": "Welcome to gRPC web template proxy"}
	context.JSON(200, res)
}

func init() {
	Router.GET("/", getIndex)
}
