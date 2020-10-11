# Go Gin Framework Crash Course
## This is a video catalog example implementation  
https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=1

## Requests (Authorization header generated from admin:admin by Postman)
- Add video
``` bash
curl \
-H "Content-Type: application/json" \
-H "Authorization: Basic YWRtaW46YWRtaW4=" \
-d \
'{
    "title": "Orwell 1984",
    "description": "Audiobook by Novel 1984 of G. Orwell",
    "url": "https://www.youtube.com/embed/scqLliarGpM",
    "author" : {
        "firstname": "George",
        "lastname":"Orwell",
        "age": 75,
        "email": "g.orwell@gmail.com"
    }
}' \
-X POST \
localhost:8080/api/videos
```

- Get all videos
```bash
curl -H "Authorization: Basic YWRtaW46YWRtaW4=" -X GET localhost:8080/api/videos
```