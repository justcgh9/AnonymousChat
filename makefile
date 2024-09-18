ENVFILETEMPLATE=.env.template

ifneq ("$(wildcard $(ENVFILETEMPLATE))","")
	include $(ENVFILETEMPLATE)
	export $(shell sed 's/=.*//' $(ENVFILETEMPLATE))
endif

ENVFILE=.env

ifneq ("$(wildcard $(ENVFILE))","")
	include $(ENVFILE)
	export $(shell sed 's/=.*//' $(ENVFILE))
endif

DB_CONTAINER_NAME=db
create-db:
	docker run --name $(DB_CONTAINER_NAME) \
		-p $(POSTGRES_PORT):5432 \
		-e POSTGRES_USER=$(POSTGRES_USERNAME) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DATABASE) \
		-d postgres

stop-db:
	docker container stop $(DB_CONTAINER_NAME)

rm-db:
	docker container rm $(DB_CONTAINER_NAME)

start-db:
	docker container start $(DB_CONTAINER_NAME)

run:
	go run ./cmd/chat

profile:
	go run ./cmd/profiler

lint:
	golines -w .
	gofmt -w .

recovery:
	python3 recovery.py
	echo "recovered"