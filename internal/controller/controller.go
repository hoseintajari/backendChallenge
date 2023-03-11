package controller

import (
	"AecProject/internal/Utility"
	"AecProject/internal/controller/apimodels"
	"AecProject/internal/models"
	"AecProject/internal/query"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var user models.User
var wallet models.Wallet
var transaction models.Transaction

func SignIn(c echo.Context) error {
	request := new(apimodels.SignInRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if query.FindUsername(request.UserName) != false {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "There is a username with the same name please Choose another name"})
	} else {
		user.Add(*request)
		return c.String(http.StatusOK, "Account created")
	}

}

func Login(c echo.Context) error {
	request := new(apimodels.LoginRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if query.CheckPassword(request.UserName, request.Password) != true {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "username or password is incorrect"})
	} else {
		role := query.GetRoles(request.UserName)
		stringToken, err := Utility.CreateJwt(request.UserName, role)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create token")
		}
		cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(1 * time.Hour)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, map[string]string{"Token": stringToken})
	}

}

func AddWallet(c echo.Context) error {

	request := new(apimodels.NewWalletRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error ": "Failed reading the request body"})
	}

	if query.FindUserId(request.UserID) != false {
		return c.JSON(http.StatusUnauthorized, "error :There is a userID with the same name")
	} else {
		wallet.Add(*request)
		return c.String(http.StatusOK, "wallet created")
	}

}

func GetBalance(c echo.Context) error {
	request := new(apimodels.GetBalanceRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if query.FindUserId(request.UserId) != true {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Wallet not found"})
	} else {
		res := wallet.GetBalance(*request)
		return c.JSON(http.StatusOK, apimodels.GetBalanceResponse{Balance: res})
	}

}

func AddMoney(c echo.Context) error {
	request := new(apimodels.AddMoneyRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if query.FindUserId(request.UserID) != true {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User ID not found. Please create your wallet first"})
	} else {
		wallet.AddMoney(*request)
		response := transaction.Add(*request)

		return c.JSON(http.StatusOK, apimodels.AddMoneyResponse{ReferenceID: response})
	}

}

func RemoveUser(c echo.Context) error {
	request := new(apimodels.RemoveUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*Utility.JwtCustomClaims)
	if !claims.Admin {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "you dont have access!"})
	}

	if query.FindUsername(request.UserName) != true {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "The user was not found"})
	} else {
		user.Remove(*request)
		return c.JSON(http.StatusOK, "The user was deleted")
	}
}
