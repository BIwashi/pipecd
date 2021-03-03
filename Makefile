BAZEL_FLAGS =
BAZEL_COMMAND_FLAGS =
CLOUDSDK_PYTHON = "/usr/bin/python"

ifdef EXTENDED_BAZEL_RC
	BAZEL_FLAGS += --bazelrc=${EXTENDED_BAZEL_RC}
endif

ifdef IS_CI
	BAZEL_FLAGS += --output_base=/workspace/bazel_out
	BAZEL_COMMAND_FLAGS += --config=ci
endif

ifdef BUILD_PLATFORM
	BAZEL_COMMAND_FLAGS += --config=${BUILD_PLATFORM}
endif

ifdef PYTHON_PATH
	CLOUDSDK_PYTHON=${PYTHON_PATH}
endif

.PHONY: build
build:
	bazelisk ${BAZEL_FLAGS} build ${BAZEL_COMMAND_FLAGS} -- //...

.PHONY: test
test:
	bazelisk ${BAZEL_FLAGS} test ${BAZEL_COMMAND_FLAGS} -- //pkg/...

.PHONY: test-debug
test-debug:
	bazelisk ${BAZEL_FLAGS} test ${BAZEL_COMMAND_FLAGS} --test_output=all -- //pkg/...

.PHONY: test-integration
test-integration:
	bazelisk ${BAZEL_FLAGS} test ${BAZEL_COMMAND_FLAGS} --action_env=CLOUDSDK_PYTHON=${CLOUDSDK_PYTHON} -- //test/integration/...

.PHONY: coverage
coverage:
	bazelisk ${BAZEL_FLAGS} coverage ${BAZEL_COMMAND_FLAGS} //pkg/...

.PHONY: dep
dep:
	go mod tidy
	go mod vendor
	bazelisk run //:gazelle -- update-repos -from_file=go.mod -prune -to_macro=repositories.bzl%go_repositories

.PHONY: gazelle
gazelle:
	bazelisk run //:gazelle

.PHONY: buildifier
buildifier:
	bazelisk run //:buildifier

.PHONY: clean
clean:
	bazelisk clean --expunge

.PHONY: expose-generated-go
expose-generated-go:
	./hack/expose-generated-go.sh pipe-cd pipe

.PHONY: site
site:
	hugo server --source=docs

.PHONY: web-dep
web-dep:
	bazelisk build //pkg/app/web:build_api //pkg/app/web:build_model

.PHONY: web-dev
web-dev:
	cd pkg/app/web; yarn dev

.PHONY: generate-test-tls
generate-test-tls:
	openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
		-keyout pkg/rpc/testdata/tls.key \
		-out pkg/rpc/testdata/tls.crt \
		-subj "/CN=localhost" \
		-config pkg/rpc/testdata/tls.config
