# Hotel Reservation backend

## Project outline
- users -> book room in hotels
- admins -> going to check for reservations
- Authentication and authorization -> JWT Token
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> db management -> seeding, migration

## Resources
### MongoDB driver
Documentation
```
https://www.mongodb.com/docs/drivers/go/current/
```

Installing MongoDB client
```
go get https://go.mongodb.org/mongo-driver/mongo/
```

### gofiber
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```