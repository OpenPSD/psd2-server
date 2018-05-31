# OpenPSD PSD2 Server

## code generation 

Basic server implementation is done by code generation based on the a swagger 2.0 specification. The swagger 2.0 spec is derived from the original openapi 3.0 spec.

https://goswagger.io/generate/server.html

```
swagger generate server --name=openpsd --copyright-file=LICENSEHEADER --spec=psd2_swagger2.yaml
swagger generate client --name=openpsd --copyright-file=LICENSEHEADER --spec=psd2_swagger2.yaml
go get -u -f ./...
```