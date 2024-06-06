package pack

import (
	"tiktok/internal/errno"
	"tiktok/kitex_gen/video"
)

func BuildBaseResp(err error) *video.BaseResp {
	if err == nil {
		return &video.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.Success.ErrMsg,
		}
	}
	Errno := errno.ConvertErr(err)
	return &video.BaseResp{
		Code: Errno.ErrCode,
		Msg:  Errno.ErrMsg,
	}
}
