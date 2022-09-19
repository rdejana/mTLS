# Refs

https://letsencrypt.org/docs/certificates-for-localhost/

https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go

## Notes
Cert/key require 'clientAuth' if used for client outh.  

e.g.
`extendedKeyUsage=serverAuth,clientAuth`


```
 openssl req -x509 -out localhost.crt -keyout localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth,clientAuth")
```
