---
layout: post
title: 02-client-go第一次使用
author: lijiaocn
createdate: 2017/12/27 13:40:54
changedate: 2017/12/27 14:10:18
categories: chapter03
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 创建项目

新建一个`$GOPATH/src/<指定的目录>/01-first`子目录，用glide引入client-go：

{% highlight shell  %}
$ glide create
[INFO]	Generating a YAML configuration file and guessing the dependencies
[INFO]	Attempting to import from other package managers (use --skip-import to skip)
[INFO]	Scanning code to look for dependencies
[INFO]	Writing configuration file (glide.yaml)
[INFO]	Would you like Glide to help you find ways to improve your glide.yaml configuration?
[INFO]	If you want to revisit this step you can use the config-wizard command at any time.
[INFO]	Yes (Y) or No (N)?
N
[INFO]	You can now edit the glide.yaml file. Consider:
[INFO]	--> Using versions and ranges. See https://glide.sh/docs/versions/
[INFO]	--> Adding additional metadata. See https://glide.sh/docs/glide.yaml/
[INFO]	--> Running the config-wizard command to improve the versions in your configuration

$ glide get k8s/client-go
[INFO]	Preparing to install 1 package.
[INFO]	Attempting to get package k8s.io/client-go
[INFO]	--> Gathering release information for k8s.io/client-go
[INFO]	The package k8s.io/client-go appears to have Semantic Version releases (http://semver.org).
[INFO]	The latest release is v6.0.0. You are currently not using a release. Would you like
[INFO]	to use this release? Yes (Y) or No (N)
Y
[INFO]	The package k8s.io/client-go appears to use semantic versions (http://semver.org).
[INFO]	Would you like to track the latest minor or patch releases (major.minor.patch)?
[INFO]	The choices are:
[INFO]	 - Tracking minor version releases would use '>= 6.0.0, < 7.0.0' ('^6.0.0')
[INFO]	 - Tracking patch version releases would use '>= 6.0.0, < 6.1.0' ('~6.0.0')
[INFO]	 - Skip using ranges
[INFO]	For more information on Glide versions and ranges see https://glide.sh/docs/versions
[INFO]	Minor (M), Patch (P), or Skip Ranges (S)?
P
[INFO]	--> Adding k8s.io/client-go to your configuration with the version ~6.0.0
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for k8s.io/client-go
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Detected semantic version. Setting version for k8s.io/client-go to v6.0.0
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting k8s.io/client-go
[INFO]	Replacing existing vendor dependencies
{% endhighlight %}

## 开发

创建源码文件main.go，输入下面的代码：

{% highlight go  %}
//create: 2017/11/15 16:05:17 change: 2017/12/27 13:47:48 author:lijiao
package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config := rest.Config{
		Host:            "https://10.39.0.105:6443",   //替换成你的集群地址
		APIPath:         "/",
		Prefix:          "",
		BearerToken:     "bf8cb8725efab8c4",           //替换成你的token
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}
	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("lijiaocn").List(v1.ListOptions{})  //替换成你namespace
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster.\n", len(pods.Items))
}
{% endhighlight %}

## 运行

运行之前，先用glide将依赖的代码下载到项目本地：

{% highlight shell %}
$ glide update
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for k8s.io/client-go
[INFO]	--> Detected semantic version. Setting version for k8s.io/client-go to v6.0.0
[INFO]	Resolving imports
[INFO]	Found Godeps.json file in /Users/lijiao/.glide/cache/src/https-k8s.io-apimachinery
[INFO]	--> Parsing Godeps metadata...
[INFO]	--> Fetching updates for github.com/go-openapi/spec
[INFO]	--> Setting version for github.com/go-openapi/spec to 7abd5745472fff5eb3685386d5fb8bf38683154d.
[INFO]	--> Fetching updates for github.com/gogo/protobuf
[INFO]	--> Setting version for github.com/gogo/protobuf to c0656edd0d9eab7c66d1eb0c568f9039345796f7.
[INFO]	--> Fetching updates for github.com/google/gofuzz
[INFO]	--> Setting version for github.com/google/gofuzz to 44d81051d367757e1c7c6a5a86423ece9afcf63c.
...
[INFO]	--> Exporting k8s.io/apimachinery
[INFO]	--> Exporting gopkg.in/yaml.v2
[INFO]	--> Exporting k8s.io/kube-openapi
[INFO]	--> Exporting k8s.io/api
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 28 dependencies.
{% endhighlight %}

编译运行：
{% highlight shell  %}
$ go build
$ ls
01-first   glide.lock glide.yaml main.go    vendor
$ ./01-first
There are 0 pods in the cluster.
{% endhighlight %}
