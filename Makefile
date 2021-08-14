docker-compose/up:
	docker-compose up -d --build
docker-compose/down:
	docker-compose down -v
docker-compose/ps:
	docker-compose ps
docker-compose/logsf:
	docker-compose logs -f
app/start:
	docker exec sqlx-pg-sample_app_1 go run main.go