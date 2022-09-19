# Refs

https://letsencrypt.org/docs/certificates-for-localhost/

https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go

## Notes
Cert/key require 'clientAuth' if used for client outh.  

e.g.
`extendedKeyUsage=serverAuth,clientAuth`
