# Kraken
Monorepo for Software Containerization project

## Dependencies
Make sure you have installed:
- Docker (Compose)
- Go

## Building & Running
The docker-compose file creates 3 containers: api, web and db.

1. Build the web application
    ```bash
    cd web
    go build
    ```

2. Copy required environment file
    ```bash
    cp .env.example .env
    # make changes to .env file
    ```

3. Build & run the managed containers
    ```bash
    # build containers with latest code changes
    docker compose build
    # start managed containers
    docker compose up -d
    ```

4. Access each deployed component  
   If you want to access the web interface or API, view the website you have to visit the hostname of the docker container.
   Run:
    ```bash
    docker inspect web
    docker inspect api
    ```
This will print the properties of the docker container, in there you will find the hostname. Visit this address in your
browser and you should see the website or API.