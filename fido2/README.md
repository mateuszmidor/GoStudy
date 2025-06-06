# e2e fido2 example

FIDO2 is made of 
- webauthn(Web Authentication API) - protocol between client(browser) and relaying party(server)
- ctap(Client To Authenticator Protocol) - protocol beteween client(browser) and authenticator(e.g. yubikey, TouchID, samsung pass)

CTAP is already implemented by major browser, and out of scope here.  
webauthn consists of registration flow and login flow. Eventually the server receives public key used for checking if the user possesses the right private key

## sources

- https://github.com/go-webauthn/webauthn?
- https://webauthn.guide/
- https://developer.mozilla.org/en-US/docs/Web/API/Web_Authentication_API

## how it works

### user registration

1. frontend asks backend to begin-registration and receives registration options to be used with webauthn
1. frontend uses navigator.credentials.create to trigger webauthn form asking user to select authenticator
1. selected authenticator generates private and public key; private key is stored on device, public key wrapped as Credentials is sent back to fronted, and then to backend
1. backend receives the Credentials and stores them with the User in the database. Credentials is not confidential

### user log in

1. frontend asks backend to begin user login and receives login options to be used with webauthn
1. frontend uses navigator.credentials.get to trigger webauthn form asking user to select authenticator eligible for provided login options
1. selected authenticator signs options.challenge string, which is sent back to frontend and then to backend
1. backend checks with Public Key if the signature was created using associated Private Key

## steps to run the example

webauthn requires HTTPS to even start working, so we must generate keys to run TLS HTTP server:
```sh
go get github.com/jsha/minica # will use minica for generating keys for our little tls web server
go install github.com/jsha/minica
minica --domains localhost # generate the keys for localhost
# manual step: import generated "minica.pem" as firefox trusted CA - the browser need to know it can trust the keys generated by minica CA (Certificate Authority)
go run . 
firefox https://localhost:8888 # see the app in action
 ```