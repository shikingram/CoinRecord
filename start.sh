#!/bin/bash
export GOPROXY=https://goproxy.io,direct
go build CoinRecord
chmod +x CoinRecord
export https_proxy=http://127.0.0.1:1087/
export http_proxy=http://127.0.0.1:1087/
./CoinRecord
