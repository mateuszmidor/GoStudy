# Go Gin Framework Crash Course (with GORM and JWT and BasicAuth)
## This is a video catalog example  
https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=1

## API
- Basic Auth test (credentials in base64 - admin:pass)
```bash
curl -H "Authorization: Basic YWRtaW46cGFzcw==" -X GET localhost:8080/batest
```

- Login (generates JWT)
```bash
curl -X POST 'localhost:8080/login?username=admin&password=pass'
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiYWRtaW4iOnRydWUsImV4cCI6MTYwMjc1MjQxNSwiaWF0IjoxNjAyNDkzMjE1LCJpc3MiOiJtYXRldXN6bWlkb3IuY29tIn0.EuoHF1zVYkMvfjLD58BJFOYVXnh6EsaLb5RMwkhXTwM"}
```

- Get all videos (requires JWT)
```bash
curl \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiYWRtaW4iOnRydWUsImV4cCI6MTYwMjc1MjQxNSwiaWF0IjoxNjAyNDkzMjE1LCJpc3MiOiJtYXRldXN6bWlkb3IuY29tIn0.EuoHF1zVYkMvfjLD58BJFOYVXnh6EsaLb5RMwkhXTwM" \
-X GET localhost:8080/api/videos
```

- Add video (requires JWT)
``` bash
curl \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiYWRtaW4iOnRydWUsImV4cCI6MTYwMjc1MjQxNSwiaWF0IjoxNjAyNDkzMjE1LCJpc3MiOiJtYXRldXN6bWlkb3IuY29tIn0.EuoHF1zVYkMvfjLD58BJFOYVXnh6EsaLb5RMwkhXTwM" \
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

## HTML view
http://localhost:8080/view/videos
