# Config Gorm

### To use the settings of this project you must perform the following instructions:

## create the .env file:
```sh
mkdir .env
```

## specifications:
```sh
host=<host>
port=<port>
database=<database-name>
user=<user-name>
password=<your-password>
charset=utf8
sql_log=true
time_zone=""

MaxIdleConns=10 
MaxOpenConns=100 
```

## the above specifications set the values of the following structure:

```go
type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
	print_log    bool
}
```