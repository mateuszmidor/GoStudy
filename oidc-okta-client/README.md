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
1. Obtain CLIENT_SECRET (like v4M04a3PLvkvvInaz6o1Q-jzWTPR5mGRITpkqSKo)
1. Obtain ISSUER (this is your subdomain in Okta, like https://dev-39423526.okta.com)
1. Export CLIENT_ID, CLIENT_SECRET, ISSUER as env variables (you can store them in [.env](./.env) file
1. `make run`
1. Browse `localhost:9999/login` and login with your user/password used to register in Okta
1. Observe profile data like:
    ```json
    {
    "email":"m.midor@devopsbay.com","email_verified":true,"family_name":"Midor","given_name":"Mateusz","locale":"en-US","name":"Mateusz Midor","preferred_username":"m.midor@devopsbay.com","sub":"00u9v93yfoI38P9fv5d7","zoneinfo":"America/Los_Angeles"
    }
    ```