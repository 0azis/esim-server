build:
	docker compose up --build -d

down:
	docker compose down --rmi all

server-run:
	docker compose build http
	docker compose up -d http
