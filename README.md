# sqlc-demo
```
# Checkout old version with just one table
$ git checkout 0.1

# Generate the .go files via sqlc
$ cd sqlc && sqlc generate

# Build and start containers
$ docker-compose up --build

# Create post
$ http POST http://localhost:8080/add-post < add-post.json

# Get post
$ http http://localhost:8080/get-post/1
```

```
# Checkout latest version with two tables
$ git checkout main

# Re-generate the sqlc built .go files
$ cd sqlc && sqlc generate

# Build and start containers
$ docker-compose up --build

# Create post
$ http POST http://localhost:8080/add-post < add-post.json

# Get post
$ http http://localhost:8080/get-post/1
```