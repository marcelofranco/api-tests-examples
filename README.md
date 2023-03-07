# API Testing examples
Project to write api tests using multiple tools and techniches.

## postman
Testing the API [Restful-Booker](https://restful-booker.herokuapp.com/) using Postman.

Validating responses and trying to break inputs.

## contract-test
Simple consumer and provider project in [Go](https://go.dev/) using [gin](https://github.com/gin-gonic/gin) and [gorm](https://gorm.io/gorm).

Test with [pact-go](https://github.com/pact-foundation/pact-go)

### Running with docker-compose
- Need to have installed [docker](https://docs.docker.com/engine/install/ubuntu/) and [docker-compose](https://docs.docker.com/compose/install/linux/)


#### With make
``` bash
cd contract-test
make up_build


# Consumer tests
cd consumer/cmd/api
go test

# Provider tests - broker is not configured, so its necessary to copy pact files to test provider
cd provider/cmd/api
go test
```

