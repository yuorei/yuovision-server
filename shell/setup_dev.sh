#!/bin/bash

# MinIO コンテナを起動
docker run --name minio \
  -p 9000:9000 \
  -p 9001:9001 \
  -v ~/minio/data:/data \
  -e "MINIO_ROOT_USER=admin" \
  -e "MINIO_ROOT_PASSWORD=password" \
  -d \
  quay.io/minio/minio server /data --console-address ":9001"

# Redis コンテナを起動
docker run --name redis \
  -d -p 6379:6379 \
  redis
