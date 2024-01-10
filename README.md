# KrakenCluster
Monorepo for software containerization assignment

## Setup instructions
Make sure you have installed:
- Docker
- Go

The docker-compose file creates 3 containers: api, web and db.

First, build the web application by going to: `projectroot/web`.
Then run: `go build`.

Next, go to the root of your project and copy the `.env.example` file.
Call the copy: `.env`, you can change the values inside that file if you like.

Then you should run:
```bash
# Linux
docker-compose build
docker-compose up -d

# Mac
docker compose build
docker compose up -d
```

If you want to view the website you have to visit the ip address of the web docker container.
To do that run: `docker inspect web`.

This will print the properties of the docker container, in there you will find the ipaddress.

Visit this addres in your browser and you should see the website.

You can do the same for the api. Inspect the api container and then make requests to that address.

### Run the website and api without docker
To run the api and website without using docker you must build both projects and run the binary, or call `go run .`

So running the api is done going to the api directory. Then run `go run .`, or run `go build -o filename` and then `./filename`.

Running the website is exactly the same, but make sure you are in the web directory when running the commands.

## vscode
Use the following extension for better highlighting in the html templates: Go Template Support by jinliming2
