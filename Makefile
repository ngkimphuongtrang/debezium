up:
	@docker compose -f docker/docker-compose.yml up

down:
	@docker compose -f docker/docker-compose.yml down -v --rmi local