package rpcs

import (
	"github.com/micro/go-micro"
)

type RouteRegister func(server micro.Service)

//export
func RegisterRoute(service micro.Service, registers ...RouteRegister) {
	if len(registers) > 0 {
		for _, reg := range registers {
			reg(service)
		}
	}
}
