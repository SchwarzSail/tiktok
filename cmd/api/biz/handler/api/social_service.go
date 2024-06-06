// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/api/biz/pack"
	"tiktok/cmd/api/biz/rpc"
	"tiktok/internal/errno"
	"tiktok/internal/utils"
	"tiktok/kitex_gen/social"

	"github.com/cloudwego/hertz/pkg/app"
	api "tiktok/cmd/api/biz/model/api"
)

// Follow .
// @router tiktok/relation/action [POST]
func Follow(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowRequest
	//先对token进行拦截
	userInfo, err := pack.GetUserInfo(ctx)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}
	err = rpc.Follow(ctx, &social.FollowRequest{
		UserId:     userInfo.ID,
		ToUserId:   req.UserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	pack.RespSuccess(c)
}

// FollowList .
// @router tiktok/following/list [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FollowListRequest
	//先对token进行拦截
	_, err = pack.GetUserInfo(ctx)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}
	if req.PageSize == nil {
		req.PageSize = new(int64)
		*req.PageSize = 10
	}
	if req.PageNum == nil {
		req.PageNum = new(int64)
		*req.PageNum = 1
	}
	resp := new(api.FollowListResponse)
	list, err := rpc.FollowList(ctx, &social.FollowListRequest{
		UserId:   req.UserID,
		PageNum:  *req.PageNum,
		PageSize: *req.PageSize,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	resp.Users = pack.BuildUserList(list)
	pack.RespList(c, resp, int64(len(list)))
}

// FansList .
// @router tiktok/follower/list [GET]
func FansList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FansListRequest
	//先对token进行拦截
	_, err = pack.GetUserInfo(ctx)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}
	if req.PageSize == nil {
		req.PageSize = new(int64)
		*req.PageSize = 10
	}
	if req.PageNum == nil {
		req.PageNum = new(int64)
		*req.PageNum = 1
	}
	resp := new(api.FansListResponse)
	list, err := rpc.FansList(ctx, &social.FansListRequest{
		Uid:      req.UserID,
		PageNum:  *req.PageNum,
		PageSize: *req.PageSize,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	resp.Users = pack.BuildUserList(list)
	pack.RespList(c, resp, int64(len(list)))
}

// FriendsList .
// @router tiktok/friends/list [GET]
func FriendsList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FriendsListRequest
	//先对token进行拦截
	userInfo, err := pack.GetUserInfo(ctx)
	if err != nil {
		pack.RespError(c, err)
		return
	}
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}
	if req.PageSize == nil {
		req.PageSize = new(int64)
		*req.PageSize = 10
	}
	if req.PageNum == nil {
		req.PageNum = new(int64)
		*req.PageNum = 1
	}
	resp := new(api.FriendsListResponse)
	list, err := rpc.FriendsList(ctx, &social.FriendsListRequest{
		Uid:      userInfo.ID,
		PageNum:  *req.PageNum,
		PageSize: *req.PageSize,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	resp.Users = pack.BuildUserList(list)
	pack.RespList(c, resp, int64(len(list)))
}