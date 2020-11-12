openapi-generator-go
====================

Generate go code from an openapi spec!

```bash
# generate structs, enums and a router from a spec file
openapi-generator-go generate --spec api.yaml --output ./openapi --package-name openapi
# generate models and router independently
openapi-generator-go generate models --spec api.yaml --output ./models --package-name models
openapi-generator-go generate router --spec api.yaml --output ./router --package-name router
```
