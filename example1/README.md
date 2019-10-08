
generate proto model

```
  protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. example1/proto/*.proto 
```

TO FIX panic: /debug/requests is already registered.
```
  rm -rf $GOPATH/src/go.etcd.io/etcd/vendor/golang.org/x/net/trace
  rm -rf $GOPATH/src/github.com/coreos/etcd/vendor/golang.org/x/net/trace
```
