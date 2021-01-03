CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o svcp github.com/morvencao/kube-mutating-webhook-tutorial/cmd

IMAGE=registry.cn-hangzhou.aliyuncs.com/launcher-agent-only-test/svcp:v1
docker build -t ${IMAGE} .
docker push ${IMAGE}
docker rmi -f ${IMAGE}
