package handlers

import (
	"github.com/gofiber/fiber"
)

func CreateRedisPrefix(c *fiber.Ctx) error {
	type CreateRedisPrefixRequest struct {
		ServiceName string `json:"serviceName"`
		Prefix      string `json:"prefix"`
		JsonName    string `json:"jsonName"`
	}

}
