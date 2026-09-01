IMAGE      := scheduler
PORT       := 7540
DATA_DIR   := $(CURDIR)/data

.PHONY: build run test docker-build docker-run docker-stop

build:
	go build -o scheduler .

run:
	go run main.go

test:
	go test ./tests

docker-build:
	docker build -t $(IMAGE) .

docker-run: docker-build
	mkdir -p $(DATA_DIR)
	-docker rm -f $(IMAGE) 2>/dev/null
	docker run -d \
		--name $(IMAGE) \
		-p $(PORT):7540 \
		--env-file .env \
		-e TODO_DBFILE=/app/data/scheduler.db \
		-v "$(DATA_DIR):/app/data" \
		$(IMAGE)

docker-stop:
	docker stop $(IMAGE)
	docker rm $(IMAGE)
