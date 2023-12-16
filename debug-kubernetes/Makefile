IMAGE_NAME=debug-kubernetes

default: image run

image:
	minikube image build -t ${IMAGE_NAME} .

run:
	kubectl run hello-world --image=docker.io/library/debug-kubernetes:latest --image-pull-policy=Never
	kubectl wait --for=condition=ready pod hello-world
	kubectl logs hello-world -f

debug:
	kubectl run hello-world --image=docker.io/library/debug-kubernetes:latest --image-pull-policy=Never -- /go/bin/dlv --continue --listen=:2345 --headless=true --accept-multiclient --api-version=2 exec /app/HelloWorld
	kubectl wait --for=condition=ready pod hello-world
	kubectl port-forward pod/hello-world 2345:2345 &
	kubectl logs hello-world -f

kill:
	kubectl delete pod hello-world