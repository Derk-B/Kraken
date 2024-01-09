# KrakenCluster
Monorepo for software containerization assignment

## How to run
In the root directory run: 
```bash
$ go run .
```

Or if you want to compile and run: 
```bash
$ go build
$ ./KrakenSite
```

## Documentation
This project uses the "gin" package for creating a REST api.
There are 2 submodules in this project, `api` and `website`.

All requests starting with "/api" are sent to the routes that are defined in the `api` submodule.

All other routes are send to the `website` submodule.

## vscode
Use the following extension for better highlighting in the html templates: Go Template Support by jinliming2
