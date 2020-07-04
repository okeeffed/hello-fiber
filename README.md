# Hello Fiber

This is a repo to accompany my [blog post](https://blog.dennisokeeffe.com/blog/2020-07-024-hello-fiber/).

## Running

```s
go run main.go
```

## cURL examples

```s
# JSON
> curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"Dennis\",\"age\":27}" localhost:3000/person
Success%
# XML
> curl -X POST -H "Content-Type: application/xml" --data "<login><name>Dennis</name><age>27</age></login>" localhost:3000/person
Success%
# x-www-form-urlencoded
> curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "name=Dennis&age=27" localhost:3000/person
Success%
# Form data
> curl -X POST -F name=Dennis -F age=27 http://localhost:3000/person
Success%
# Query values
> curl -X POST "http://localhost:3000/person?name=Dennis&age=27"
Success%
# Error Example passing age as a string
> curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"Dennis\",\"age\":\"27\"}" localhost:3000/person
Failed%
```
