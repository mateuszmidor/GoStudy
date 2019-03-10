# answerapp
Web app like stack overflow: ask questions, get answers, rate anwers.
Run it on google appengine

## Install gcloud for GO:
    https://cloud.google.com/appengine/docs/standard/go/download?hl=pl

## Install appengine go components
    go get google.golang.org/appengine/cmd/aefix
    go get github.com/golang/appengine
    go get github.com/golang/protobuf/proto
    
## Run API:
    cd api
    python2 /home/mateusz/bin/google-cloud-sdk/bin/dev_appserver.py .
    check localhost:8080

# Deploy on gcloud
- Create new project on google cloud: https://console.cloud.google.com/  
- Deploy the app: gcloud app deploy