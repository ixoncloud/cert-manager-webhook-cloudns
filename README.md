# Cert-Manager CloudDNS support

Application which runs a Cert-Manager DNS01 challenge webhook server.

## Environment

|Name|Required|Description|
|---|---|---|
|`GROUP_NAME`|yes|Used to organise cert-manager webhooks by name|
|`CLOUDNS_AUTH_ID`|yes|ClouDNS Auth ID|
|`CLOUDNS_AUTH_PASSWORD`|yes|ClouDNS Auth password|
|`CLOUDNS_TTL`|no, default: 60|ClouDNS TTL|
|`CLOUDNS_HTTP_TIMEOUT`|no, default: 30 seconds|ClouDNS API request timeout|
