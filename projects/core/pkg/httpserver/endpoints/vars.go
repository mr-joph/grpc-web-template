package endpoints

import "github.com/gin-gonic/gin"

var Router = gin.Default()
var v1 = Router.Group("/rest/v1")
