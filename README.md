openapi-generator-go
====================

Generate go code from an openapi spec!

```bash
# generate structs, enums and a router from a spec file
openapi-generator-go generate --spec ./example/api.yaml --output ./example/generated
# generate models and router independently
openapi-generator-go generate models --spec ./example/api.yaml --output ./example/models --package-name models
openapi-generator-go generate router --spec ./example/api.yaml --output ./example/router --package-name router
```
