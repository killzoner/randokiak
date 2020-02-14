module randokiak-gen

go 1.13

require (
	github.com/Pallinder/go-randomdata v1.2.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
	google.golang.org/grpc v1.27.1
	rdk.io/protocols v0.0.0
)

replace rdk.io/protocols v0.0.0 => ../protocols
