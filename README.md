# Events REST API

A simple example REST API written in Go with a PostgresSQL
database and sqlx, containerised with Docker

## Running

## Supported Routes

`GET /events`

> Gets all events from the database

`GET /events/:id`

> Gets the event in the database with the given id if it
> exists

`POST /events`

> Creates an event in the database


`PUT /events`

> Updates an event in the database

`DELETE /events/:id`

> Deletes the event with the given id in the database

