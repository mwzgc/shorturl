# shorturl
https://github.com/tal-tech/go-zero/blob/master/doc/shorturl.md


go run api/shorturl.go -f api/etc/shorturl-api.yaml
go run rpc/transform/transform.go -f rpc/transform/etc/transform.yaml

curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn"
curl -i "http://localhost:8888/expand?shorten=f35b2a"