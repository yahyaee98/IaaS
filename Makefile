APP := iaas
CI_COMMIT_REF_SLUG?=$(shell cat .git/HEAD | cut -d '/' -f 3)

compile:
	go build -race -o bin/$(APP) .

run: compile
	bin/$(APP)

test:
	go test -race ./...

docker:
	docker build \
          -t $(APP):$(CI_COMMIT_REF_SLUG) .

check-golint:
	which golint || (go get -u golang.org/x/lint/golint)

lint: check-golint
	find . -type f -name "*.go" | xargs -n 1 -I R golint -set_exit_status R
