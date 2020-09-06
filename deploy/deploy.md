
# docker 
- docker run -p 6379:6379 --name redis -d redis 
- docker run -d -p 2379:2379 -p 2380:2380 --name etcd quay.io/coreos/etcd /usr/local/bin/etcd --name s1 --listen-client-urls http://0.0.0.0:2379  --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380 --initial-advertise-peer-urls http://0.0.0.0:2380 --initial-cluster s1=http://0.0.0.0:2380 --initial-cluster-token tkn  --initial-cluster-state new
- docker run --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

.run.sh

curl -i "http://18.162.143.215:8888/shorten?url=http://www.xiaoheiban.cn"
curl -i "http://18.162.143.215:8888/expand?shorten=f35b2a"
