module github.com/libdns/huaweicloud

go 1.23

require github.com/libdns/libdns v0.2.2

require (
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.1.118
	go.mongodb.org/mongo-driver v1.12.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/huaweicloud/huaweicloud-sdk-go-v3 => github.com/devhaozi/huaweicloud-sdk-go-v3 v0.1.85-0.20241018201231-a42a624bc53e
