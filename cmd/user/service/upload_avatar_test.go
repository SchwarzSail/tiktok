package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/user/dal/db"
	"tiktok/config"
)

func TestUploadAvatar(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	userService := GetUserService()
	u, err := userService.UploadAvatar(context.Background(), []byte(""), "3")
	assert.NotNil(t, u)
	assert.NoError(t, err)
}
