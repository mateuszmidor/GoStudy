# redis cache

Redis shell tutorial: <https://youtu.be/Hbt56gFj998?t=466>  
Example use in GO: <https://github.com/go-redis/redis>

## Highlights

- in redis no password is set by default
- no encryption possible at all
- mostly get/set operations; 50'000 operations per second in single thread, 100'000 in multithreaded mode:
  - redis-server --io-threads 8
  - redis-benchmark --threads 8
- https://hub.docker.com/_/redis stores DB under /data, so use ```docker run -v /home:/data redis``` for persistent DB

## Play around with redis shell (first: ./run_all.sh)

Redis shell commands: <https://redis.io/commands/#string>

```bash
docker exec -it myredis redis-cli

# authenticate if password has been set
127.0.0.1:6379> auth mypass
OK

# check connection
127.0.0.1:6379> ping
PONG

# set key to string value
127.0.0.1:6379> set color RED
OK
127.0.0.1:6379> get color
"RED"

# set set key to numeric value
127.0.0.1:6379> set price 19
OK
127.0.0.1:6379> get price
"19"

# increment then decrement number
127.0.0.1:6379> incr price
(integer) 20
127.0.0.1:6379> decr price
(integer) 19

# check if key exists
127.0.0.1:6379> exists color
(integer) 1
127.0.0.1:6379> exists volume
(integer) 0

# delete key
127.0.0.1:6379> del price
(integer) 1

# keyspaces
127.0.0.1:6379> set server:name localhost
OK
127.0.0.1:6379> set server:ip 127.0.0.1
OK
127.0.0.1:6379> get server
(nil)
127.0.0.1:6379> get server:name
"localhost"
127.0.0.1:6379> get server:ip
"127.0.0.1"

# expiring keys
127.0.0.1:6379> set boardingpass Mateusz
OK
127.0.0.1:6379> expire boardingpass 50 # expire in 50 seconds
(integer) 1
127.0.0.1:6379> ttl boardingpass
(integer) 43 # 43 seconds remaining

# expiring key oneliner
127.0.0.1:6379> setex boardingpass 30 Mateusz # expire in 30 seconds
OK

# disabling expiration on key
127.0.0.1:6379> persist boardingpass
(integer) 1

# appending to already existing color value "RED"
127.0.0.1:6379> append color _PALE
(integer) 8
127.0.0.1:6379> get color
"RED_PALE"

# rename key
127.0.0.1:6379> rename color hue
OK
127.0.0.1:6379> get hue
"RED_PALE"

# lists - add 
127.0.0.1:6379> lpush util bowl # add to list head
(integer) 1
127.0.0.1:6379> lpush util pot # add to list head
(integer) 2
127.0.0.1:6379> lpush util pan # add to list head
(integer) 3
127.0.0.1:6379> lrange util 0 -1 # get elemnts from 0 to the end
1) "pan"
2) "pot"
3) "bowl"
127.0.0.1:6379> rpush util shotglass # add to list TAIL
(integer) 4
127.0.0.1:6379> lrange util 0 -1
1) "pan"
2) "pot"
3) "bowl"
4) "shotglass"

# lists - get length
127.0.0.1:6379> llen util
(integer) 4

# sets
127.0.0.1:6379> sadd cars honda
(integer) 1
127.0.0.1:6379> sadd cars ford
(integer) 1
127.0.0.1:6379> sadd cars bmw
(integer) 1
127.0.0.1:6379> sismember cars ford # check if ford is in cars
(integer) 1
127.0.0.1:6379> scard cars # get num members in set
(integer) 3
127.0.0.1:6379> smembers cars # get set members
1) "honda"
2) "bmw"
3) "ford"

# hsets
127.0.0.1:6379> hset andrzej age 33
(integer) 1
127.0.0.1:6379> hset andrzej email dexlab@o2.pl
(integer) 1
127.0.0.1:6379> hget andrzej email
"dexlab@o2.pl"
127.0.0.1:6379> hget andrzej age
"33"
127.0.0.1:6379> hgetall andrzej
1) "age"
2) "33"
3) "email"
4) "dexlab@o2.pl"


# save changes to file
127.0.0.1:6379> save
7:M 17 Nov 2020 17:23:49.765 * DB saved on disk
OK

# delete all keys
127.0.0.1:6379> FLUSHALL
7:M 17 Nov 2020 17:35:00.078 * DB saved on disk
OK

# exit CLI with saving changes
quit
 ```