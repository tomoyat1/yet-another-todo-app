PKG=github.com/tomoyat1/yet-another-todo-app
.PHONY: all
all: dep-ensure build

.PHONY: build
build:
	go build -o build/todo ${PKG}/cmd/todo-server

.PHONY: dep-ensure
dep-ensure:
	dep version || go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

install: dep-ensure
	go install ${PKG}/cmd/todo-server

.PHONY: docker
docker:
	docker build -t tomoyat1/yet-another-todo-app:latest ./
	docker push tomoyat1/yet-another-todo-app:latest
