PREFIX=harbor.enncloud.cn/lijiaob
TAG="latest"
IMAGE=${PREFIX}/04-incluster:${TAG}

.PHONY: build image push run

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0  go build
image:
	docker build -t ${IMAGE} .
push:
	docker push ${IMAGE}
run:
	kubectl -n kube-system run -i --rm informer --image=${IMAGE} --restart=Never
