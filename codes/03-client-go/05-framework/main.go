//create: 2018/01/02 18:25:55 change: 2018/01/03 16:02:02 lijiaocn@foxmail.com
package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/lijiaocn/handbook-kubernetes/codes/03-client-go/05-framework/client"
	"github.com/lijiaocn/handbook-kubernetes/codes/03-client-go/05-framework/config"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func run() {

	{
		clientset := client.GetClientSet()
		pods, err := clientset.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{})
		if err != nil {
			println("error: ", err.Error())
		}
		fmt.Printf("There are %d pods in the cluster.\n", len(pods.Items))
	}

	{
		restclient := client.GetRESTClient(v1.SchemeGroupVersion)
		pods := &v1.PodList{}
		if err := restclient.Get().
			Namespace(v1.NamespaceAll).
			Resource("pods").
			Do().
			Into(pods); err != nil {
			println("error: ", err.Error())
		}
		fmt.Printf("There are %d pods in the cluster.\n", len(pods.Items))
	}
}

func main() {
	flag.Parse()

	cmdline := config.GetCmdLine()

	if cmdline.Help {
		flag.Usage()
		return
	}

	if err := config.ValidCheck(); err != nil {
		glog.Exitln(err.Error())
	}

	if err := client.InitClientSet(cmdline); err != nil {
		glog.Exitln(err.Error())
	}

	if err := client.InitRESTClient(cmdline, v1.SchemeGroupVersion); err != nil {
		glog.Exitln(err.Error())
	}

	run()
}
