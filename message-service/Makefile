build:
	docker build -t message_service_image .
start:
	docker compose down && docker compose up --build -d
stop:
	docker compose down
images: 
	chmod ugo+x tools/images.sh && tools/./images.sh
swag:
	swag init -g internal/app/app.go -o api/docs