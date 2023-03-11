package routing

import (
	"AecProject/internal/Utility"
	"AecProject/internal/controller"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MainRout(e *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(Utility.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	e.POST("/signin", controller.SignIn)
	e.GET("/login", controller.Login)
	loginGroup := e.Group("/login")
	loginGroup.POST("/new-wallet", controller.AddWallet, echojwt.WithConfig(config))
	loginGroup.GET("/get-balance", controller.GetBalance, echojwt.WithConfig(config))
	loginGroup.POST("/add-money", controller.AddMoney, echojwt.WithConfig(config))
	loginGroup.DELETE("/remove-user", controller.RemoveUser, echojwt.WithConfig(config))
}
