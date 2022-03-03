package controllers

import "github.com/morkid/paginate"

// func sendJSON(c *fiber.Ctx,status int, obj interface{}) error{

// }
var pg = paginate.New(&paginate.Config{})

type json struct {
	Data	interface{} 	`json:"data"`
	Message string 			`json:"message,omitempty"`
}