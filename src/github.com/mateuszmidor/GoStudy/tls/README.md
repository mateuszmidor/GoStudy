# TLS demo

## Steps

- install minica from <https://github.com/jsha/minica>
- generate cert for mydomain.com 
```shell
./minica --domains  mydomain.com
```

, this will get you:
```
mydomain.com
 - key.pem
 - cert.pem
```