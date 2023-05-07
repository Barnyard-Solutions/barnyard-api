`docker build -t barnyard-api . && docker create  -e DB_HOST=barnyard-database --network alpha  -p 5000:5000 --name barnyard-api1 barnyard-api && docker network connect beta barnyard-api1 && docker start -a barnyard-api1`




GET/POST /api/user: creates a new account or logs in with an existing one
POST /api/user/token: generates an access token for a given account
GET/POST /api/feed: gets the list of available feeds or creates a new one
DELETE /api/feed/:id: removes a feed by its ID
GET/POST /api/feed/:id/event: gets the list of available events in a given feed or creates a new one
DELETE /api/feed/:id/event/:id: removes an event by its ID from a feed
GET/POST /api/feed/:id/milestone: gets the list of available classes in a given feed or creates a new one
DELETE /api/feed/:id/milestone/:id: removes a class by its ID from a feed