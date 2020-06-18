APP := iaas
CI_COMMIT_REF_SLUG?=$(shell cat .git/HEAD | cut -d '/' -f 3)

compile:
	go build -race -o bin/$(APP) cmd/$(APP)/main.go

run: compile
	bin/$(APP)

docker:
	docker build \
          -t $(APP):$(CI_COMMIT_REF_SLUG) .

