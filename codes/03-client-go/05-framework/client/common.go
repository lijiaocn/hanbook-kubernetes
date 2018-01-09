//create: 2018/01/02 19:56:37 change: 2018/01/03 14:50:28 lijiaocn@foxmail.com
package client

import (
	"errors"
	"github.com/lijiaocn/handbook-kubernetes/codes/03-client-go/05-framework/config"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func ConvertToRestConfig(cmd *config.CmdLine) (kconfig *rest.Config, err error) {
	switch cmd.Auth {
	case config.AUTH_TOKEN:
		kconfig = &rest.Config{
			Host:            cmd.Host,
			BearerToken:     cmd.Token,
			TLSClientConfig: rest.TLSClientConfig{Insecure: cmd.SkipTLS},
		}
	case config.AUTH_INCLUSTER:
		kconfig, err = rest.InClusterConfig()
	case config.AUTH_KUBECONFIG:
		kconfig, err = clientcmd.BuildConfigFromFlags("", cmd.KubeConfig)
	default:
		err = errors.New("unkown auth method")
	}
	return
}
