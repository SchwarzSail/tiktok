namespace go interaction

include "video.thrift"

struct BaseResp {
    1: i64 code
    2: string msg
}

struct Comment {
    1: string id
    2: string user_id
    3: string video_id
    4: string parent_id
    5: i64 child_count
    6: i64 like_count
    7: string content
    8: string created_at
    9: string updated_at
    10: string deleted_at
}

//点赞操作
struct LikeRequest {
    1: string user_id
    2: string video_id
    3: string comment_id
    4: string action_type
}

struct LikeResponse {
    1: BaseResp baseResp
}

//指定用户点赞的视频列表
struct LikeListRequest {
    1: string user_id
    2: i64 page_num
    3: i64 page_size
}

struct LikeListResponse {
    1: BaseResp baseResp
    2: list<video.Video> videos
}

//发布评论
struct CommentPublishRequest {
    1: string video_id
    2: string comment_id
    3: string content
    4: string user_id
}

struct CommentPublishResponse {
    1: BaseResp baseResp
    3: optional list<Comment> comments
}

//评论列表
struct CommentListRequest {
    1: string video_id
    2: string comment_id
    3: i64 page_num
    4: i64 page_size
}

struct CommentListResponse {
    1: BaseResp baseResp
    2: list<Comment> comments
}

//删除评论
struct DeleteCommentRequest {
    1: string video_id
    2: string comment_id
}

struct DeleteCommentResponse {
    1: BaseResp baseResp
}

//返回video的likes, comments
struct GetVideoInfoRequest {
    1: string video_id
}

struct GetVideoInfoResponse {
    1: BaseResp baseResp
    2: i64 like_count
    3: i64 comment_count
}

service InteractionService {
    LikeResponse Like(1: LikeRequest req)
    LikeListResponse LikeList(1: LikeListRequest req)
    CommentPublishResponse CommentPublish(1: CommentPublishRequest req)
    CommentListResponse CommentList(1: CommentListRequest req)
    DeleteCommentResponse DeleteComment(1: DeleteCommentRequest req)
    GetVideoInfoResponse GetVideoInfo(1: GetVideoInfoRequest req)
}