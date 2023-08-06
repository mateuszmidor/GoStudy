# System for Cross-domain Identity Management (SCIM 2.0) server integration with Okta Identity Provider (IDP)
  
SCIM + Okta explained: https://www.youtube.com/watch?v=JmA83cy0uVc  
Server app is basicaly this: https://github.com/elimity-com/scim

## Highlights

* SCIM works in **push** manner: 
  * Okta pushes users and groups to your application
  * your application can't pull users and groups from Okta
* your application runs actually an HTTP server that serves SCIM protocol requests (CRUD) at designated endpoint
  * this means, for local development you must expose your machine to the internet
  * we use `localtunnel` for that - provides static dns name for free
* Okta authenticates in your application using User/Pasword or Bearer Token in request's header
* while configuring your application in Okta, Okta will send a connection test request to `/Users` endpoint, so your app must implement Users even if it only needs Groups

## Steps

As Okta only supports SCIM as part of SAML 2.0 integration, we unfortunatelly need to setup a minimal SAML 2.0 integration first:

1. First create an account in Okta https://developer.okta.com/signup/ ("Access the Okta Developer Edition Service" tile).
1. Then after logging-in, `Applications` -> `Applications`  -> `Create App Integration` 
1. Select `SAML2.0`
1. Next  
1. Input `http://www.example.com` for `Single sign-on URL` - it is required but we won't use it
1. Input `abc` for `Audience URI (SP Entity ID)` - it is required but we won't use it
1. Next
1. Select `I'm an Okta customer adding an internal app`
1. Finish
1. In `General` tab, check `Enable SCIM provisioning`
1. In `Provisioning` tab -> `Integration`:
    * `SCIM connector base URL `, input `https://dob-mateusz-scim.loca.lt/scim`
    * `Unique identifier field for users`, input `userName`
    * Select all `Supported provisioning actions` (maybe, unselect Import Groups, if error on Save)
    * Configure Basic Auth with User/Password
    * Save
1. (when your SCIM server is already running): go to `Push Groups` tab, and push some groups!

## Run

```bash
make 
firefox localhost:32000/scim/v2/Users
```

```json
{
  "Resources": [
    {
      "id": "sample-user",
      "userName": "John Doe",
      "active": true,
      "name": {
        "familyName": "Doe",
        "givenName": "John"
      },
      "emails": [
        {
          "value": "john_doe@gmail.com"
        }
      ],
      "meta": {
        "resourceType": "User",
        "location": "Users/sample-user"
      },
      "schemas": [
        "urn:ietf:params:scim:schemas:core:2.0:User"
      ]
    }
  ],
  "itemsPerPage": 100,
  "schemas": [
    "urn:ietf:params:scim:api:messages:2.0:ListResponse"
  ],
  "startIndex": 1,
  "totalResults": 1
}
```

## Validate 

* the SCIM server implementation can be validated against Okta specification: https://developer.okta.com/docs/guides/scim-provisioning-integration-prepare/main/#customize-the-imported-runscope-test-for-your-scim-integration
* this scim-okta-server has resource filtering implemented and passes the test 100%

## Example pagination

pagination may be used for /scim/Users? and /scim/Groups?:
* startIndex=10&count=5

## Example filters

filter may be used for /scim/Users? and /scim/Groups?:
* filter=userName eq "John Doe" # equal
* filter=not(userName eq "John Doe") # not equal
* filter=name.familyName co "Do" # contains
* filetr=userName eq "John Doe" or userName eq "Johny Bravo" # alternative
* filter=emails[value co "gmail"] # any of the configured emails matches

## Example query

* curl 'http://localhost:32000/scim/Users?startIndex=2&count=1&filter=userName%20eq%20%22John%20Doe%22'

## TODO

Query string params:
* sortBy
* orderBy
* attributes