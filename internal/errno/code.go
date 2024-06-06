package errno

const (
	SuccessCode    = 0
	ServiceErrCode = 10000 + iota
	ParamErrCode
	PageOutOfRangeCode
	//user
	UserAlreadyExistErrCode
	UserNotExistErrCode
	AuthorizationFailedErrCode

	//utils
	UploadErrCode

	//likes
	LikesEAlreadyExistErrCode
	LikesNotExistErrCode

	//relation
	AlreadyFollowedErrCode
	NotFollowedErrCode
)
