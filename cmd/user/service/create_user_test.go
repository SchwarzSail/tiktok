package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/user/dal/db"
	"tiktok/config"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/user"
)

func TestCreateUser(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	userService := GetUserService()
	err := userService.CreateUser(context.Background(), &user.RegisterRequest{
		Username: "smallpig",
		Password: "123456",
	})
	assert.Equal(t, errno.UserAlreadyExistErr, err)
}
