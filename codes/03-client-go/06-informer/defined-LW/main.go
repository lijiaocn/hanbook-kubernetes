//create: 2017/11/15 16:05:17 change: 2018/01/02 17:31:37 author:lijiao
package main

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"os"
	"os/signal"
	"syscall"
)

func ListPod(clientset *kubernetes.Clientset) func(metav1.ListOptions) (runtime.Object, error) {
	return func(options metav1.ListOptions) (runtime.Object, error) {
		return clientset.CoreV1().Pods(v1.NamespaceAll).List(options)
	}
}

func WatchPod(clientset *kubernetes.Clientset) func(metav1.ListOptions) (watch.Interface, error) {
	return func(options metav1.ListOptions) (watch.Interface, error) {
		return clientset.CoreV1().Pods(v1.NamespaceAll).Watch(options)
	}
}

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

func main() {
	receive := make(chan os.Signal)
	signal.Notify(receive, syscall.SIGUSR1, syscall.SIGUSR2)
	println("display all pods: kill -s SIGUSR1 ", syscall.Getpid())

	config := rest.Config{
		Host:            "https://10.39.0.105:6443",
		APIPath:         "/",
		Prefix:          "",
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
	podLw := cache.ListWatch{
		ListFunc:  ListPod(clientset),
		WatchFunc: WatchPod(clientset),
	}

	podHandler := cache.ResourceEventHandlerFuncs{
		AddFunc:    AddPod,
		UpdateFunc: UpdatePod,
		DeleteFunc: DeletePod,
	}

	podStore, PodController := cache.NewInformer(&podLw, &v1.Pod{}, 0, podHandler)

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
