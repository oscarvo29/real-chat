BACKEND_BINARY=backendBinary

## up: starts all the containers in the background without forcing build
up:
	@echo "starting containers"
	docker compose up -d

up_build: build_backend
	@echo "Stopping containers if running...."
	docker compose down 
	@echo "Building containers..."
	docker compose up --build -d 
	@echo "Docker images build!"

down:
	@echo "Shutting down containers.."
	docker compose down
	@echo "Docker shutdown..."

build_backend:
	@echo "Building backend..."
	cd ../real-chat-backend && env GOOS=linux CGO_ENABLED=0 go build -o ${BACKEND_BINARY} .
	@echo "Build done"