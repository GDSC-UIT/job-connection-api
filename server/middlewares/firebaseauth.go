package middlewares

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/GDSC-UIT/job-connection-api/conf"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var FirebaseAuthHandler func() func (*fiber.Ctx) error
var fireApp *firebase.App

type UserInfo struct {
	ID		string
	Name 	string
	Email	string
	Photo	string
}

func ConnectFirebase() {
	opt := option.WithCredentialsJSON([]byte(conf.Config.FirebaseServiceAccount))
	fireApp, _ = firebase.NewApp(context.Background(), nil, opt)

	FirebaseAuthHandler = func () func (*fiber.Ctx) error {
		return func (c *fiber.Ctx) error {
			// get token from header
			idToken := c.Get(fiber.HeaderAuthorization)
			if len(idToken) == 0 {
				return fiber.NewError(fiber.StatusUnauthorized,"Missing Token")
			}
			// validate token
			client, err := fireApp.Auth(context.Background())
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError)
			}
			token, err := client.VerifyIDToken(context.Background(),idToken )
			if err != nil {
				return fiber.NewError(fiber.StatusUnauthorized,err.Error())
			}
			c.Locals("info",UserInfo{
				ID: 	token.Claims["user_id"].(string),
				Name: 	token.Claims["name"].(string),
				Email: 	token.Claims["email"].(string),
				Photo: 	token.Claims["picture"].(string),
			})
			if token.Claims["type"]=="company" {
				c.Locals("type","company")
			} else {
				c.Locals("type","user")
			}
			return c.Next()
		}
	}
}

