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

// 上传阿里云图片/音频
func UploadOssHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析 multipart form
		err := r.ParseMultipartForm(10 << 20) // 10MB max
		if err != nil {
			logx.WithContext(r.Context()).Errorf("解析表单错误,err:%+v", err)
			httpx.OkJsonCtx(r.Context(), w, response.ResponseError(errcode.ErrInvalidParams))
			return
		}

		// 获取文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			logx.WithContext(r.Context()).Errorf("获取文件错误,err:%+v", err)
			httpx.OkJsonCtx(r.Context(), w, response.ResponseError(errcode.ErrInvalidParams))
			return
		}
		defer file.Close()

		// 获取文件类型
		fileType := r.FormValue("file_type")
		req := types.UploadOssReq{
			FileType: fileType,
		}

		l := aliyun.NewUploadOssLogic(r.Context(), svcCtx)
		resp, err := l.UploadOss(&req, file, handler.Filename)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("handler err:%+v", err)
			httpx.OkJsonCtx(r.Context(), w, response.ResponseError(err))
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.ResponseOk(resp))
		}
	}
}
