//create: 2018/01/02 18:52:14 change: 2018/01/03 15:38:32 lijiaocn@foxmail.com
package config

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/golang/glog"
)

const (
	//auth methods
	AUTH_INCLUSTER  = "incluster"
	AUTH_KUBECONFIG = "kubeconfig"
	AUTH_TOKEN      = "token"
)

func init() {
	flag.BoolVar(&cmdline.Help, "help", false, "show usage")

	flag.StringVar(&cmdline.Auth, "auth", "", "auth method: "+
		AUTH_INCLUSTER+","+AUTH_KUBECONFIG+","+AUTH_TOKEN)

	flag.StringVar(&cmdline.KubeConfig, "kubeconfig", "", "kubeconfig file")

	flag.StringVar(&cmdline.Host, "host", "", "kubernetes api host")
	flag.StringVar(&cmdline.Token, "token", "", "user's bearer token")
	flag.BoolVar(&cmdline.SkipTLS, "skiptls", true, "don't verify TLS certificate")
	//TODO: ADD FLAGS
}

type CmdLine struct {
	Help       bool
	Auth       string
	KubeConfig string
	Host       string
	Token      string
	SkipTLS    bool
}

var cmdline CmdLine

func ValidCheck() error {

	//TODO: ADD MORE CHECKS
	if b, err := json.Marshal(cmdline); err != nil {
		glog.Exitf("marshal cmdline fail: %s\n", err.Error())
	} else {
		glog.Infof("cmdline is: %s\n", string(b))
	}

	switch cmdline.Auth {
	case "":
		return errors.New("auth method is not set by -auth")
	case AUTH_INCLUSTER:

	case AUTH_KUBECONFIG:
		if cmdline.KubeConfig == "" {
			return errors.New("must specify the kubeconfig file by -kubeconfig")
		}
	case AUTH_TOKEN:
		if cmdline.Host == "" {
			return errors.New("must specify the host by -host")
		}
	default:
		return errors.New("unkown auth method: " + cmdline.Auth)
	}
	return nil
}

func GetCmdLine() *CmdLine {
	return &cmdline
}
