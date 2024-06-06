package pack

import (
	"tiktok/internal/errno"
	"tiktok/kitex_gen/social"
)

func BuildBaseResp(err error) *social.BaseResp {
	if err == nil {
		return &social.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.Success.ErrMsg,
		}
	}
	Errno := errno.ConvertErr(err)
	return &social.BaseResp{
		Code: Errno.ErrCode,
		Msg:  Errno.ErrMsg,
	}
}
