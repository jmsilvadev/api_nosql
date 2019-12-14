# NoSql API
API RestFULL to interacte with Big Data & NoSQL DB. Code created using Go Programming Language version 1.12.3 (https://golang.org/).

## Current Version
The actual version is 0.1.0

## Overview

In this current version, the default NoSql is ElasticSearch.

## API endpoints

#### GET
- `/v1/`<br />
- `/v1/info`<br />
- `/v1/license`<br />
- `v1/{indexname}`<br />

#### POST
- `/v1/{indexname}`<br/>

#### DELETE
- `/v1/{indexname}`<br/>

___


## API Endpoint Details

## GET /v1/


Call the HealthInfo method - Shows if this API is UP and Running

##### Response

```ruby
You know, to interact with the NoSQL api. :)
```


## GET /v1/info


Call the ShowInfo method - Shows information about this API

##### Response

```ruby
[
    {
        "title": "NoSQL Manager Service.",
        "description": "API RestFULL to interacte with Big Data & NoSQL DB.",
        "version": "0.1.1",
        "request_ip": "<ipaddress>:<port>"
    }
]
```

## GET /v1/license


Call the ShowLicense method - Shows information about the License of this API


##### Response

```ruby
[
    {
        "name": "JMSilvaDev © 2019 | Todos os direitos reservados",
        "request_ip": "<ipaddress>:<port>"
    }
]
```


## GET /v1/{indexname}


Shows Index Data

##### Response

```ruby

```



## DELETE /v1/{indexname}


Delete an Index

##### Response

```ruby
{
    "Name": "indexname",
    "response": "Index Deleted Successfully"
}
```

or

```ruby
{
    "Name": "indexname",
    "response": "Fail! Index Not Deleted."
}
```


## POST /v1/{indexname}


Create an Index

##### Response


```ruby
{
    "Name": "indexname",
    "response": "Index Created Successfully"
}
```

or

```ruby
{
    "Name": "indexname",
    "response": "Fail! Index Not Created."
}
```

## Dev Work

1 - Setting environment variables: 

``` bash
export GOPATH=$HOME/go; export GOROOT=/usr/local/go; export PATH=$GOPATH/bin:$GOROOT/bin:$PATH;
```

2 - Clone the repo:

``` bash
git@github.com/jmsilvadev/api-nosql.git
```

3 - Check if there are any other dependencies to install:

```bash
go build -o build/apinosql (inside folder /go/src/api-nosql/)
```

Install required dependences!


## Testing Using REST Client

``` bash
go run .
```