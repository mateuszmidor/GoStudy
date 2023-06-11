# OpenID Connect (OIDC) client integration with Okta Identity Provider (IDP)

Client app is basically this: https://github.com/zitadel/oidc/tree/main/example/client.

Steps:
1. First create an account in Okta https://developer.okta.com/signup/ ("Access the Okta Developer Edition Service" tile).
1. Then after logging-in, `Applications` -> `Applications`  -> `Create App Integration` 
1. Select `OIDC` and then `Web Application`
1. Set `Sign-in redirect URIs` to `http://localhost:9999/auth/callback` (endpoint served by our client app)
1. Set `Controlled access` to `Allow everyone in your organization to access`
1. Save
1. Obtain CLIENT_ID (like: 0oa3v9n54yf4zbKJ85d7)
1. Obtain CLIENT_SECRET (like: v4M04a3PLvkvvInaz6o1Q-jzWTPR5mGRITpkqSKo)
1. Obtain ISSUER (this is your subdomain in Okta, like: https://dev-39423526.okta.com)
1. Export CLIENT_ID, CLIENT_SECRET, ISSUER as env variables - do it in [.env](./.env) file
1. `make run`
1. Browse `localhost:9999/login` and login with your user/password used to register in Okta
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
        locale: en_US
        middle_name: TheGuy
        name: Peter Griffin
        preferred_username: p.griffin@familyguy.com
        sub: 00u9w8nj1qLtTeilf5d7
        zoneinfo: America/Los_Angeles
    ```

## Bazel

Bazel setup is here for educational purposes only; you can ignore it.  
* `source .env && bazel run //app:app` - run the app