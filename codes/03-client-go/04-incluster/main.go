//create: 2018/01/02 10:24:03 change: 2018/01/02 10:33:17 lijiaocn@foxmail.com
package main

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		println("clientset: ", err.Error())
	}

	pods := &v1.PodList{}
	if pods, err = clientset.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{}); err != nil {
		println("clientset: ", err.Error())
	}

	fmt.Printf("clientset: There are %d pods in the cluster.\n", len(pods.Items))
}
