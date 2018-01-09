//create: 2018/01/02 19:41:47 change: 2018/01/03 16:00:03 lijiaocn@foxmail.com
package client

import (
	"github.com/lijiaocn/handbook-kubernetes/codes/03-client-go/05-framework/config"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

var (
	restclients map[schema.GroupVersion]*rest.RESTClient
)

func InitRESTClient(cmd *config.CmdLine, groups ...schema.GroupVersion) error {

	kconfig, err := ConvertToRestConfig(cmd)
	if err != nil {
		return err
	}

	restclients = make(map[schema.GroupVersion]*rest.RESTClient)

	for _, group := range groups {
		kconfig.ContentConfig.GroupVersion = &group
		kconfig.ContentConfig.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
		kconfig.APIPath = "/api"

		restclient, err := rest.RESTClientFor(kconfig)
		if err != nil {
			return err
		}
		restclients[group] = restclient
	}
	return nil
}

func GetRESTClient(group schema.GroupVersion) *rest.RESTClient {
	return restclients[group]
}
