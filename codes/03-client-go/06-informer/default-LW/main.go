//create: 2018/01/02 11:01:07 change: 2018/01/02 11:26:15 lijiaocn@foxmail.com
package main

import (
	"flag"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
	"syscall"
)

func AddPod(obj interface{}) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		println("Wrong! not a pod in AddFunc")
		return
	}
	println("add a pod: ", pod.Name)
	return
}

func UpdatePod(oldObj, newObj interface{}) {
	pod1, ok1 := oldObj.(*v1.Pod)
	pod2, ok2 := newObj.(*v1.Pod)
	if !(ok1 && ok2) {
		println("Wrong! not a pod in OnUpdate")
		return
	}
	println("old pod: ", pod1.Name, "\t", pod1.ResourceVersion)
	println("new pod: ", pod2.Name, "\t", pod1.ResourceVersion)
	return
}

func DeletePod(obj interface{}) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		println("Wrong! not a pod in OnDelete")
		return
	}
	println("delete pod: ", pod.Name)
	return
}

func DisplayPod(podStore cache.Store) {
	for i, obj := range podStore.List() {
		pod, ok := obj.(*v1.Pod)
		if ok != true {
			println(i, "\t", "not pod")
			continue
		}
		println(i, "\t", pod.Name)
	}
}

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "./kubeconfig", "kubeconfig file")
}

func main() {

	receive := make(chan os.Signal)
	signal.Notify(receive, syscall.SIGUSR1, syscall.SIGUSR2)
	println("display all pods: kill -s SIGUSR1 ", syscall.Getpid())

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.APIPath = "/api"

	restclient, err := rest.RESTClientFor(config)
	if err != nil {
		println("restclient: ", err.Error())
		return
	}

	lw := cache.NewListWatchFromClient(restclient, "pods", v1.NamespaceAll, fields.SelectorFromSet(nil))

	podHandler := cache.ResourceEventHandlerFuncs{
		AddFunc:    AddPod,
		UpdateFunc: UpdatePod,
		DeleteFunc: DeletePod,
	}

	podStore, PodController := cache.NewInformer(lw, &v1.Pod{}, 0, podHandler)

	stopChan := make(chan struct{}, 1)

	go PodController.Run(stopChan)

	for {
		select {
		case s := <-receive:
			switch s {
			case syscall.SIGUSR1:
				DisplayPod(podStore)
			case syscall.SIGUSR2:
				var stop struct{}
				stopChan <- stop
			default:
				continue
			}
		}
	}
}
