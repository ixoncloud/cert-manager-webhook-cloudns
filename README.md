# Cert-Manager ClouDNS DNS01 Provider

A Cert-Manager DNS01 provider for ClouDNS.

> :warning: **Out of date**: This version has not been updated for up-to-date versions of Kubernetes and cert-manager. Please check out [this fork](https://github.com/mschirrmeister/cert-manager-webhook-cloudns) instead.

## Configuration

Cert-Manager expects DNS01 providers to parse configuration from incoming webhook requests.

This can be used to have multiple Cert-Manager `Issuer` resources use the same instance of the provider with different credentials or configuration.

Because we currently don't need this and the LEGO library already has support for parsing environment variables (and files), we have opted to not use this.

The `testdata/config.json` file is there because the DNS01 provider conformance testing suite wants to mock the requests away, and needs a folder to load the data from.

### Environment Options

|Name|Required|Description|
|---|---|---|
|`GROUP_NAME`|yes|Used to organise cert-manager providers, this is usually a domain|
|`CLOUDNS_AUTH_ID_FILE`|yes|Path to file which contains ClouDNS Auth ID|
|`CLOUDNS_AUTH_PASSWORD_FILE`|yes|Path to file which contains ClouDNS Auth password|
|`CLOUDNS_TTL`|no, default: 60|ClouDNS TTL|
|`CLOUDNS_HTTP_TIMEOUT`|no, default: 30 seconds|ClouDNS API request timeout|

## Development

### Running DNS01 provider conformance testing suite

```bash
# Get kubebuilder
./scripts/fetch-test-binaries.sh

# Run testing suite
TEST_ZONE_NAME=<domain> CLOUDNS_AUTH_ID_FILE=.creds/auth_id CLOUDNS_AUTH_PASSWORD_FILE=.creds/auth_password go test -v .
```
