seed:
	docker exec -i go-testing-db psql -U admin go_testing < sql/seed.sql

db:
	docker exec -it go-testing-db psql -d go_testing -U admin
