module github.com/ixoncloud/cert-manager-webhook-cloudns

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-acme/lego v2.7.1+incompatible
	github.com/jetstack/cert-manager v0.9.0
	github.com/stretchr/testify v1.3.0
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190413052642-108c485f896e
