# Go Roller Coasters REST API

Simple REST API built in GO.

Data is only stored in memory, not persistent.

Following [kubucation/go-rollercoaster-api](https://github.com/kubucation/go-rollercoaster-api) tutorial

### Data type

Coaster object should look like

```
{
    "name": "name of the coaster",
    "inPark": "the amusement park the ride is in",
    "manufacturer": "name of the manufacturer",
    "height": 30,
}
```

## Running it

```
go run server.go
```

Api will be available on `localhost:8090/coasters`

### Endpoints

- `GET /coasters` returns list of coasters as JSON
- `GET /coasters/{id}` returns details of specific coaster as JSON
- `POST /coasters` accepts a new coaster to be added in `application/json` format
- `GET /coasters/random` redirects (Status 302) to a random coaster
