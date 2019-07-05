# INVENTORY SERVICE


## Requirements
1. Instal go
2. [Install soda](https://gobuffalo.io/en/docs/db/toolbox/) (to create database and running migration)


## Running Service
- `go run main.go`

## Database
1. soda create -e development (to create database development) [more](https://gobuffalo.io/en/docs/db/toolbox/)
2. soda drop -e development (to drop or delete database) [more](https://gobuffalo.io/en/docs/db/toolbox/)
3. migration : `soda migrate -p database up` [more](https://gobuffalo.io/en/docs/db/migrations/)


