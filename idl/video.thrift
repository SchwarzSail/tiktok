namespace go video

struct BaseResp {
    1: i64 code
    2: string msg
}

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
    1: i64 timeStamp
}

struct FeedResponse {
    1: BaseResp baseResp
    2: list<Video> videos
}

struct PublishRequest {
    1: binary data
    2: binary cover
    3: string title
    4: string description
    5: string uid
    6: string user_name
}

struct PublishResponse {
    1: BaseResp baseResp
}

struct PublishListRequest {
    1: string uid
    2: i64 page_num
    3: i64 page_size
}

struct PublishListResponse {
    1: BaseResp baseResp
    2: list<Video> videos
}

struct PopularListRequest {
    1: i64 page_num
    2: i64 page_size
}

struct PopularListResponse {
    1: BaseResp baseResp
    2: list<Video> videos
}

struct SearchRequest {
    1: string keyword
    2: i64 page_num
    3: i64 page_size
    4: optional i64 from_date
    5: optional i64 to_date
    6: optional string username
}

struct SearchResponse {
    1: BaseResp baseResp
    2: list<Video> videos
}

struct WatchVideoRequest {
    1: string video_id
}

struct WatchVideoResponse {
    1: BaseResp baseResp
    2: Video video
}

struct GetVideoInfoRequest {
    1: string video_id
}

struct GetVideoInfoResponse {
    1: BaseResp baseResp
    2: Video video
}

service VideoService {
    GetVideoInfoResponse GetVideoInfo(1: GetVideoInfoRequest req)
    FeedResponse Feed(1: FeedRequest req)
    PublishResponse Publish(1: PublishRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    PopularListResponse PopularList(1: PopularListRequest req)
    SearchResponse Search(1: SearchRequest req)
    WatchVideoResponse WatchVideo(1: WatchVideoRequest req)
}