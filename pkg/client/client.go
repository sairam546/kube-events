package client

import (
	"fmt"
	"github.com/sairam546/kube-events/pkg/controller"
)

func Run() {
	fmt.Println("in client.go Run")
	controller.Start()
}