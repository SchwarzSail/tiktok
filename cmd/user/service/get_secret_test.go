package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/user/dal/db"
	"tiktok/config"
)

func TestGetSecret(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	userService := GetUserService()
	secret, err := userService.GetSecret(context.Background(), "3")
	assert.Equal(t, "", secret)
	assert.NoError(t, err)
}
