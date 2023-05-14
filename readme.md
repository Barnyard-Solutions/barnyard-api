<div align="center">



<img width=200 src="https://user-images.githubusercontent.com/37497007/236691167-e476560b-f87a-43e3-aa9a-197fc204d175.svg">

# Barnyard-API

Barnyard-API is a component of the Barnyard project, which is a progressive web app designed to help manage barnyard animals, including pigeons, chickens, rabbits, pigs, and more. The API component is responsible for retrieving and updating barnyard animal information.

</div>


## 🛠️ Technologies Used

Docker <br>
Golang <br>
Python <br>

## 🌐 Usage
To use the Barnyard-API component, follow these steps:

Build the Docker image using the following command:
```
docker build -t barnyard-api .
```
Create a Docker container using the following command:
```
docker create -e DB_HOST=barnyard-database --network alpha -p 5000:5000 --name barnyard-api1 barnyard-api
```
Connect the Docker container to the beta network using the following command:
```
docker network connect beta barnyard-api1
```
Start the Docker container using the following command:
```
docker start -a barnyard-api1
```
## 🚀 API Endpoints

GET/POST /api/user : Creates a new account or logs in with an existing one.

POST /api/user/token : Generates an access token for a given account.

GET/POST /api/feed : Gets the list of available feeds or creates a new one.

DELETE /api/feed/:id : Removes a feed by its ID.

GET/POST /api/feed/:id/event : Gets the list of available events in a given feed or creates a new one.

DELETE /api/feed/:id/event/:id : Removes an event by its ID from a feed.

GET/POST /api/feed/:id/milestone : Gets the list of available classes in a given feed or creates a new one.

DELETE /api/feed/:id/milestone/:id : Removes a class by its ID from a feed.

GET /api/feed/:id/subscribe : Retrieves the subscription status of a user by their ID.

POST /api/feed/:id/subscribe : Subscribes a user to push notifications by their ID.

POST /api/user/notification-time : Sets the notification time for a user by their ID.

GET /api/feed/:id/member : Retrieves the list of members in a feed by its ID.

POST /api/feed/:id/member : Adds a member to a feed by its ID.

PUT /api/feed/:id/member : Updates the permission of a member in a feed by their IDs.

DELETE /api/feed/:id/member : Removes a member from a feed by their IDs.

## 🙋 Contributors

Contributions to the Barnyard project are always welcome! If you would like to contribute, please feel free to submit a pull request or open an issue on the Barnyard-API repository.

Thank you to all of our contributors who have helped make the Barnyard project a reality!
