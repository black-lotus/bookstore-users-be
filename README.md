## RUN DEVELOPMET MODE
```
$ mysql_username=root mysql_password=root mysql_host=127.0.0.1 mysql_database=go_users go run main.go
```

## DATABASE SETUP
Create database & users tables

```
CREATE TABLE `go_users`.`users` ( `id` BIGINT NOT NULL AUTO_INCREMENT , `first_name` VARCHAR(16) NOT NULL , `last_name` VARCHAR(16) NOT NULL , `email` VARCHAR(32) NOT NULL , `date_created` VARCHAR(32) NOT NULL , `status` VARCHAR(16) NOT NULL , `password` VARCHAR(32) NOT NULL , PRIMARY KEY (`id`), UNIQUE `go_users_email_unique` (`email`)) ENGINE = InnoDB;
```