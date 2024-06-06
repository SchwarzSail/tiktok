namespace go api

//user
struct User {
    1: string id
    2: string name
    3: string avatar_url
    4: string created_at
    5: string updated_at
    6: string deleted_at
}

//注册
struct RegisterRequest {
    1: required string username
    2: required string password
}

struct RegisterResponse {
}

//登录
struct LoginRequest {
    1: required string username
    2: required string password
    3: optional string otp
}

struct LoginResponse {
    1: required User user
    2: required string refresh_token
    3: required string access_token
}

//用户信息获取
struct InfoRequest {
    1: required string id
}

struct InfoResponse {
    1: required User user
}

//头像上传
struct AvatarUploadRequest {
    //1: required binary data
}

struct AvatarUploadResponse {
    1: required User user
}

//获取MFAqrcode
struct GetMFAqrcodeRequest {
}

struct GetMFAqrcodeResponse {
    1: required string secret
    2: required string qrcode
}

//MFA认证
struct MFABindRequest {
    1: required string code
    2: required string secret
}

struct MFABindResponse {
}


//video
struct Video {
    1: string id
    2: string user_id
    3: string video_url
    4: string cover_url
    5: string title
    6: string description
    7: i64 visit_count
    8: i64 like_count
    9: i64 comment_count
    10: string created_at
    11: string updated_at
    12: string deleted_at
}

struct FeedRequest {
    1: required i64 timeStamp
}

struct FeedResponse {
    1: required list<Video> videos
}

struct PublishRequest {
    //1: required binary data
    //2: required binary cover
    3: required string title
    4: required string description
}

struct PublishResponse {

}

struct PublishListRequest {
    1: required string user_id
    2: required i64 page_num
    3: required i64 page_size
}

struct PublishListResponse {
    1: required list<Video> videos
}

struct PopularListRequest {
    1: required i64 page_num
    2: required i64 page_size
}

struct PopularListResponse {
    1: required list<Video> videos
}

struct SearchRequest {
    1: required string keyword
    2: required i64 page_num
    3: required i64 page_size
    4: optional i64 from_date
    5: optional i64 to_date
    6: optional string username
}

struct SearchResponse {
    1: required list<Video> videos
}

struct Comment {
    1: string id
    2: string user_id
    3: string video_id
    4: string parent_id
    5: i64 like_count
    6: i64 child_count
    7: string content
    8: string created_at
    9: string updated_at
    10: string deleted_at
}

//点赞操作
struct LikeRequest {
    1: optional string video_id
    2: optional string comment_id
    3: required string action_type
}

struct LikeResponse {
}

//指定用户点赞的视频列表
struct LikeListRequest {
    1: required string user_id
    2: optional i64 page_num
    3: optional i64 page_size
}

struct LikeListResponse {
    1: required list<Video> videos
}

//发布评论
struct CommentPublishRequest {
    1: optional string video_id
    2: optional string comment_id
    3: required string content
}

struct CommentPublishResponse {
}

//评论列表
struct CommentListRequest {
    1: optional string video_id
    2: optional string comment_id
    3: optional i64 page_num
    4: optional i64 page_size
}

struct CommentListResponse {
    1: optional list<Comment> comments
}

//删除评论
struct DeleteCommentRequest {
    1: optional string video_id
    2: optional string comment_id
}

struct DeleteCommentResponse {
}



struct FollowRequest {
    1: required string user_id
    2: required string action_type
}

struct FollowResponse {
}

struct FollowListRequest {
    1: required string user_id
    2: optional i64 page_num
    3: optional i64 page_size
}

struct FollowListResponse {
    1: required list<User> users
}

struct FansListRequest {
    1: required string user_id
    2: optional i64 page_num
    3: optional i64 page_size
}

struct FansListResponse {
    1:required list<User> users
}

struct FriendsListRequest {
    1: optional i64 page_num
    2: optional i64 page_size
}

struct FriendsListResponse {
    1: required list<User> users
}


//service

service UserService {
    RegisterResponse Register(1: RegisterRequest req) (api.post="tiktok/user/register")
    LoginResponse Login(1: LoginRequest req) (api.post="tiktok/user/login")
    InfoResponse GetInfo(1 :InfoRequest req) (api.get="tiktok/user/info")
    AvatarUploadResponse AvatarUpload(1: AvatarUploadRequest req) (api.put="tiktok/user/avatar/upload")
    GetMFAqrcodeResponse GetMFAqrcode(1: GetMFAqrcodeRequest req) (api.get="tiktok/auth/mfa/qrcode")
    MFABindResponse MFABind(1: MFABindRequest req) (api.post="tiktok/auth/mfa/bind")
}

service VideoService {
    FeedResponse Feed(1: FeedRequest req) (api.get="tiktok/video/feed")
    PublishResponse Publish(1: PublishRequest req) (api.post="tiktok/video/publish")
    PublishListResponse PublishList(1: PublishListRequest req) (api.get="tiktok/video/list")
    PopularListResponse PopularList(1: PopularListRequest req) (api.get="tiktok/video/popular")
    SearchResponse Search(1: SearchRequest req) (api.post="tiktok/video/search")
}

service InteractionService {
    LikeResponse Like(1: LikeRequest req) (api.post="tiktok/like/action")
    LikeListResponse LikeList(1: LikeListRequest req) (api.get="tiktok/like/list")
    CommentPublishResponse CommentPublish(1: CommentPublishRequest req) (api.post="tiktok/comment/publish")
    CommentListResponse CommentList(1: CommentListRequest req) (api.get="tiktok/comment/list")
    DeleteCommentResponse DeleteComment(1: DeleteCommentRequest req) (api.delete="tiktok/comment/delete")
}

service SocialService {
    FollowResponse Follow(1: FollowRequest req) (api.post="tiktok/relation/action")
    FollowListResponse FollowList(1: FollowListRequest req) (api.get="tiktok/following/list")
    FansListResponse FansList(1: FansListRequest req) (api.get="tiktok/follower/list")
    FriendsListResponse FriendsList(1: FriendsListRequest req) (api.get="tiktok/friends/list")
}