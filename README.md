## Teravision
Test project.
The project runs by default on port 1323

### Configuration
You must create or edit the env.yml file with the following content.
```yaml
host:
port: 1323
driver_name: postgres
database_dsn: 
```
Example:
```yaml
host:
port: 1323
driver_name: postgres
database_dsn: postgres://teravision:teravision@172.17.0.1/teravision
```

### Dependencies

The first step is to execute

```
go get -v
```

### Run
```bash
go run main.go
```
Or
```bash
sh build.sh && docker-compose up
```

### Build docker image
```bash
sh build.sh
```
or
```bash
go build -o bin/teravision
```

### Endpoint
To **insert** a new item you must make the following request
```http
@host = http://127.0.0.1:1323

PUT {{host}}/api/users
Content-Type: application/json

{
  "name": "julio torres",
  "dni": "12826254",
  "date": "1990-04-30"
}
```
in curl
```bash
host="http://127.0.0.1:1323"

curl -X PUT $host/api/users -H "Content-Type: application/json" \
--data '{"name": "julio torres","dni": "12826254","date":"1990-04-30"}'
```

### Note
* You must edit the postgresql.conf postgres pg_hba.conf files to allow traffic from the container.
* The database must exist, however, the tables are generated automatically when you start the project.
