# go-api

## Description
This is a simple API built using Go and MySQL.

## Tools
The following tools/packages/modules were used:

- Database: `MySQL`
- Web server: `Mux`
- ORM: `Gorm`

## API Design
The API was build using the following folder structure:

```
├── config          // Database settings
├── controllers     // Business logic
├── models          // Database tables
├── repository      // Functions used to query data from the database
├── routes          // Routes of the API
├── schemas         // Responses
└── utils           // Common functions to be used
```

We can make it easier to understand with the picture below.
![API Flow](/docs/flow.png)

1- User makes a HTTP request to the API.
2- The specified route will execute a specific controller function.
3- The function needs to interact with the database to gather information. To better isolation of each layer, the controller function will call the repository function to interact with the database and gather the data.
4- The repository is responsible for direct communication with the database.
5- Repository will return the gathered information to the controller.
6- Controller will execute the business logic and use the schema layer to create the response.
7- HTTP response is returned to the user.


## Running the API
A docker-compose was created to setup, using Docker, the API and MySQL dabase. 

So simply running `docker-compose up` will spin up all the infrastructure.

## Routes

The following routes were created.

- GET `/user`

Get all users.

- GET `/user/{id}`

Return information about a specific user given its ID.

- POST `/user`

Adds a new user to the database. Accepts the following body:
```json
{
    "name": << name of the user >>,
    "email": << email of the user>>,
    "status": << bool indicating if user is active or not>>
}
```

- PUT `/user/{id}`

Edit information about a user given its ID. Accepts same body as the POST `/user` route.

- DELETE `/user/{id}`

Delete a user given its ID.

