run_db:
	sudo podman run --rm -d -p 3306:3306 --name go-mariadb -e MARIADB_DATABASE=mariadb -e MARIADB_ROOT_PASSWORD=chrrot  -v /home/lelouch/data_docker/mysql:/var/lib/mysql docker.io/library/mariadb:10.6.8-focal
test:
	go test gitlab.com/eliotandelon/gotesting/controllers
	go test gitlab.com/eliotandelon/gotesting/implementation
