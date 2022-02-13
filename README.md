#Go Helper v0.0.1

To install Go Helper package, you need to install Go and set your Go workspace first. 
```sh
$ go get -u github.com/MuhammadSuryono/go-helper
```

Add connection database on your `.env` file and change by your config database
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=homestead
DB_PASS=homestead
DB_NAME=golang_db
DB_DRIVER=mysql
```
Support connection to database
```
- MySql (mysql)
- Postgre (postgres)
- SQL Server (sql-server)
```
