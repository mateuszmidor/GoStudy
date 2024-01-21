# Debug go program running in local Kubernetes (using minikube)

What is impotant: in case of Deployment, the ReadinessProbe and LivenessProbe must be disabled or a breakpoint will trigger a problem

1. Build docker image, it has the app itself and delve debugger inside
    ```sh
	minikube image build -t debug-kubernetes .
    ```

1. Run the app under control of delve debugger
    ```sh
	kubectl run hello-world --image=docker.io/library/debug-kubernetes:latest --image-pull-policy=Never -- /go/bin/dlv --continue --listen=:2345 --headless=true --accept-multiclient --api-version=2 exec /app/HelloWorld
	kubectl wait --for=condition=ready pod hello-world
	kubectl logs hello-world -f &
	kubectl port-forward pod/hello-world 2345:2345
    ```

1. Attach VS Code to delve in container at port 2345. `launch.json` config for remote debugging:
    ```json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Attach to remote Delve session",
                "type": "go",
                "request": "attach",
                "mode": "remote",
                "substitutePath": [
                    { "from": "${workspaceFolder}/", "to": "/app/" },
                ],
                "port": 2345,
                "host": "127.0.0.1",
                "showLog": true,
                "trace": "log",
                "logOutput": "rpc"
            }
        ]
    }
    ```
    **Note** for Delve sources and breakpoints to work, the substitutePath must map directory where the source code is on the host to where the `go build` was executed from (in our case `go build` is executed from /app/ folder in docker container)