<div align="center">
  <h1>Pathao Technical Assessment</h1>
  <p>
    This is a simple text animation program. When you hover on the text it will be animated.
  </p>
  <p>Submitted By: <a href="https://github.com/IstiyakRiyad" target="_blank">Md. Istiyak Hossain</a> </p>
</div>

## Prerequisite
For manual build install golang & postgresql database locally. Setup the configuration file. Pass the config file with --config flag. Default config file `dev.env`
## Menual Build Process:

``` bash
# Download repository:
git clone https://gitlab.com/IstiyakRiyad/technical-assessment-pathao.git

# Go to parent directory:
cd technical-assessment-pathao

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
## Project Structure:

* `etl` - code for etl(extract, transform and load) the given data.

