include .env
export 

service-run:
	@go run main.go

service-deploy:
	@sudo docker compose up -d application

service-undeploy:
	@sudo docker compose down application