# intro

這是一個透過 golang 製作的 JWT generate server ，  
可以透過 API，建立一個合法的 JWT token

# api

[POST] /jwt

[GET] /jwt/verify

# build

```sh
# mac_os
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build

# linux_os
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

# 測試覆蓋率

```sh
# 執行測試
go test ./... -coverprofile=size_coverage.out 

# 使用網頁查看測試報告
go tool cover -html=size_coverage.out 
```

# refer

[jwt](https://jwt.io/)

[GO語言web框架Gin之完全指南](https://segmentfault.com/a/1190000022066618)
