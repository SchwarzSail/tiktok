package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/user/dal/db"
	"tiktok/config"
)

func TestGetUserInfo(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	userService := GetUserService()
	u, err := userService.GetUserInfo(context.Background(), "3")
	assert.NotEqual(t, nil, u)
	assert.NoError(t, err)
}
