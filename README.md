# Currency conversion API

Small application to convert currencies based on `fixer.io` data.

Tech stack:
 * go 1.16
 * gorilla mux
 * docker
 * docker-compose
 * make

## Project structure

* `api/fixer` - code to query Fixer API
* `convert` - currency conversion code
* `env` - code to work with environment variables
* `http` - HTTP API
* `project` - logical structure for storing projects
* `sync` - logic to query Fixer API and process results
* `docs` - example API calls and example response

## Development steps

* `make` - build and run test for application
* `make clean` - clean working directory
* `make test` - run tests and generate coverage results
* `make container` - build docker container with application

Example API calls can be found in `docs` directory, basic flow:
 * Create project and obtain API key
 * Query service with API key and exchange payload

## Important environment variables

* `FIXER_API_KEY` - API key to access Fixer data, app will fail if missing
* `SYNC_INTERVAL_MIN` - background sync interval of currencies, default 60 min
* `TRACE` - if defined additional logs will be printed in output
* `HTTP_PORT` - listen HTTP port

Starting through docker-compose application:
```
FIXER_API_KEY=<api-key> docker-compose up -d
```

## Future work

* CI/CD integration
* Integration tests
* Persistence for projects
* Admin API endpoints
* Example client to API

# License

Only for reference
