module github.com/realotz/cratos

go 1.15

require (
	github.com/go-kratos/kratos/v2 v2.0.0-alpha5
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jinzhu/copier v0.2.5 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210228012217-479acdf4ea46 // indirect
	google.golang.org/genproto v0.0.0-20210226172003-ab064af71705
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/datatypes v1.0.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.12
	istio.io/api v0.0.0-20210302211031-2e1e4d7e6f4b
	istio.io/client-go v0.0.0-20210302212237-0ceb5261e848
	k8s.io/apiextensions-apiserver v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
)

replace google.golang.org/grpc => google.golang.org/grpc v1.35.0
