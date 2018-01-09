PREFIX=harbor.enncloud.cn/lijiaob
TAG="latest"
IMAGE=${PREFIX}/05-framework:${TAG}

.PHONY: host build image push run

host:
	go build
build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0  go build
image:
	docker build -t ${IMAGE} .
push:
	docker push ${IMAGE}
run:
#	kubectl -n kube-system run -i --rm --restart=Never framework --image=${IMAGE}  -- -auth=incluster
#	./05-framework -auth=token -host=https://10.39.0.105:6443 -token=bf8cb8725efab8c4
#	./05-framework -auth=kubeconfig -kubeconfig=./kubeconfig
