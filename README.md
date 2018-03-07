<<<<<<< HEAD
# Basset Go Microservice Docker Application

This project was bootstrapped with [Create React App](https://github.com/facebookincubator/create-react-app).

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

In the project directory, you can run:

### `go build serve.go; ./serve`

Run App the app .<br>
Open [http://localhost:4000/users](http://localhost:4000/users) to view it in the browser.<br><br>




### `go test`

Launches the test runner.<br><br>



### `go test -test.bench .`

To run the benchmarks




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
basset-front-reactJs/
  README.md
  serve.go
  server_test.go
 
```


=======
# basset-go-docker
>>>>>>> ddd6aaac3870d1e77690c7f4afd85171b347f6d9
