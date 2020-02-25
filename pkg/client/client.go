package client

import (
	"github.com/sairam546/kube-events/pkg/controller"
)

func Run() {
	controller.Start()
}