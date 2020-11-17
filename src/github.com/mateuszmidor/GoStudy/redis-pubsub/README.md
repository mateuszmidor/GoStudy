# redis pubsub

Redis pubsub tutorial: <https://www.youtube.com/watch?v=33N1mgiRYK0>
<!-- Example use in GO: <https://github.com/go-redis/redis> -->

## Highlights


## Play around with redis shell (first: ./run_all.sh)

Redis shell commands: <https://redis.io/commands/#pubsub>

```bash
# start server
docker run --rm --name pubsub redis

# start subscriber in terminal 1
docker exec -it pubsub redis-cli
SUBSCRIBE messages

# start publisher in terminal 2
docker exec -it pubsub redis-cli
PUBLISH messages Hello

# subscriber receives: Hello
 ```