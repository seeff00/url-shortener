# URL Shortener
This repository contains two projects. They works together and there purpose is to generate short URLs.

1. Web application:

Frontend - The project represents a Web application build on React v18.2.0. The project
can be found in the directory [url-shortener-web](https://github.com/seeff00/url-shortener/tree/main/url-shortener-web).

2. HTTP API service:

Backend - The HTTP API service is build on GoLang v1.21.5. It has own Docker setup
which provides all required prerequisites. Such as MySQL database and GoLang v1.21.5. 
When Docker setup is started it will import a prepared database dump from file 
named 'dump.sql'. The database dump creates a database named 
'url_shortener' and all required tables and columns.

The project can be found in the directory [url-shortener-api](https://github.com/seeff00/url-shortener/tree/main/url-shortener-api).

### Clone repository

1. Open Linux terminal and navigate to your projects directory.
```shell
# Setup
cd /projects/
git clone https://github.com/seeff00/url-shortener.git
```
