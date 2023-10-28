# Debug go program running in Docker

https://golangforall.com/en/post/go-docker-delve-remote-debug.html

1. Build docker image, it has the app itself and delve debugger inside
    ```sh
    docker build . -t debug-docker
    ```

1. Run the app under control of delve debugger
    ```
    docker run -t --rm -p 2345:2345 debug-docker /go/bin/dlv --continue --listen=:2345 --headless=true --accept-multiclient --api-version=2 exec /app/HelloWorld
    ```

1. Attach VS Code to delve in container at port 2345. `launch.json` config for remote debugging:
    ```json
    {
        "name": "Attach to remote Delve session",
        "type": "go",
        "request": "attach",
        "mode": "remote",
        "remotePath": "/app/",
        "port": 2345,
        "host": "127.0.0.1",
        "showLog": true,
        "trace": "log",
        "logOutput": "rpc"
    }
    ```
    **Note** the remotePath needs to reflect the folder in which the go application has been compiled for sources and breakpoints to work