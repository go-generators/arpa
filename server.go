package main

import (
	"arpa/controller"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Port int
}

func loadConfig() (Config, error) {
	config := Config{}

	jsonFile, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(&config)
	return config, err
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	config, err := loadConfig()
	if err != nil {
		panic("config load error!")
	}

	e.GET("/sign_in", handleSignIn)
	e.POST("/verify", handleVerify)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

func handleSignIn(ctx echo.Context) error {
	request := struct{}{}
	response := struct {
		Message string `json:"message"`
	}{}
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	message, err := controller.GetRandomMessage(20)
	if err != nil {
		return err
	}
	response.Message = message
	return ctx.JSON(http.StatusOK, response)

}
func handleVerify(ctx echo.Context) error {
	request := struct {
		Address       string `json:"address"`
		SignedMessage string `json:"SignedMessage"`
	}{}
	response := struct {
		Verified bool `json:"verified"`
	}{}
	err := ctx.Bind(&request)
	if err != nil {
		return err
	}
	signedMessages := strings.Split(request.SignedMessage, "|")
	if len(signedMessages) != 2 {
		return errors.New("Invalid SignedMessage")
	}
	message := signedMessages[0]
	signature := signedMessages[1]

	verified, err := controller.Verify(request.Address, message, signature)
	if err != nil {
		return err
	}
	response.Verified = verified
	return ctx.JSON(http.StatusOK, response)
}
