package pack

import (
	"tiktok/internal/errno"
	"tiktok/kitex_gen/interaction"
)

func BuildBaseResp(err error) *interaction.BaseResp {
	if err == nil {
		return &interaction.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.Success.ErrMsg,
		}
	}
	Errno := errno.ConvertErr(err)
	return &interaction.BaseResp{
		Code: Errno.ErrCode,
		Msg:  Errno.ErrMsg,
	}
}
