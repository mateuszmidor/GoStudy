# chat 
This app is a chat server that supports google login and avatars.

# Configure OAuth2 authentication (logging with google)
Setup google OAuth client ID under https://console.developers.google.com/apis/credentials?project=api-project-44918022082
, and put the id and pass in "initOAuth2" function in application.go

# Build
To build it, run: ./build.sh

# Run
To run it, run bin/application and in web browser navigate to http://localhost:8080/chat
When redirected to login page, use the google login method.
This chat uses google account picture for avatar. You can also upload your own avatar under http://localhost:8080/upload