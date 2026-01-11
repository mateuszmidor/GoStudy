# go-feature-flag

Go module for working with feature flags:
- The feature flag module automatically re-loads the config file [flags.yaml](./flags.yaml) every 3 seconds.
- Try to change the flag in [flags.yaml](./flags.yaml) and see it updates live!

Note: OpenFeature can use go-feature-flag as it's provider: https://gofeatureflag.org/docs/getting-started
## Run

```sh
go run .

Flag is true
Flag is true
Flag is true
...
```