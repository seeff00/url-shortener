# URL Shortener API
It's an API service, and it's built on Go v1.21.4.

## Prerequisites & Setup

### Prerequisites

1. Docker - Can be downloaded/installed from the official [get docker page](https://docs.docker.com/get-docker/)

### Config
The project have a config file in the root directory of the project '.env.dev'. Rename config file '.env.dev' -> '.env'.
```shell
mv .env.dev .env
```

### Setup & Run

1. Open Linux terminal and navigate to your project's directory.
```shell
# Setup
cd /projects/
git clone https://github.com/seeff00/url-shortener.git
cd url-shortener/url-shortener-api/

# Run
docker-compose up -d
docker exec -it url-shortener-api bash
go mod tidy
go mod vendor
go run main.go
```

## Endpoints
Edpoint | Method | Params | Description
--- |--------|-------- | -----------
/short | POST   | url | Generates short url and return it in response.   
/:url_code | GET    | :url_code | Redirect from short url to original.

Example Request to "/short" endpoint:

```shell
curl --location 'http://localhost:8080/short?url=https%3A%2F%2Fioflood.com%2Fblog%2Fdocker-stop-all-containers-one-command-to-stop-and-or-remove-every-docker-container%2F' \
--header 'Content-Type: text/plain' \
--data '{
    "url":"https://translate.google.bg/?hl=bg&sl=auto&tl=bg&text=Hello%20world&op=translate"
}'
```

Example Response of "/short" endpoint:
```json
{
    "code": "H4mHc-orDQ",
    "url": "https://translate.google.bg/?hl=bg&sl=auto&tl=bg&text=Hello%20world&op=translate",
    "ip": "172.18.0.1",
    "created_at": "2023-12-18T06:26:14.073575741Z"
}
```

## Keep in mind

In the root of the project there is a 'dump.sql' file. On executing 'docker-compose up -d' it will up 
MySql docker image and container and will import everything from this dump file to DB.