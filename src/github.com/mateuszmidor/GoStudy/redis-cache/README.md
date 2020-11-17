# redis cache

Example use: <https://github.com/go-redis/redis>

```bash
docker exec -it myredis redis-cli # after ./run_all.sh

127.0.0.1:6379> AUTH mypass
OK

127.0.0.1:6379> set COLOR RED
OK

127.0.0.1:6379> get COLOR
"RED"
 ```