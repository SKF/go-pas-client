.PHONY: clean internal/models
.INTERMEDIATE: swagger.json

RM     ?= rm
WGET   ?= wget
MKDIR  ?= mkdir
DOCKER ?= docker

API_URL := "https://api.point-alarm-status.sandbox.iot.enlight.skf.com/v1"
GOSWAGGER_VERSION := v0.29.0

swagger.json:
	$(WGET) "$(API_URL)/docs/swagger/doc.json" -O "$@"

internal/models: swagger.json
	$(RM) -rf "$@" && $(MKDIR) -p "$@"
	$(DOCKER) run --rm \
		--volume "$(shell pwd):/src" \
		--workdir /src \
		--user "$(shell id -u):$(shell id -g)" \
		quay.io/goswagger/swagger:${GOSWAGGER_VERSION} \
			generate model --spec="$<" --target="$(@D)"

clean:
	$(RM) -rf models
