package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/user/dal/db"
	"tiktok/config"
	"tiktok/kitex_gen/user"
)

func TestCheckUser(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	userService := GetUserService()
	result, err := userService.CheckUser(context.Background(), &user.LoginRequest{
		Username: "smallpig",
		Password: "123456",
		Otp:      nil,
	})

	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, result, "User should not be nil")
	assert.Equal(t, "smallpig", result.Username, "Username should match")
}
