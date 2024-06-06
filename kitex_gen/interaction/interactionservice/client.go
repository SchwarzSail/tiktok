// Code generated by Kitex v0.9.0. DO NOT EDIT.

package interactionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	interaction "tiktok/kitex_gen/interaction"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Like(ctx context.Context, req *interaction.LikeRequest, callOptions ...callopt.Option) (r *interaction.LikeResponse, err error)
	LikeList(ctx context.Context, req *interaction.LikeListRequest, callOptions ...callopt.Option) (r *interaction.LikeListResponse, err error)
	CommentPublish(ctx context.Context, req *interaction.CommentPublishRequest, callOptions ...callopt.Option) (r *interaction.CommentPublishResponse, err error)
	CommentList(ctx context.Context, req *interaction.CommentListRequest, callOptions ...callopt.Option) (r *interaction.CommentListResponse, err error)
	DeleteComment(ctx context.Context, req *interaction.DeleteCommentRequest, callOptions ...callopt.Option) (r *interaction.DeleteCommentResponse, err error)
	GetVideoInfo(ctx context.Context, req *interaction.GetVideoInfoRequest, callOptions ...callopt.Option) (r *interaction.GetVideoInfoResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kInteractionServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kInteractionServiceClient struct {
	*kClient
}

func (p *kInteractionServiceClient) Like(ctx context.Context, req *interaction.LikeRequest, callOptions ...callopt.Option) (r *interaction.LikeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Like(ctx, req)
}

func (p *kInteractionServiceClient) LikeList(ctx context.Context, req *interaction.LikeListRequest, callOptions ...callopt.Option) (r *interaction.LikeListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeList(ctx, req)
}

func (p *kInteractionServiceClient) CommentPublish(ctx context.Context, req *interaction.CommentPublishRequest, callOptions ...callopt.Option) (r *interaction.CommentPublishResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentPublish(ctx, req)
}

func (p *kInteractionServiceClient) CommentList(ctx context.Context, req *interaction.CommentListRequest, callOptions ...callopt.Option) (r *interaction.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, req)
}

func (p *kInteractionServiceClient) DeleteComment(ctx context.Context, req *interaction.DeleteCommentRequest, callOptions ...callopt.Option) (r *interaction.DeleteCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, req)
}

func (p *kInteractionServiceClient) GetVideoInfo(ctx context.Context, req *interaction.GetVideoInfoRequest, callOptions ...callopt.Option) (r *interaction.GetVideoInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoInfo(ctx, req)
}
