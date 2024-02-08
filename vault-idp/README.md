# vault-idp

This example shows how to use Vault as Identity Provider (IDP) with OIDC (OpenID Connect) support.
- SSH: https://brian-candler.medium.com/using-hashicorp-vault-as-an-ssh-certificate-authority-14d713673c9a
- OIDC: https://brian-candler.medium.com/using-vault-as-an-openid-connect-identity-provider-ee0aaef2bba2

## Run Vault with `userpass` engine enabled

```bash
make vaut # run vault container and enable userpass secret engine - for user/pass authentication.
firefox localhost:8200 # login with token: root-token
make stop # kill vault container
```

## Programatically setup Vault as IDP

```sh
make idp
```
, output:
```text
ClientID: uHJ0PvHMnRJbguQG1Wohm9Cy2MeKpdCR
ClientSecret: hvo_secret_Nj1dvY3d7SoAqrwhwogbU9iVxMuuAzTAPccnp2gNYJ1Dtj48b3bi7dkcZz9bPPAR
IssuerURL: http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider
```

## Manually setup Vault as IDP


Connect to vault container, set server URL to http as we are running vault in http mode:
```sh
docker exec -e 'VAULT_ADDR=http://127.0.0.1:8200' -it my-vault sh
```

create user:
```sh
vault login root-token # need root credentials to configure vault
# vault auth enable userpass # enable userpass secret engine in vault; already enabled with make command
vault write auth/userpass/users/user password=pass # create new user
vault list auth/userpass/users
#vault login -method=userpass username=user password=pass # can already login as "user"
vault auth list
```

add custom scope for OIDC purposes, this will expose user data to OIDC client apps:
(https://developer.hashicorp.com/vault/docs/concepts/policies#templated-policies):
```sh
vault auth list
TOKEN_TEMPLATE=$(cat << EOF
{
    "username": {{identity.entity.name}},
    "contact": {
        "email": {{identity.entity.metadata.email}},
        "phone_number": {{identity.entity.metadata.phone_number}}
    },
    "groups": {{identity.entity.groups.names}}
}
EOF
)
vault write identity/oidc/scope/user \
  description="Scope for user metadata" \
  template="$(echo $TOKEN_TEMPLATE | base64 -)"

TOKEN_TEMPLATE=$(cat << EOF
{
    "userinfoprofile": {
        "givenname": "mateusz",
        "familyname": "midor"
    }
}
EOF
)
vault write identity/oidc/scope/userinfoprofile \
  description="Scope for user metadata" \
  template="$(echo $TOKEN_TEMPLATE | base64 -)"
```

create OIDC provider:
```sh
vault write identity/oidc/provider/my-provider \
  allowed_client_ids="*" \
  scopes_supported="user"
vault read identity/oidc/provider/my-provider
```

create OIDC client app:
```sh
vault write identity/oidc/client/my-webapp \
  redirect_uris="http://localhost:8000/auth/callback" \
  assignments="allow_all" \
  key="default" \
  id_token_ttl="300m" \
  access_token_ttl="10h"
vault read identity/oidc/client/my-webapp
```

```sh
vault write identity/oidc/client/my-webapp \
  redirect_uris="https://id.local.mydatamarx.com:3000" \
  assignments="allow_all" \
  key="default" \
  id_token_ttl="300m" \
  access_token_ttl="10h"
vault read identity/oidc/client/my-webapp
```

**DONE**

Issuer URL:
```sh
vault read identity/oidc/provider/my-provider
```
```text
Key                   Value
---                   -----
allowed_client_ids    [*]
issuer                http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider
scopes_supported      []
```

ClientID+ClientSecret:
```sh
vault read identity/oidc/client/my-webapp
```
```text
Key                 Value
---                 -----
access_token_ttl    10h
assignments         [allow_all]
client_id           M64R1BLVBWHneIHANIINAf8KS91pg4RN
client_secret       hvo_secret_fQjiOG0TJKuOy9bpdlDyCir10QMGRViWj1DU7Acskd9Xeb42hvZAkHgjMOtSp6E2
client_type         confidential
id_token_ttl        5h
key                 default
redirect_uris       [http://localhost:8000/auth/callback]
```

## Provider config

```sh
curl -Ss http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider/.well-known/openid-configuration | jq
```
, output:
```json
{
  "issuer": "http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider",
  "jwks_uri": "http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider/.well-known/keys",
  "authorization_endpoint": "http://127.0.0.1:8200/ui/vault/identity/oidc/provider/my-provider/authorize",
  "token_endpoint": "http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider/token",
  "userinfo_endpoint": "http://127.0.0.1:8200/v1/identity/oidc/provider/my-provider/userinfo",
  "request_parameter_supported": false,
  "request_uri_parameter_supported": false,
  "id_token_signing_alg_values_supported": [
    "RS256",
    "RS384",
    "RS512",
    "ES256",
    "ES384",
    "ES512",
    "EdDSA"
  ],
  "response_types_supported": [
    "code"
  ],
  "scopes_supported": [
    "user",
    "openid"
  ],
  "claims_supported": [],
  "subject_types_supported": [
    "public"
  ],
  "grant_types_supported": [
    "authorization_code"
  ],
  "token_endpoint_auth_methods_supported": [
    "none",
    "client_secret_basic",
    "client_secret_post"
  ]
}
```
