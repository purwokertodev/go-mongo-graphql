### Golang, GraphQL & Mongo Demo

## Recipes
   - Golang version 1.7+ https://golang.org/

   - Glide https://github.com/Masterminds/glide

   - MongoDb https://www.mongodb.com/

## Install Dependencies
  ```shell
  $ glide install
  ```

## Running
  ```shell
  $ MONGO_HOST=localhost MONGO_DB_NAME=tutorial1 go run main.go
  ```

## graphql
  - Create Profile
  ```shell
  $ curl -g 'http://localhost:8080/graphql?query=mutation+_{createProfile(id:"U5",firstName:"Andre",lastName:"De",email:"andre@yahoo.com",password:"123456"){id,firstName,lastName,email,password,createdAt,updatedAt}}'
  ```

  - Get Profile
  ```shell
  $ curl -g 'http://localhost:8080/graphql?query={profile(id:"U3"){id,firstName,email}}'
  ```

  - Get Profiles
  ```shell
  $ curl -g 'http://localhost:8080/graphql?query={profiles{id,firstName,email,password,updatedAt}}'
  ```

###

2018 Purwokerto Dev
