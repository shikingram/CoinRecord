@echo off
chcp 65001
go env -w GOPROXY=https://goproxy.cn,direct
go build
set https_proxy=http://127.0.0.1:1080
set http_proxy=http://127.0.0.1:1080
CoinRecord.exe