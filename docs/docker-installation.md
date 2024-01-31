# Building & Running (the vanilla Docker way)
The docker-compose file creates 3 containers: api, web and db.

1. Dependencies
   Make sure you have installed:
   * Docker (w/ Compose)

2. Copy required environment file
    ```bash
    cp .env.example .env
    # make necessary changes to .env file
    ```

3. Build & run the managed containers
    ```bash
    # build containers
    docker compose build
    # start managed containers
    docker compose up -d
    ```

4. Access the app
   To access the web interface of Kraken, navigate to [http://localhost](http://localhost) in your browser.  
   To access the REST API of Kraken, you can access it at [http://localhost:8888](http://localhost:8888). For example, the following command checks the status of the API server:
   ```bash
   curl localhost:8888/ping 
   ```
   If everything works out, you should see `{"message":"Kraken API is running on host: HOSTNAME","status":"ok"}` (assuming your hostname is `HOSTNAME`).