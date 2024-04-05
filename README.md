<div align="center">
  <h1>Pathao Technical Assessment</h1>
  <p>
    This is a simple text animation program. When you hover on the text it will be animated.
  </p>
  <p>Submitted By: <a href="https://github.com/IstiyakRiyad" target="_blank">Md. Istiyak Hossain</a> </p>
</div>

## Prerequisite
For manually build the server, install golang & postgresql locally. Setup the configuration file. Pass the config file with `--config` flag. Default config file `dev.env`


### Clone The Repository:

``` bash
# Download repository:
git clone https://gitlab.com/IstiyakRiyad/technical-assessment-pathao.git

# Go to parent directory:
cd technical-assessment-pathao
```

### Menual Build Process:

``` bash
# Install dependencies:
go mod download

# Build the Project
go build -o bin/restaurant

# Run migrations
./bin/restaurant migrate up --config .env

# ETL 
./bin/restaurant etl --config .env

# provide the config file to --config flag
./bin/restaurant start --config .env
```

### Docker Process:
<b>Build Image:</b> <br />
First clone the repository then enter into the repository.
To build the just the docker file run:
``` bash
docker build -t pathao/restaurant .
```

<b>Run Docker Image:</b><br />
Here the docker image take `/app/prod.env` as config file by default. I used `--network host` because my database on the local machine.
``` bash
docker run --rm -d -v ./dev.env:/app/prod.env --network host pathao/restaurant
```

<b>Run with Docker Compose:</b><br />
Here the docker image take `/app/prod.env` as config file by default. I used `--network host` because my database on the local machine.
``` bash
# copy docker.env to .env & modify
cp docker.env .env

# Build
docker compose build

# Run docker compose
docker compose up -d
```

<b>ETL(Extract, Transform and Load):</b> <br />

``` bash
docker exec restaurant restaurant etl --config /app/prod.env
```

## Project Structure:

* `etl` - code for etl(extract, transform and load) the given data.

