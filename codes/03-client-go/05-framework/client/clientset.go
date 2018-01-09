//create: 2018/01/02 19:20:09 change: 2018/01/03 14:08:32 lijiaocn@foxmail.com
package client

import (
	"github.com/lijiaocn/handbook-kubernetes/codes/03-client-go/05-framework/config"
	"k8s.io/client-go/kubernetes"
)

var (
	clientset *kubernetes.Clientset
)

func InitClientSet(cmd *config.CmdLine) error {

	kconfig, err := ConvertToRestConfig(cmd)
	if err != nil {
		return err
	}

	clientset, err = kubernetes.NewForConfig(kconfig)
	if err != nil {
		return err
	}
	return nil
}

func GetClientSet() *kubernetes.Clientset {
	return clientset
}
