<div align="center">
  <h1>Restaurant Service</h1>
  <p>
    It is a restaurant backend server. Here I try to use Twelve-Factor App Methodology. 
  </p>
  <p>Submitted By: <a href="https://github.com/IstiyakRiyad" target="_blank">Md. Istiyak Hossain</a> </p>
</div>

## Prerequisite
For manually build the server, install golang & postgresql locally. Setup the configuration file. Pass the config file with `--config` flag. Default config file `dev.env`


### Clone The Repository:

``` bash
# Download repository:
git clone https://github.com/IstiyakRiyad/restaurant-service.git

# Go to parent directory:
cd restaurant-service
```

### Menual Build Process:
Please put local db config to the .env & the migrations location on .env file(`MIGRATION_FILE_URL="file:///$PROJECT_DIR/migrations"`). 
``` bash
# Install dependencies:
go mod download

# Build the Project
go build -o bin/restaurant

# Setup config file
cp dev.env .env

# Run migrations
./bin/restaurant migrate up --config .env

# ETL 
./bin/restaurant etl --config .env

# provide the config file to --config flag
./bin/restaurant start --config .env
```

### Docker Process:
<b>Build Image Manually:</b> <br />
First clone the repository then enter into the repository. <br />
To build the just the docker file run:
``` bash
docker build -t istiyak/restaurant .
```

<b>Run Docker Image Manually:</b><br />
Here the docker image take `/app/prod.env` as config file by default. I used `--network host` because my database on the local machine. Please put local db config to the .env & the migrations location on .env file (`MIGRATION_FILE_URL="file:///migrations"`).
``` bash
docker run --rm -d -v ./.env:/app/prod.env --network host istiyak/restaurant
```

<b>Run with Docker Compose:</b><br />
``` bash
# copy docker.env to .env & modify
cp docker.env .env

# Build
docker compose build

# Run docker compose
docker compose up -d
```

<b>ETL(Extract, Transform and Load):</b> <br />
If you don't load data to the database before then run following command.
``` bash
docker exec restaurant restaurant etl --config /app/prod.env
```
<b>Testing:</b> <br />
I have added some integration testing. 
``` bash
go test -cover ./...
```

## API Documentations:
* [README API DOCS](https://github.com/IstiyakRiyad/restaurant-service/-/blob/main/docs/README.md)
* [Insomnia Docs](https://github.com/IstiyakRiyad/restaurant-service/-/blob/main/docs/Insomnia_docs.json)

## Technology:

### Golang Libraries:
* `gin`
* `cobra`
* `viper`
* `golang-migrate`
* `pq`
* `database/sql`
* `context`
* `encoding/json`
* `stretchr/testify`

## Tools
* `nvim`
* `air`
* `Makefile`
* `docker`
* `docker compose`

## Project Structure:

* `etl` - code for etl(encoding/jsonextract, transform and load) the given data.
* `db`  - database connect & database query files
* `cmd` - Different command line code.
* `migrations` - Migration sql files
* `transport` - HTTP connection functionality code
* `internal` - Core logic of the application
* `utils` - Different helper functions
* `bin` - Binary executable files
* `restaurantData` - Given json files(restaurant & user data)
___
