# Echo Proxy + Jaeger Example
This project aims to demonstrate the use of Echo with the proxy and jaeger tracing middleware. 

## Getting Started
You can launch the project by running the following command:
```
docker compose up -d --build
```
And test by calling the following endpoint:
```
curl localhost:8080/hello
```
You can then view the tracing by going to `localhost:16686` in your browser.

## Tracing Issue
There is currently an [issue](https://github.com/labstack/echo-contrib/pull/110) with the jaegertracing middleware in that when it is combined with the builtin proxy middleware. It creates a new span within the current `echo.Context` but it is not propagated in the headers which the proxy middleware forwards to the next service. Therefore rather than one continuous span from the proxy to the service and back. It ends up with two spans, one from the proxy and one from the service.