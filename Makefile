APP_NAME=godest

.PHONY: run test fmt vet build docker-build k8s-dev k8s-prod

run:
	go run ./cmd

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build -o bin/$(APP_NAME) ./cmd

docker-build:
	docker build -f deploy/docker/Dockerfile -t $(APP_NAME):latest .

k8s-dev:
	kubectl apply -k deploy/k8s/overlays/dev

k8s-prod:
	kubectl apply -k deploy/k8s/overlays/prod
