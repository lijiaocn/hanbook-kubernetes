---
layout: post
title: "01-client-go的安装"
author: lijiaocn
createdate: 2017/12/26 14:02:16
changedate: 2017/12/27 13:40:51
categories: chapter03
tags:
keywords:
description: 

---

* auto-gen TOC:
{:toc}

## 说明

[client-go][1]是kubernetes的一个子项目，它是一个用来与kubernetes集群交互的library，用go语言实现。

client-go主要包含下面几个package：

	kubernetes:  提供用来访问kubernetes集群的Clientset
	discovery:   提供用来查询kubernetes集群支持的api的DiscoveryClient
	dynamic:     提供用来操作kubernetes集群的任意资源的Client
	transport:   用于设置认证信息、建立连接
	tools/cache: 用于Controller的开发

client-go的版本命名使用的是[Semantic Versioning][3]规则，详情可以参考：[怎样为软件的不同版本命名？][2]。

## 安装

client-go将大部分依赖的package都放入了vendor，除了`k8s.io/apimachinery`和`github.com/glog`。

因此使用`go get`安装了client-go的代码之后还需要安装apimachinery和glog。

	go get k8s.io/client-go/...
	go get -u k8s.io/apimachinery/...
	go get -u github.com/glog/...

## 版本

在使用client-go的时候，为了避免混乱，最好明确指定要使用的版本，这个过程可以用依赖管理工实现。

### Godep

[godep][5]是一个比较早期的依赖管理工具，client-go和kubernetes都在使用。

可以用下面的命令安装godep：

	go get github.com/tools/godep

这里以要使用client-go的`v6.0.0`版本为例，示范如何用godep指定依赖版本。

首先要到本地的client-go目录中，将client-go切换到目标分支`v6.0.0`。

	cd $GOPATH/src/k8s.io/client-go
	git checkout v6.0.0

然后在client-go目录中，设置client-go依赖的包添加到GOPATH中：

	godep restore ./...

之后，在正在进行的项目中创建一个引用了client-go的go文件：

	import (
	    "k8s.io/client-go"
	)

>项目中没有引用的依赖包，不会被下一步操作中的godep保存到项目的本地。

最后，用godep将项目的依赖保存到项目的本地：

	godep save ./...

将项目的依赖保存到项目本地之后，当前项目就不在使用系统中安装的client-go了。

可以将系统中的client-go切换为其它版本，供其它项目使用。

### Glide

[glide][6]是另一个使用比较多的依赖管理工具。

在linux上可以用下面的方法安装：

	curl https://glide.sh/get | sh

Mac上还可以用brew安装：

	brew install glide

可以用glide的命令引入client-go:

	glide create
	glide get k8s.io/client-go

之后可以编辑项目中的`glide.yaml`文件，指定要使用的版本:

	package: ( your project's import path ) # e.g. github.com/foo/bar
	import:
	- package: k8s.io/client-go
	  version: v6.0.0

修改之后执行:

	glide up -v

### Dep

[dep][7]是一个计划成为go的标准工具的依赖管理工具。截至`2017-12-26 16:39:42`，client-go不支持dep。

## 参考

1. [github: client-go][1]
2. [怎样为软件的不同版本命名？][2]
3. [Semantic Versioning][3]
4. [Installing client-go][4]
5. [godep][5]
6. [glide][6]
7. [dep][7]

[1]: https://github.com/kubernetes/client-go  "github: client-go" 
[2]: http://www.lijiaocn.com/%E6%96%B9%E6%B3%95/2017/12/26/sofeware-version-semver.html "怎样为软件的不同版本命名？"
[3]: https://semver.org/ "Semantic Versioning"
[4]: https://github.com/kubernetes/client-go/blob/master/INSTALL.md "Installing client-go"
[5]: https://github.com/tools/godep "godep"
[6]: https://github.com/Masterminds/glide "glide"
[7]: https://github.com/golang/dep  "dep"
