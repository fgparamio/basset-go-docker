# Basset Go Microservice Docker Application


## Overview

Application built with:

```
Go Verion "1.10"
Visual Studio Code "1.20.1"

```

Go Libraries in this project:

```
github.com/gorilla/handlers
github.com/gorilla/mux
github.com/gorilla/context
github.com/stretchr/testify/suite <= Only To Testing 

```

## Demo

Open: [http://bassetbackgo-env.us-east-2.elasticbeanstalk.com/users](http://bassetbackgo-env.us-east-2.elasticbeanstalk.com/users) 

to view API App in the browser.


## Available Scripts

### To Run Application

```
make run 

```

### Makefile options

    all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	install \
	test \
	clean 


In the project directory, you can run:

### `go build serve.go; ./serve`

Run App the app .<br>
Open [http://localhost:4000/users](http://localhost:4000/users) to view it in the browser.<br><br>


### `go test`

Launches the test runner.<br><br>


### `go test -test.bench .`

To run the benchmarks


## API Rest Examples

- Get List Users

```
GET /users   <== Return first Page
Accept: application/json
Content-Type: application/json


RESPONSE: HTTP 200 (OK)


{
page: 1,
perPage: 3,
total: 12,
totalPages: 4,
users: [
	{
		ID: 1,
		name: "George Bluth",
		avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/calebogden/128.jpg"
	},
	{
		ID: 2,
		name: "Janet Weaver",
		avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/josephstein/128.jpg"
	},
	{
		ID: 3,
		name: "Emma Wong",
		avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/olegpogodaev/128.jpg"
	}
 ]}
)
```

```
GET /users?page=2
Accept: application/json
Content-Type: application/json

RESPONSE: HTTP 200 (OK)

{
page: 2,
perPage: 3,
total: 12,
totalPages: 4,
users: [
	{
		ID: 4,
		name: "Eve Holt",
		avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/marcoramires/128.jpg"
	},
	{
		ID: 5,
		name: "Charles Morris",
		avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/stephenmoon/128.jpg"
	},
	{
		ID: 6,
		name: "Tracey Ramos",
		avatar: "https://s3.amazonaws.com/uifaces/faces/twitter/bigmancho/128.jpg"
	}
  ]}
}

```


## Docker

You can build your image:

```
docker build -t server/gobasset .

```

 Run Docker image:
 
```
docker run --rm -p 5000:5000 server/gobasset


```
Surf to http://localhost:5000/ to see it running. <br/><br/>




You can see that if you re-run the build command from earlier:
 
```
time docker build -t goBack-docker .

```

## Folder Structure

```
basset-go-docker/
  README.md         <= This file
  serve.go          <= API Rest implementation
  serve_test.go     <= Test API Application
  Dockerfile        <= Docker container configuration
 
```

