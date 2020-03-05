module randokiak-api

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/apache/pulsar-client-go v0.0.0-20200214184451-fc390a6a37f3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/golang/mock v1.3.0
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	golang.org/x/tools v0.0.0-20200211205636-11eff242d136 // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
	rdk.io/protocols v0.0.0
	rdk.io/utils v0.0.0
)

replace rdk.io/protocols v0.0.0 => ../protocols

replace rdk.io/utils v0.0.0 => ../utils
