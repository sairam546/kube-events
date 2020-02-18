package cmd

import (
	"fmt"
	c "github.com/sairam546/kube-events/pkg/client"
)

func Execute() {
	fmt.Println("in cmd/root.go Execute")
	c.Run()
}