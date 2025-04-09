#!/bin/bash

BUCKET_NAME="speedix-meepo-api-docs"
echo "使用存储桶: $BUCKET_NAME"

# 检查是否已生成index.html
if [ ! -f "index.html" ]; then
  echo "未找到index.html文件，尝试生成..."
  exit 1
fi

# 检查gsutil是否可用
if ! command -v gsutil &> /dev/null; then
  echo "错误: 未找到gsutil命令"
  echo "请确保已安装并配置Google Cloud SDK"
  exit 1
fi

# 检查存储桶是否存在，如果不存在则创建
echo "确保存储桶存在..."
gsutil ls -b "gs://$BUCKET_NAME" &>/dev/null || gsutil mb "gs://$BUCKET_NAME"

# 上传文件到GCS
echo "上传文档到GCS..."
gsutil -h "Cache-Control:public,max-age=3600" cp index.html "gs://$BUCKET_NAME/docs/api/index.html"

# 设置访问权限
echo "设置公共访问权限..."
gsutil acl ch -u AllUsers:R "gs://$BUCKET_NAME/docs/api/index.html"

# 输出访问链接
echo "部署完成!"
echo "访问链接: https://storage.googleapis.com/$BUCKET_NAME/docs/api/index.html"