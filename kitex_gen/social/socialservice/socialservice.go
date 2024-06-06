// Code generated by Kitex v0.9.0. DO NOT EDIT.

package socialservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	social "tiktok/kitex_gen/social"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Follow": kitex.NewMethodInfo(
		followHandler,
		newSocialServiceFollowArgs,
		newSocialServiceFollowResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"FollowList": kitex.NewMethodInfo(
		followListHandler,
		newSocialServiceFollowListArgs,
		newSocialServiceFollowListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"FansList": kitex.NewMethodInfo(
		fansListHandler,
		newSocialServiceFansListArgs,
		newSocialServiceFansListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"FriendsList": kitex.NewMethodInfo(
		friendsListHandler,
		newSocialServiceFriendsListArgs,
		newSocialServiceFriendsListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	socialServiceServiceInfo                = NewServiceInfo()
	socialServiceServiceInfoForClient       = NewServiceInfoForClient()
	socialServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return socialServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return socialServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return socialServiceServiceInfoForClient
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
	serviceName := "SocialService"
	handlerType := (*social.SocialService)(nil)
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
		"PackageName": "social",
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

func followHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFollowArgs)
	realResult := result.(*social.SocialServiceFollowResult)
	success, err := handler.(social.SocialService).Follow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFollowArgs() interface{} {
	return social.NewSocialServiceFollowArgs()
}

func newSocialServiceFollowResult() interface{} {
	return social.NewSocialServiceFollowResult()
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFollowListArgs)
	realResult := result.(*social.SocialServiceFollowListResult)
	success, err := handler.(social.SocialService).FollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFollowListArgs() interface{} {
	return social.NewSocialServiceFollowListArgs()
}

func newSocialServiceFollowListResult() interface{} {
	return social.NewSocialServiceFollowListResult()
}

func fansListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFansListArgs)
	realResult := result.(*social.SocialServiceFansListResult)
	success, err := handler.(social.SocialService).FansList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFansListArgs() interface{} {
	return social.NewSocialServiceFansListArgs()
}

func newSocialServiceFansListResult() interface{} {
	return social.NewSocialServiceFansListResult()
}

func friendsListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFriendsListArgs)
	realResult := result.(*social.SocialServiceFriendsListResult)
	success, err := handler.(social.SocialService).FriendsList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFriendsListArgs() interface{} {
	return social.NewSocialServiceFriendsListArgs()
}

func newSocialServiceFriendsListResult() interface{} {
	return social.NewSocialServiceFriendsListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Follow(ctx context.Context, req *social.FollowRequest) (r *social.FollowResponse, err error) {
	var _args social.SocialServiceFollowArgs
	_args.Req = req
	var _result social.SocialServiceFollowResult
	if err = p.c.Call(ctx, "Follow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, req *social.FollowListRequest) (r *social.FollowListResponse, err error) {
	var _args social.SocialServiceFollowListArgs
	_args.Req = req
	var _result social.SocialServiceFollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FansList(ctx context.Context, req *social.FansListRequest) (r *social.FansListResponse, err error) {
	var _args social.SocialServiceFansListArgs
	_args.Req = req
	var _result social.SocialServiceFansListResult
	if err = p.c.Call(ctx, "FansList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendsList(ctx context.Context, req *social.FriendsListRequest) (r *social.FriendsListResponse, err error) {
	var _args social.SocialServiceFriendsListArgs
	_args.Req = req
	var _result social.SocialServiceFriendsListResult
	if err = p.c.Call(ctx, "FriendsList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
