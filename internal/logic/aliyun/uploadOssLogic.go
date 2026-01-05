// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package aliyun

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"github.com/gpencil/upload/internal/svc"
	"github.com/gpencil/upload/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传阿里云图片/音频
func NewUploadOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadOssLogic {
	return &UploadOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadOssLogic) UploadOss(req *types.UploadOssReq, file multipart.File, filename string) (resp *types.UploadOssResp, err error) {
	// 从环境变量读取阿里云配置
	accessKeyID := os.Getenv(l.svcCtx.Config.Aliyun.AccessKeyID)
	accessKeySecret := os.Getenv(l.svcCtx.Config.Aliyun.AccessKeySecret)

	// 调试日志：打印环境变量名称和是否获取到值
	l.Infof("读取环境变量: %s, 是否获取到: %v (长度: %d)",
		l.svcCtx.Config.Aliyun.AccessKeyID,
		accessKeyID != "",
		len(accessKeyID))

	if accessKeyID == "" || accessKeySecret == "" {
		return nil, fmt.Errorf("阿里云配置未设置，请设置环境变量 %s 和 %s",
			l.svcCtx.Config.Aliyun.AccessKeyID,
			l.svcCtx.Config.Aliyun.AccessKeySecret)
	}

	// 调试日志：打印部分 AccessKeyID（前4位和后4位）
	if len(accessKeyID) > 8 {
		l.Infof("使用 AccessKeyID: %s...%s", accessKeyID[:4], accessKeyID[len(accessKeyID)-4:])
	}

	// 创建 OSS 客户端
	client, err := oss.New(l.svcCtx.Config.Aliyun.Endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		l.Errorf("创建 OSS 客户端失败: %v", err)
		return nil, fmt.Errorf("创建 OSS 客户端失败")
	}

	// 获取 Bucket
	bucket, err := client.Bucket(l.svcCtx.Config.Aliyun.BucketName)
	if err != nil {
		l.Errorf("获取 Bucket 失败: %v", err)
		return nil, fmt.Errorf("获取 Bucket 失败")
	}

	// 根据文件类型确定上传目录
	var uploadDir string
	switch req.FileType {
	case "icon":
		uploadDir = l.svcCtx.Config.Aliyun.IconDir
	case "preview":
		uploadDir = l.svcCtx.Config.Aliyun.PreviewDir
	default:
		return nil, fmt.Errorf("不支持的文件类型: %s", req.FileType)
	}

	// 生成唯一文件名
	ext := filepath.Ext(filename)
	uniqueFilename := fmt.Sprintf("%s_%s%s", time.Now().Format("20060102150405"), uuid.New().String()[:8], ext)
	objectKey := fmt.Sprintf("%s/%s", uploadDir, uniqueFilename)

	// 上传文件
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		l.Errorf("上传文件到 OSS 失败: %v", err)
		return nil, fmt.Errorf("上传文件失败")
	}

	// 返回文件 URL
	fileURL := fmt.Sprintf("https://%s.%s/%s", l.svcCtx.Config.Aliyun.BucketName, l.svcCtx.Config.Aliyun.Endpoint, objectKey)

	return &types.UploadOssResp{
		Name: fileURL,
	}, nil
}
