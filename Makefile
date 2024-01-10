build:
	docker compose -f docker-compose-hot.yml build --no-cache
hot:
	docker compose -p netadim -f docker-compose-hot.yml up --build 
down:
	docker compose -p vnetadim -f docker-compose-hot.yml down
