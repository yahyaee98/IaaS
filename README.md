# IaaS 
[![Build Status](https://travis-ci.org/yahyaee98/IaaS.svg?branch=master)](https://travis-ci.org/yahyaee98/IaaS)
[![Go Report Card](https://goreportcard.com/badge/github.com/yahyaee98/IaaS)](https://goreportcard.com/report/github.com/yahyaee98/IaaS)
[![Coverage Status](https://coveralls.io/repos/github/yahyaee98/IaaS/badge.svg?branch=master)](https://coveralls.io/github/yahyaee98/IaaS?branch=master)

"Items as a Service" is a sample project written in golang which finds books and musics from Google Books and iTunes API and presents them in a simple API.

Books and musics are referred to as "Item"s. 

## Usage

### Run
Just run `docker-compse up --build` to build the application in a container and have it running on port `8080` by default. A Redis instance is also included.

### Fetch results
A sample `GET` request can be made to `http://127.0.0.1:8080/api/results?search=beautiful`.

This is a result:

```json
{
  "items": [
    {
      "title": "Beautiful (feat. Camila Cabello)",
      "type": "music",
      "creators": [
        "Bazzi"
      ]
    },
    {
      "title": "Beautiful Red",
      "type": "book",
      "creators": [
        "M. Darusha Wehm"
      ]
    }
  ]
}
```

You can find the detailed API documentation [here](https://github.com/yahyaee98/IaaS/blob/master/api/swagger.yml).

### Unit tests
There are some unit tests included which you can run them by running `make test`.

## Project structure

The project structure is inspired by [project-layout](https://github.com/golang-standards/project-layout). 

`/api`: contains API documentation.

`/bin`: will contain the compiled application binary.
 
`/internal`: contains the main part of the code. These are the private codes that can not be easily used in other projects.

`/pkg`: includes some packages that can easily be used in other projects as they are not coupled to the internal codes.

## Packages

### API
Holds the logic for the HTTP server including the handlers for different endpoints.

### Cache
Introduces a `Cache` interface and also a `RedisCache` implementation.

### Data
It is responsible for holding various structs that are used in the application including responses.

### Log
I made this to make it possible to change the logger library later.
It also includes a global `log` variable which eases the logic process, although it's not the best way and the logger should also be injected.

### Metric
Includes histograms which we use to report third party upstream response times. 

### Repository
It contains all the logic for retrieving "Item"s. It takes care of either fetching fresh results using an `Upstream` or returning the cached result.

### Upstream
Holds various adapters to Google Books and iTunes client libraries.

### GoogleBooks
It contains a client for retrieving books from Google Books API.

### Itunes
It contains a client for retrieving musics from iTunes API.
