# OpenID Connect (OIDC) client integration with Okta Identity Provider (IDP)

OAuth2 + OIDC explained: https://www.youtube.com/watch?v=9LRZXg0NK5k  
Client app is basically this: https://github.com/zitadel/oidc/tree/main/example/client

## Steps:
1. First create an account in Okta https://developer.okta.com/signup/ ("Access the Okta Developer Edition Service" tile).
1. Then after logging-in, `Applications` -> `Applications`  -> `Create App Integration` 
1. Select `OIDC` and then `Web Application`
1. Set `Sign-in redirect URIs` to `http://localhost:8000/auth/callback` (endpoint served by our client app)
1. Set `Controlled access` to `Allow everyone in your organization to access` (so you don't have to assign permitted user to the app)
1. Save
1. For "groups" claim to be returned in id_token, in the app config page on tab `Sign On`, set `Groups claim filter` regex to `.*`  
See: https://developer.okta.com/docs/guides/customize-tokens-groups-claim/main/#add-a-groups-claim-for-the-org-authorization-server

1. Obtain CLIENT_ID (like: 0oa3v9n54yf4zbKJ85d7)
1. Obtain CLIENT_SECRET (like: v4M04a3PLvkvvInaz6o1Q-jzWTPR5mGRITpkqSKo)
1. Obtain ISSUER (this is your subdomain in Okta, like: https://dev-39423526.okta.com)
1. Export CLIENT_ID, CLIENT_SECRET, ISSUER as env variables - do it in [.env](./.env) file
1. `make run`
1. Browse `localhost:8000/` and login with your user/password used to register in Okta
1. Observe the returned token contents, like:
    ```yaml
    subject: 00u9w8nj1qLtTeilf5d7
    userinfoprofile:
        name: Peter Griffin
        givenname: Peter
        familyname: Griffin
        middlename: TheGuy
        nickname: ""
        profile: ""
        picture: ""
        website: ""
        gender: ""
        birthdate: ""
        zoneinfo: America/Los_Angeles
        locale: {}
        updatedat: 0
        preferredusername: p.griffin@familyguy.com
    userinfoemail:
        email: p.griffin@familyguy.com
        emailverified: true
    userinfophone:
        phonenumber: ""
        phonenumberverified: false
    address: null
    claims:
        email: p.griffin@familyguy.com
        email_verified: true
        family_name: Griffin
        given_name: Peter
        groups:
          - viewer
          - editor
        locale: en_US
        middle_name: TheGuy
        name: Peter Griffin
        preferred_username: p.griffin@familyguy.com
        sub: 00u9w8nj1qLtTeilf5d7
        zoneinfo: America/Los_Angeles
    ```
1. You can also test the Refresh Token endpoint: `localhost:8000/auth/refresh`

## OAuth2 client app credentials

clientID + clientSecret - just another names for user and password, registered in IDP to authorize the client application

## OAuth2 client types

* confidential - client app runs on web server and can safely store clientSecret
* public - Single Page Application, android native, etc - runs on user device, can be infected by malicius software, can't store clientSecret

## OAuth2 grant types

* code grant - for confidential client apps; use clientSecret and clientID to get code grant, then exchange it for authorization token
* code grant + PKCE - for public client apps; instead of using clientSecret, client app uses code_verifier and code_challenge
* device grant - for public client apps without GUI (no browser to show login page), user gets login URL and opens login page on a separate device
* client credentials - for confidential client apps that authorize themself directly with its secret without the user even knowing (triggered without user action).

## OIDC

* builds on top of OAuth2
* can use code grant authorization for web server confidential apps
* client requests additional scopes like profile email phone address
* IDP returns additional info for requested scopes in id_token (JWT format), attached in addition to auth_token. It may be "thin" token - with only basic user info
* IDP exposes additional endpoint `/userinfo` for fetching full and most recent information about the user; requires authorization with auth_token

## Bazel

Bazel setup is here for educational purposes only; you can ignore it.  
* `bazel run //:gazelle` - generate BUILD.bazel files with go targets
* `bazel run //:gazelle-update-repos` - generate deps.bzl file with go dependencies
* `source .env && bazel run //app:app` - run the app