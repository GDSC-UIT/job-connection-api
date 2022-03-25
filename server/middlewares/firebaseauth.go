package middlewares

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/GDSC-UIT/job-connection-api/conf"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var FirebaseAuthHandler func() func(*fiber.Ctx) error
var FireBaseInfoHandler func() func(*fiber.Ctx) error
var fireApp *firebase.App

type UserInfo struct {
	ID    string
	Name  string
	Email string
	Photo string
}

func ConnectFirebase() {
	opt := option.WithCredentialsJSON([]byte(conf.Config.FirebaseServiceAccount))
	fireApp, _ = firebase.NewApp(context.Background(), nil, opt)
	FireBaseInfoHandler = func() func(*fiber.Ctx) error {
		return func(c *fiber.Ctx) error {
			// get token from header
			idToken := c.Get(fiber.HeaderAuthorization)
			if len(idToken) == 0 {
				return c.Next()
			}
			// validate token
			client, err := fireApp.Auth(context.Background())
			if err != nil {
				return c.Next()
			}
			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return c.Next()
			}
			userInfo := UserInfo{
				ID: token.Claims["user_id"].(string),
				// Name:  token.Claims["name"].(*string),
				Email: token.Claims["email"].(string),
				// Photo: token.Claims["picture"].(*string),
			}

			name, ok := token.Claims["name"].(string)
			if ok {
				userInfo.Name = name
			} else {
				userInfo.Name = "user_" + userInfo.ID[:5]
			}

			picture, ok := token.Claims["picture"].(string)
			if ok {
				userInfo.Photo = picture
			} else {
				userInfo.Photo = "https://i.imgur.com/FuoDVfD.png"
			}

			c.Locals("info", userInfo)

			return c.Next()
		}
	}

	FirebaseAuthHandler = func() func(*fiber.Ctx) error {
		return func(c *fiber.Ctx) error {
			// get token from header
			idToken := c.Get(fiber.HeaderAuthorization)
			if len(idToken) == 0 {
				return fiber.NewError(fiber.StatusUnauthorized, "Missing Token")
			}
			// validate token
			client, err := fireApp.Auth(context.Background())
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError)
			}
			token, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}
			userInfo := UserInfo{
				ID: token.Claims["user_id"].(string),
				// Name:  token.Claims["name"].(*string),
				Email: token.Claims["email"].(string),
				// Photo: token.Claims["picture"].(*string),
			}

			name, ok := token.Claims["name"].(string)
			if ok {
				userInfo.Name = name

			} else {
				userInfo.Name = "user_" + userInfo.ID[:5]
			}

			picture, ok := token.Claims["picture"].(string)
			if ok {
				userInfo.Photo = picture
			} else {
				userInfo.Photo = "https://i.imgur.com/FuoDVfD.png"
			}

			c.Locals("info", userInfo)

			return c.Next()
		}
	}
}
