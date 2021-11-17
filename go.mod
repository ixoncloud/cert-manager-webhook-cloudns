module github.com/ixoncloud/cert-manager-webhook-cloudns

go 1.17

require (
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/go-acme/lego v2.7.2+incompatible
	github.com/jetstack/cert-manager v1.6.1
	github.com/miekg/dns v1.1.43 // indirect
	github.com/stretchr/testify v1.7.0
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	k8s.io/client-go v0.22.3
)
