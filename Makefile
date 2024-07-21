build:
	docker build -t message_service_image ./message-service && docker build -t processed_service_image ./processed
run:
	docker compose down && docker compose up --build -d
initilize:
	chmod ugo+x message-service/tools/initilize.sh && message-service/tools/./initilize.sh
migrate:
	migrate -path message-service/migrations/ -database "postgresql://postgres:test123456@localhost:5436/messages?sslmode=disable" -verbose up
stop:
	docker compose down

