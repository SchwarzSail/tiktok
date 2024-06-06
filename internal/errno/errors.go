package errno

var (
	Success        = NewErrNo(SuccessCode, "Success")
	ServiceErr     = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr       = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	PageOutOfRange = NewErrNo(PageOutOfRangeCode, "分页查询超出范围")
	// user
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	UserNotExist           = NewErrNo(UserNotExistErrCode, "User doesn't exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")

	//utils
	UploadErr = NewErrNo(UploadErrCode, "Upload failed")

	//likes
	LikesAlreadyExistErr = NewErrNo(LikesEAlreadyExistErrCode, "likes already exists")
	LikesNotExistErr     = NewErrNo(LikesNotExistErrCode, "likes doesn't exists")
	//relation
	AlreadyFollowed = NewErrNo(AlreadyFollowedErrCode, "already followed")
	NotFollowed     = NewErrNo(NotFollowedErrCode, "not yet followed")
)
