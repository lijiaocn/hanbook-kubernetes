//create: 2017/12/28 16:50:01 change: 2017/12/28 18:33:41 lijiaocn@foxmail.com
package main

import (
	"flag"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "./kubeconfig", "kubeconfig file")
}

func main() {
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	{ //RESTCLient
		config.GroupVersion = &v1.SchemeGroupVersion
		config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
		config.APIPath = "/api"
		restclient, err := rest.RESTClientFor(config)
		if err != nil {
			println("restclient: ", err.Error())
		}

		pods := &v1.PodList{}
		if err = restclient.Get().
			Namespace(v1.NamespaceAll).
			Resource("pods").
			Do().
			Into(pods); err != nil {
			println("restclient: ", err.Error())
		}

		fmt.Printf("restclient: There are %d pods in the cluster.\n", len(pods.Items))
	}

	{ //Clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			println("clientset: ", err.Error())
		}

		pods := &v1.PodList{}
		if pods, err = clientset.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{}); err != nil {
			println("clientst: ", err.Error())
		}

		fmt.Printf("clientset: There are %d pods in the cluster.\n", len(pods.Items))
	}
}
