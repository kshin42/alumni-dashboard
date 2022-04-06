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

## Mysql setup 
service mysql start
mysql -u root -proot

## Migrations
CREATE DATABASE dev

CREATE TABLE organizations
(id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
 name VARCHAR(256),
 orgCode VARCHAR(256)
 description VARCHAR(256),
 timeCreated TIMESTAMP DEFAULT now()
)

CREATE TABLE logins 
(id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
 email VARCHAR(256), 
 firstName VARCHAR(256),
 lastName VARCHAR(256),
 passwordHash VARCHAR(256)
 timeCreated TIMESTAMP DEFAULT now(),
 timeUpdated TIMESTAMP);

INSERT into logins set (email, firstName, lastName, passwordHash)
VALUES ('test@test.com', 'Ftest', 'Ltest', 'zp8Th8LLoyvPhI2Jo2wZ/w$BHmCXAh5zB15LBnwyP6LsqGzdnhyMsbC87hR8MGIIwU');


	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Metadata     string