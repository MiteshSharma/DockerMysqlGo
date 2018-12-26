# DockerMysqlGo

Running mysql docker container using go code

### Steps to run mysql docker container using command line:

Step 1: Check docker is installed by running command "docker ps". If docker is not installed, install docker from https://docs.docker.com/install/.

Step 2: Start docker mysql container with user name "gouser", password "gopassword" and database name "godb" using command:

```
docker run --name our-mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=gouser -e MYSQL_PASSWORD=gopassword -e MYSQL_DATABASE=godb -p 3306:3306 --tmpfs /var/lib/mysql mysql:5.7
```

Step 3: Check docker status using command with container name our-mysql-container: 

```
docker ps -a
```

Step 4: Once docker is up and running, connect using connection string:

```
gouser:gopassword@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local
```

### Docker using go

We are using same steps to run docker mysql container using go.

docker.go : This file contains generic code to run any type of database using docker.

mysql.go : This file runs mysql container using docker.go
