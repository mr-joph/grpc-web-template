package httpserver

import (
	"core/pkg/helper/config"
	"core/pkg/httpserver/endpoints"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	mode := config.Env.Info.Mode
	gin.SetMode(mode)
}

func GetHandler() http.Handler {
	return endpoints.Router
}

func RunStandalone() {
	port := strconv.Itoa(config.Env.Webserver.Port)
	endpoints.Router.Run(":" + port)
}
