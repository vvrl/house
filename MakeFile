test:
	@echo 'Мы сделали Makefile'

up:
	docker-compose up --build project

stop:
	docker-compose stop