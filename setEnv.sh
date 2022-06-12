#!/bin/sh

# Local envs for gateway
export GATEWAY_PORT=8888

# Local envs for user service
export USER_SVC_PORT=8080
export ETCD_HOST=127.0.0.1:2379
export POSTGRES_URL='postgres://go:password@localhost:5432/cyntra_local?sslmode=disable'
export REDIS_HOST=localhost:6379
export JWT_SECRET=supersecret
export JWT_EXPIRY_SECONDS=900
export JWT_REFRESH_EXPIRY_SECONDS=604800
export REDIS_URL=redis://localhost:6379/0
