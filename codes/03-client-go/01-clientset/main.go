//create: 2017/11/15 16:05:17 change: 2018/01/03 14:01:43 author:lijiao
package main

import (
	"fmt"
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config := rest.Config{
		Host:            "https://10.39.0.105:6443",
		BearerToken:     "bf8cb8725efab8c4",
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}
	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster.\n", len(pods.Items))
}
