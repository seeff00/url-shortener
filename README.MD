# URL Shortener
This repository contains two projects. Both works together. They generate short URLs.

1. Web application:

Frontend - The project represents a Web application build on React v18.2.0. The project
can be found in the directory [url-shortener-api](https://github.com/seeff00/url-shortener/tree/main/url-shortener-api).

2. HTTP API service:

Backend - The HTTP API service is build on GoLang v1.21.5. It has own Docker setup
which provides all required prerequisites. Such as MySQL database and GoLang v1.21.5. 
On running the Docker setup will automatically import prepared database dump file 
named 'dump.sql' and located here "Link". The database dump creates a database named 
'url_shortener' and all required tables and columns in it.

The project can be found in the directory [url-shortener-web](https://github.com/seeff00/url-shortener/tree/main/url-shortener-web)

### Clone repository

1. Open Linux terminal and navigate to your projects directory.
```shell
# Setup
cd /projects/
git clone https://github.com/seeff00/url-shortener.git
```