# REGO policy examples

* Quick usage reference: https://www.openpolicyagent.org/docs/latest/policy-reference/
* Try your policy here: https://play.openpolicyagent.org

## input

```json
{
    "sites": [
        {
            "region": "east",
            "name": "prod",
            "servers": [
                {
                    "name": "web-0",
                    "hostname": "hydrogen"
                },
                {
                    "name": "web-1",
                    "hostname": "helium"
                },
                {
                    "name": "db-0",
                    "hostname": "lithium"
                }
            ]
        },
        {
            "region": "west",
            "name": "smoke",
            "servers": [
                {
                    "name": "web-1000",
                    "hostname": "beryllium"
                },
                {
                    "name": "web-1001",
                    "hostname": "boron"
                },
                {
                    "name": "db-1000",
                    "hostname": "carbon"
                }
            ]
        },
        {
            "region": "west",
            "name": "dev",
            "servers": [
                {
                    "name": "web-dev",
                    "hostname": "nitrogen"
                },
                {
                    "name": "db-dev",
                    "hostname": "oxygen"
                }
            ]
        }
    ],
    "apps": [
        {
            "name": "web",
            "servers": [
                "web-0",
                "web-1",
                "web-1000",
                "web-1001",
                "web-dev"
            ]
        },
        {
            "name": "mysql",
            "servers": [
                "db-0",
                "db-1000"
            ]
        },
        {
            "name": "mongodb",
            "servers": [
                "db-dev"
            ]
        }
    ],
    "users": [
        {
            "id": "alice",
            "role": "employee",
            "country": "USA"
        },
        {
            "id": "bob",
            "role": "customer",
            "country": "USA"
        },
        {
            "id": "dora",
            "role": "admin",
            "country": "Sweden"
        }
    ]
}
```

## "with" keyword - allows to run other rule with substituted input

```opa
package play

import future.keywords

mateusz_allowed if {
	input.imie == "mateusz"
}  

test_mateusz_allowed if {
	mateusz_allowed with input as {"imie" : "mateusz"}
}  
```
```json
{
    "test_mateusz_allowed": true
}

```
##  For each app, print list of associated hostnames

```opa
package play

import future.keywords

# returns appname:list of hostnames
apps_hostnames[app.name] := hostnames {
	some app in input.apps
	hostnames := [host | some server in input.sites[_].servers; server.name in app.servers; host := server.hostname]
}
```

```json
{
    "apps_hostnames": {
        "mongodb": [
            "oxygen"
        ],
        "mysql": [
            "lithium",
            "carbon"
        ],
        "web": [
            "hydrogen",
            "helium",
            "beryllium",
            "boron",
            "nitrogen"
        ]
    }
}
```