#!/bin/bash

# 检查是否安装了Node.js和npm
if ! command -v node &> /dev/null || ! command -v npm &> /dev/null; then
    echo "错误: 需要安装Node.js和npm"
    echo "请访问 https://nodejs.org/ 下载并安装"
    exit 1
fi

# 检查是否提供了OpenAPI规范文件路径
if [ $# -eq 0 ]; then
    echo "错误: 缺少OpenAPI规范文件路径"
    echo "用法: $0 <openapi.yaml路径>"
    exit 1
fi

OPENAPI_PATH="$1"

# 检查OpenAPI规范文件是否存在
if [ ! -f "$OPENAPI_PATH" ]; then
    echo "错误: 找不到OpenAPI规范文件: $OPENAPI_PATH"
    echo "请确保提供了正确的文件路径"
    exit 1
fi

# 生成Redoc HTML文档
echo "生成Redoc文档..."
npx @redocly/cli build-docs "$OPENAPI_PATH" --output index.html

echo "成功生成文档: index.html"