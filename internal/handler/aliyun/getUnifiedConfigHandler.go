// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package aliyun

import (
	"github.com/gpencil/upload/internal/logic/aliyun"
	"github.com/gpencil/upload/internal/svc"
	"github.com/gpencil/upload/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ysgit.lunalabs.cn/products/go-common/errcode"
	"ysgit.lunalabs.cn/products/go-common/response"
)

// 获取统一配置
func GetUnifiedConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUnifiedConfigReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.WithContext(r.Context()).Errorf("参数错误,err:%+v", err)
			httpx.OkJsonCtx(r.Context(), w, response.ResponseError(errcode.ErrInvalidParams))
			return
		}

		l := aliyun.NewGetUnifiedConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetUnifiedConfig(&req)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("handler err:%+v", err)
			httpx.OkJsonCtx(r.Context(), w, response.ResponseError(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.ResponseOk(resp))
		}
	}
}
