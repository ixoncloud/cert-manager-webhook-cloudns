IMAGE_NAME := "ixoncloud/cert-manager-webhook-cloudns"
IMAGE_TAG := "1.0.3"

OUT := $(shell pwd)/.out

$(shell mkdir -p "$(OUT)")

verify:
	go test -v .

build:
	docker build -t "$(IMAGE_NAME):$(IMAGE_TAG)" .

.PHONY: rendered-manifest.yaml
rendered-manifest.yaml:
	helm template \
	    --name cert-manager-webhook-cloudns \
        --set image.repository=$(IMAGE_NAME) \
        --set image.tag=$(IMAGE_TAG) \
        deploy/cert-manager-webhook-cloudns > "$(OUT)/rendered-manifest.yaml"
