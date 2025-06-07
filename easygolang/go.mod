module github.com/mrchuanxu/fresh_golang/easygolang

go 1.24.3

require github.com/mrchuanxu/vito_infra v0.1.2

replace src => ../src

replace google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace github.com/mrchuanxu/vito_infra v0.1.2 => ../../vito_infra
