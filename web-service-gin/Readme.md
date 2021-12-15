## An API that provides access to a store selling vintage recordings on vinyl. Exposed endpoints through which a client can get and add albums for users.

### Here are the endpoints we’ll create in this tutorial.

- /albums
 - GET – Get a list of all albums, returned as JSON.
   example:- http://localhost:9999/albums
 - POST – Add a new album from request data sent as JSON.
   example:- post some JSON payload {"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99} to 
             http://localhost:9999/albums and then use GET Api to validate.

- /albums/:id
 - GET – Get an album by its ID, returning the album data as JSON.
   example http://localhost:9999/albums/3

### Create the data
We’ll store data in memory. A more typical API would interact with a database.
> **Note** that storing data in memory means that the set of albums will be lost each time we stop the server, then recreated when we start it.

### Instructions to run
``
- cd web-service-gin
- go run .
- this will run a HTTP server to port 9999, which can accept requests.
``

**Author**
>Gaurav