package controller

import (
	"fmt"
	"github.com/sairam546/kube-events/pkg/utils"


	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Start() {
	fmt.Println("in controller.go Start")

	var kubeClient kubernetes.Interface
	_, err := rest.InClusterConfig()
	if err != nil {
		kubeClient = utils.GetClientOutOfCluster()
	} else {
		kubeClient = utils.GetClient()
	}

	fmt.Println(kubeClient)
}