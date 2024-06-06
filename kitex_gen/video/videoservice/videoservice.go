// Code generated by Kitex v0.9.0. DO NOT EDIT.

package videoservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	video "tiktok/kitex_gen/video"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetVideoInfo": kitex.NewMethodInfo(
		getVideoInfoHandler,
		newVideoServiceGetVideoInfoArgs,
		newVideoServiceGetVideoInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Feed": kitex.NewMethodInfo(
		feedHandler,
		newVideoServiceFeedArgs,
		newVideoServiceFeedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Publish": kitex.NewMethodInfo(
		publishHandler,
		newVideoServicePublishArgs,
		newVideoServicePublishResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"PublishList": kitex.NewMethodInfo(
		publishListHandler,
		newVideoServicePublishListArgs,
		newVideoServicePublishListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"PopularList": kitex.NewMethodInfo(
		popularListHandler,
		newVideoServicePopularListArgs,
		newVideoServicePopularListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Search": kitex.NewMethodInfo(
		searchHandler,
		newVideoServiceSearchArgs,
		newVideoServiceSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"WatchVideo": kitex.NewMethodInfo(
		watchVideoHandler,
		newVideoServiceWatchVideoArgs,
		newVideoServiceWatchVideoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	videoServiceServiceInfo                = NewServiceInfo()
	videoServiceServiceInfoForClient       = NewServiceInfoForClient()
	videoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return videoServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return videoServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.0",
		Extra:           extra,
	}
	return svcInfo
}

func getVideoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetVideoInfoArgs)
	realResult := result.(*video.VideoServiceGetVideoInfoResult)
	success, err := handler.(video.VideoService).GetVideoInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetVideoInfoArgs() interface{} {
	return video.NewVideoServiceGetVideoInfoArgs()
}

func newVideoServiceGetVideoInfoResult() interface{} {
	return video.NewVideoServiceGetVideoInfoResult()
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

func publishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishArgs)
	realResult := result.(*video.VideoServicePublishResult)
	success, err := handler.(video.VideoService).Publish(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishArgs() interface{} {
	return video.NewVideoServicePublishArgs()
}

func newVideoServicePublishResult() interface{} {
	return video.NewVideoServicePublishResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishListArgs)
	realResult := result.(*video.VideoServicePublishListResult)
	success, err := handler.(video.VideoService).PublishList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishListArgs() interface{} {
	return video.NewVideoServicePublishListArgs()
}

func newVideoServicePublishListResult() interface{} {
	return video.NewVideoServicePublishListResult()
}

func popularListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePopularListArgs)
	realResult := result.(*video.VideoServicePopularListResult)
	success, err := handler.(video.VideoService).PopularList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePopularListArgs() interface{} {
	return video.NewVideoServicePopularListArgs()
}

func newVideoServicePopularListResult() interface{} {
	return video.NewVideoServicePopularListResult()
}

func searchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceSearchArgs)
	realResult := result.(*video.VideoServiceSearchResult)
	success, err := handler.(video.VideoService).Search(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceSearchArgs() interface{} {
	return video.NewVideoServiceSearchArgs()
}

func newVideoServiceSearchResult() interface{} {
	return video.NewVideoServiceSearchResult()
}

func watchVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceWatchVideoArgs)
	realResult := result.(*video.VideoServiceWatchVideoResult)
	success, err := handler.(video.VideoService).WatchVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceWatchVideoArgs() interface{} {
	return video.NewVideoServiceWatchVideoArgs()
}

func newVideoServiceWatchVideoResult() interface{} {
	return video.NewVideoServiceWatchVideoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetVideoInfo(ctx context.Context, req *video.GetVideoInfoRequest) (r *video.GetVideoInfoResponse, err error) {
	var _args video.VideoServiceGetVideoInfoArgs
	_args.Req = req
	var _result video.VideoServiceGetVideoInfoResult
	if err = p.c.Call(ctx, "GetVideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Feed(ctx context.Context, req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Publish(ctx context.Context, req *video.PublishRequest) (r *video.PublishResponse, err error) {
	var _args video.VideoServicePublishArgs
	_args.Req = req
	var _result video.VideoServicePublishResult
	if err = p.c.Call(ctx, "Publish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, req *video.PublishListRequest) (r *video.PublishListResponse, err error) {
	var _args video.VideoServicePublishListArgs
	_args.Req = req
	var _result video.VideoServicePublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PopularList(ctx context.Context, req *video.PopularListRequest) (r *video.PopularListResponse, err error) {
	var _args video.VideoServicePopularListArgs
	_args.Req = req
	var _result video.VideoServicePopularListResult
	if err = p.c.Call(ctx, "PopularList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Search(ctx context.Context, req *video.SearchRequest) (r *video.SearchResponse, err error) {
	var _args video.VideoServiceSearchArgs
	_args.Req = req
	var _result video.VideoServiceSearchResult
	if err = p.c.Call(ctx, "Search", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) WatchVideo(ctx context.Context, req *video.WatchVideoRequest) (r *video.WatchVideoResponse, err error) {
	var _args video.VideoServiceWatchVideoArgs
	_args.Req = req
	var _result video.VideoServiceWatchVideoResult
	if err = p.c.Call(ctx, "WatchVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
