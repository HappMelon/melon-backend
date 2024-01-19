package routes

import (
	"github.com/FlareZone/melon-backend/config"
	"github.com/FlareZone/melon-backend/internal/components"
	"github.com/FlareZone/melon-backend/internal/handler"
	"github.com/FlareZone/melon-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"golang.org/x/oauth2"
	"net/http"
)

var (
	log = log15.New("m", "routes")
)

func auth(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authHandler := handler.NewAuthHandler(service.NewUser(components.DBEngine), service.NewNonce(components.DBEngine))
		// google 登录
		authGroup.GET("/google/login", func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect,
				config.GoogleOauthCfg.AuthCodeURL("state", oauth2.AccessTypeOffline))
		})

		// google 登录回调
		authGroup.GET("/google/callback", authHandler.GoogleOauthCallback)

		// eth eip712 登录
		authGroup.POST("/ethereum/signature/nonce", authHandler.EthereumEip712SignatureNonce)
		authGroup.POST("/ethereum/signature/login", authHandler.EthereumEip712Signature)

	}
}