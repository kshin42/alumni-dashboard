# alumni-dashboard

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your tests
```
npm run test
```

### Lints and fixes files
```
npm run lint
```

## Devvm Setup
### Download or build Dev VM
docker build -t devvm:latest .

### Run devVM
docker run -d -p 8080:8080 -p 8081:3000 -p 8082:3306 -v [working dir]:/src devvm:latest sleep infinity

### Go into dev vm
docker ps (copy CONTAINER ID)
docker exec -it <CONTAINER_ID> fish

## API Setup
env GOOS=linux go build .

## Mysql setup 
service mysql start
mysql -u root -proot

CREATE DATABASE dev

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';

## Migrations
To create db tables you start the go service and hit localhost:3000/dbsetup
