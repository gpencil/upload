# 声线上传管理系统

goctl api go -api upload.api -dir . -style goZero

声线上传管理系统是一个用于管理 TTS（文本转语音）声线的 Web 应用，支持上传声线图标、试听音频，并将声线信息存储到数据库。

## 功能特性

- 文件上传到阿里云 OSS
  - 声线图标上传（存储到 ttsIconV1 目录）
  - 试听音频上传（存储到 ttsPreV1 目录）
- 声线信息管理
  - 创建新声线
  - 更新现有声线
- 统一配置查询接口
- 友好的 Web 界面

## 配置说明

### 1. 数据库配置

```yaml
Mysql:
  DataSource: root:password@tcp(127.0.0.1:3306)/tts_db?charset=utf8mb4&parseTime=True&loc=Local
```

### 2. 阿里云配置

```yaml
Aliyun:
  Endpoint: oss-cn-hangzhou.aliyuncs.com
  BucketName: lunalab-res
  IconDir: ttsIconV1
  PreviewDir: ttsPreV1
```

### 3. 环境变量

设置阿里云访问密钥（必须）：

```bash
export AliyunAccessKeyID=your_access_key_id
export AliyunAccessKeySecret=your_access_key_secret
```


## 启动服务

服务将在 `http://localhost:8889` 启动。

## 使用说明

### Web 界面

访问 `http://localhost:8889` 打开 Web 管理界面。

**操作流程：**

1. **上传文件**
   - 点击"上传声线图标"区域，选择图片文件
   - 点击"上传试听音频"区域，选择音频文件
   - 上传成功后，URL 会自动填充到下方表单

2. **保存**
   - 点击"保存声线信息"按钮
   - 系统会自动判断是创建新声线还是更新现有声线

### API 接口

#### 1. 上传文件到 OSS

```bash
POST /upload/oss
Content-Type: multipart/form-data

参数：
- file: 文件内容
- file_type: 文件类型（icon=图标, preview=试听）

响应：
{
  "code": 0,
  "msg": "success",
  "data": {
    "name": "https://lunalab-res.oss-cn-hangzhou.aliyuncs.com/ttsIconV1/20240101120000_abc123.jpg"
  }
}
```

#### 2. 创建/更新声线

```bash
POST /upload/voice
Content-Type: application/json

{
  "voice_id": "alice_warm",
  "name": "温柔女声",
  "scenario_id": 1,
  "description": "温柔的女性声音",
  "icon_url": "https://...",
  "preview_url": "https://...",
  "sort_order": 100,
  "style": "温柔",
  "language": "zh-CN",
  "status": 1,
  "age_group": "adult",
  "gender": "female"
}

响应：
{
  "code": 0,
  "msg": "success",
  "data": {}
}
```

#### 3. 获取统一配置

```bash
GET /config?voice_id=alice_warm&scenario_id=1

响应：
{
  "code": 0,
  "msg": "success",
  "data": {
    "fullConfigs": [...]
  }
}
```


## 生成代码

```bash
goctl api go -api upload.api -dir . -style goZero
```


## 注意事项
上传文件大小限制为 10MB
修改文件上传大小限制
A: 修改 `uploadOssHandler.go` 中的 `ParseMultipartForm(10 << 20)` 参数。