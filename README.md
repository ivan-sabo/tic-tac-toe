# tic-tac-toe

This is an API backend project for tic-tac-toe game :)

Requirements:
    golang 1.18
    docker
    docker-compose

To run the project you will have to first run "docker-compose up -d" inside root folder.
This will run docker container with required database. As you could see in docker-compose.yaml,
volume is not used for persistant storage.

Next, you will have to run "go run ./cmd/admin/migrate.go".
This will run necessary migration and create a database table inside your docker container.

Finally, run "go run ./cmd/api/main.go".
This will fire up game and expose api on localhost:8080/v1/games.

Supported endpoints:
GET /v1/games
POST /v1/games          - expects either an empty body, or first move. This is an example of first move request body: {"board":"X--------"}
GET /v1/games/{uuid}
PUT /v1/games/{uuid}

Todo:
    - write unit tests
    - make database and api configurable (ports, username, password, etc)
    - add strategy pattern for AI player algorithms
    - add dependency injection
