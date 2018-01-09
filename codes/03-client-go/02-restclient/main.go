//create: 2017/12/28 15:25:24 change: 2017/12/28 16:47:51 lijiaocn@foxmail.com
package main

import (
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

func main() {
	config := rest.Config{
		Host:            "https://10.39.0.105:6443",
		APIPath:         "/api",
		BearerToken:     "bf8cb8725efab8c4",
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
		ContentConfig: rest.ContentConfig{
			/*
				$ kubectl api-versions
				kubectl api-versions
				apiextensions.k8s.io/v1beta1
				apiregistration.k8s.io/v1beta1
				apps/v1beta1
				authentication.k8s.io/v1
				...
				v1
			*/
			GroupVersion:         &v1.SchemeGroupVersion,
			NegotiatedSerializer: serializer.DirectCodecFactory{CodecFactory: scheme.Codecs},
		},
	}

	restclient, err := rest.RESTClientFor(&config)
	if err != nil {
		panic(err.Error())
	}

	pods := &v1.PodList{}
	if err = restclient.Get().
		Namespace(v1.NamespaceAll).
		Resource("pods").
		Do().
		Into(pods); err != nil {
		println("error: ", err.Error())
	}

	fmt.Printf("There are %d pods in the cluster.\n", len(pods.Items))
}
