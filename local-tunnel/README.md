# local tunnel

Should enable access from internet to your localhost:port, terminating https (so you only need to run simple http server)  
https://theboroer.github.io/localtunnel-www/  

## Run on linux

```bash
go run .
make macos 
> your url is: https://dob-mateusz-test.loca.lt
```

## Run on macos

```bash
go run .
make macos  
> your url is: https://dob-mateusz-test.loca.lt
```

, and then go to displayed URL and enter public IP to authorize access.

## Easier done with serveo.net 

```sh
ssh -R 80:localhost:33000 serveo.net 
> Forwarding HTTP traffic from https://ff5aa8fe6c01e1817282f053c38637c0.serveo.net
```