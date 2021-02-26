module github.com/realotz/cratos

go 1.15

replace github.com/realotz/cratos/api => ./api

require (
	github.com/go-kratos/kratos/v2 v2.0.0-alpha3
	github.com/go-kratos/kube v0.0.0-20210222030635-0603fdf376c5
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/wire v0.5.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/urfave/cli v1.22.5
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/client-go v0.20.0
)
