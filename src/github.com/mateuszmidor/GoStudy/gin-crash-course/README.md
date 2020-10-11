# Go Gin Framework Crash Course
## This is a video catalog example implementation  
https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=1

## Requests
- Add video
``` bash
curl \
-H "Content-Type: application/json" \
-d \
'{
    "title": "1984",
    "description": "Audiobook by G. Orwell",
    "url": "https://www.youtube.com/watch?v=scqLliarGpM"
}' \
-X POST \
localhost:8080/videos
```

- Get all videos
```bash
curl -X GET localhost:8080/videos
```