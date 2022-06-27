build_db:
	sudo podman build . -t mysql_go

run_db:
	sudo podman run  --rm -d -p 3306:3306 --name go-mariadb -e MARIADB_DATABASE=go_mysql -e MARIADB_ROOT_PASSWORD=chrrot  -v /home/lelouch/data_docker/mysql:/var/lib/mysql mysql_go



all_db: build_db run_db
