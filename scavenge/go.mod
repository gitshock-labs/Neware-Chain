module scavenge

go 1.18

require (
	github.com/cosmos/cosmos-sdk v0.45.5
	github.com/cosmos/ibc-go/v3 v3.0.1
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ignite/cli v0.23.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/spn v0.2.1-0.20220708132853-26a17f03c072
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	google.golang.org/genproto v0.0.0-20220805133916-01dd62135a58
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.2 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)