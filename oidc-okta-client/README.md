# OpenID Connect (OIDC) client integration with Okta Identity Provider (IDP)

OAuth2 + OIDC explained: https://www.youtube.com/watch?v=9LRZXg0NK5k  
Client app is basically this: https://github.com/zitadel/oidc/tree/main/example/client

## Steps for Okta:

1. First create an account in Okta https://developer.okta.com/signup/ ("Access the Okta Developer Edition Service" tile).
1. Then after logging-in, `Applications` -> `Applications`  -> `Create App Integration` 
1. Select `OIDC` and then `Web Application`
1. Set `Sign-in redirect URIs` to `http://localhost:8000/auth/callback` (endpoint served by our client app)
1. Set `Controlled access` to `Allow everyone in your organization to access` (so you don't have to assign permitted user to the app)
1. Save
1. For "groups" claim to be returned in id_token, in the app config page on tab `Sign On`, set `Groups claim filter` regex to `.*`  
See: https://developer.okta.com/docs/guides/customize-tokens-groups-claim/main/#add-a-groups-claim-for-the-org-authorization-server

1. Obtain CLIENT_ID (like: 0oa3v9n54yf4zbKJ85d7)
1. Obtain CLIENT_SECRET (like: v4M04a3PLvkvvInaz6o1Q-jzWTPR5mGRITpkqSKo), or use PEM private and public keys - advanced
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
1. Or Revoke Refresh Token endpoint, so user can no longer refresh: `localhost:8000/auth/revoke` (11.2023: EntraID doesn't support revocation)

## Steps for AzureAD (new name: EntraID)

1. Login to your Azure cloud
1. Then after logging-in, `Microsoft Entra ID` -> `App registrations` -> `New registration`
1. Input `Name`
1. Select `Supported account types` = `Accounts in this organizational directory only (Default Directory only - Single tenant)`
1. Select `Redirect URI (optional)`: `Platform` = `Web`, input URL = `http://localhost:8000/auth/callback`
1. Click `Register`
1. Your ISSUER is `"https://login.microsoftonline.com/<Directory (tenant) ID>/v2.0"`, the CLIENT_ID is `<Application (client) ID>`, eg:
    * ISSUER: `"https://login.microsoftonline.com/e0ae040b-2d16-41ad-bd29-faaa3ec975b9/v2.0"`
    * CLIENT_ID: `c86f603c-c425-4832-bc4a-906f492ac77f`
1. Go to `Certificates & Secrets` -> `New client secret`, click `Add` and save the `Value` field: this is your CLIENT_SECRET
1. Go to `Token configuration`, then:
    * `Add optional claim` -> `Token type` = ID, select `family_name` and `given_name`, so that Entra returns these values in OIDC token, click `Add` and agree to `Turn on the Microsoft Graph profile permission (required for claims to appear in token)`
    * `Add groups claim`, select relevant groups and click `Add`
1. Go to `API permission`, you should see: `profile` and `User.Read` (both `delegated`)
1. On first login, you will be asked by Microsoft to allow access for this application ("Żądane uprawnienia") 

## OAuth2 client app credentials

- clientID + clientSecret - just another names for user and password, registered in IDP to authorize the client application
- clientID + rsaPrivKey + rsaPrivKeyId - authentication with x.509 instead of clientSecret, more secure, public key registered in IDP

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

## openid-configuration endpoints

* Okta: https://dev-39423526.okta.com/.well-known/openid-configuration
* EntraID: https://login.microsoftonline.com/c2ceea60-3945-479c-b6a6-15be438c0c4b/v2.0/.well-known/openid-configuration

## Bazel

Bazel setup is here for educational purposes only; you can ignore it.  
* `bazel run //:gazelle` - generate BUILD.bazel files with go targets
* `bazel run //:gazelle-update-repos` - generate deps.bzl file with go dependencies
* `source .env && bazel run //app:app` - run the app