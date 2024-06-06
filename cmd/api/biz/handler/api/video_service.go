// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"tiktok/cmd/api/biz/pack"
	"tiktok/cmd/api/biz/rpc"
	"tiktok/internal/errno"
	"tiktok/internal/utils"
	"tiktok/kitex_gen/video"

	"github.com/cloudwego/hertz/pkg/app"
	api "tiktok/cmd/api/biz/model/api"
)

// Feed .
// @router tiktok/video/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}

	resp := new(api.FeedResponse)
	list, err := rpc.Feed(ctx, &video.FeedRequest{TimeStamp: req.TimeStamp})
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, err)
		return
	}
	resp.Videos = pack.BuildVideoList(list)
	pack.RespList(c, resp, int64(len(list)))
}

// Publish .
// @router tiktok/video/publish [POST]
func Publish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishRequest
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
	//将发过来的数据进行转换
	data, err := c.FormFile("data")
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, errno.ParamErr)
		return
	}
	videoFile, err := data.Open()
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, errno.ParamErr)
		return
	}
	videoFile.Close()

	cover, err := c.FormFile("cover")
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, errno.ParamErr)
		return
	}
	coverFile, err := cover.Open()
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, errno.ParamErr)
		return
	}
	coverFile.Close()
	err = c.BindAndValidate(&req)

	videoData, err := io.ReadAll(videoFile)
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, errno.ParamErr)
		return
	}
	coverData, err := io.ReadAll(coverFile)
	if err != nil {
		utils.LogrusObj.Error(err)
		pack.RespError(c, errno.ParamErr)
		return
	}
	err = rpc.Publish(ctx, &video.PublishRequest{
		Data:        videoData,
		Cover:       coverData,
		Title:       req.Title,
		Description: req.Description,
		Uid:         userInfo.ID,
		UserName:    userInfo.UserName,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	pack.RespSuccess(c)
}

// PublishList .
// @router tiktok/video/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}

	resp := new(api.PublishListResponse)
	res, err := rpc.PublishList(ctx, &video.PublishListRequest{
		Uid:      req.UserID,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	resp.Videos = pack.BuildVideoList(res)
	pack.RespList(c, resp, int64(len(res)))
}

// PopularList .
// @router tiktok/video/popular [GET]
func PopularList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PopularListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}
	resp := new(api.PopularListResponse)
	res, err := rpc.PopularList(ctx, &video.PopularListRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	resp.Videos = pack.BuildVideoList(res)
	pack.RespList(c, resp, int64(len(res)))
}

// Search .
// @router tiktok/video/search [POST]
func Search(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SearchRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.RespError(c, errno.ParamErr)
		return
	}
	resp := new(api.SearchResponse)
	res, err := rpc.Search(ctx, &video.SearchRequest{
		Keyword:  req.Keyword,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
		Username: req.Username,
	})
	if err != nil {
		utils.LogrusObj.Error(err)
		err := errors.Cause(err)
		pack.RespError(c, err)
		return
	}
	resp.Videos = pack.BuildVideoList(res)
	pack.RespList(c, resp, int64(len(res)))
}
