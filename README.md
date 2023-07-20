### Pokemon-API Docs


- `GET /api/v1/pokemon/catch-probability	`
  REST API to return probability is 50% when catching Pokemon

  > Catch pokemon

  50% or 0.5, 

  success probability is 50% 

  if the probability is less than or equal to 0.5 it will success catch pokemon, by example Probability 0.2970735
  
  if the probability is more than 0.5 it will fail to catch pokemon, by example Probability 0.5526671

  

  _Path Example_
  
  ```
  https://pokemaster-af3f8e6d4ac3.herokuapp.com/api/v1/pokemon/catch-probability
  ```
  
  _Response (208)_
  
  ```json
  {
    "status": true,
    "message": "OK",
    "code": 208,
    "data": {
      "success": true,
      "probability": 0.4434599,
      "information": "success to catch pokemon"
    }
  }
  ```
  
  _Response (500 - Internal Server Error)_
  
  ```json
  {
    "status": false,
    "message": "internal server error",
    "code": 500
  }
  ```
  
  ---
  
-  `GET /api/v1/list-pokemon	`

  List of pokemon that has been caught and added to favorites

  > Get All Pokemon 

  

  _Path Example_

  ```
  https://pokemaster-af3f8e6d4ac3.herokuapp.com/api/v1/pokemon/list-pokemon
  ```

  _Response (200)_

  ```json
  {
    "status": true,
    "message": "OK",
    "code": 200,
    "data": [
      {
        "id": "599463ff-a7cf-4bbc-b8cf-e4d013ae000b",
        "pokemon_name": "Unwon",
        "pokemon_picture": "https://img.pokemondb.net/artwork/large/unown.jpg",
        "number": 2,
        "user_id": "1",
        "created_at": "2023-07-19 17:09:06.370802191 +0700 WIB m=+1649.963860158",
        "updated_at": ""
      },
      {
        "id": "66cbdeb0-0df4-4e01-bf0d-07b95c6a7ca6",
        "pokemon_name": "Wobbu-0",
        "pokemon_picture": "https://img.pokemondb.net/artwork/large/wobbuffet.jpg",
        "number": 3,
        "user_id": "1",
        "created_at": "2023-07-19 17:09:28.145093875 +0700 WIB m=+1671.738151353",
        "updated_at": "2023-07-19 17:10:09.087322029 +0700 WIB m=+1712.680379926"
      },
      {
        "id": "a97b0038-522d-4aeb-ace6-99f7f70efdf1",
        "pokemon_name": "pineco",
        "pokemon_picture": "https://img.pokemondb.net/artwork/large/pineco.jpg",
        "number": 19,
        "user_id": "1",
        "created_at": "2023-07-19 17:11:31.456813059 +0700 WIB m=+1795.049871026",
        "updated_at": ""
      },
      {
        "id": "bf590811-79a6-48b9-9eae-b7e660d9ac75",
        "pokemon_name": "azez",
        "pokemon_picture": "https://img.pokemondb.net/artwork/large/girafarig.jpg",
        "number": 7,
        "user_id": "1",
        "created_at": "2023-07-19 17:10:40.610691008 +0700 WIB m=+1744.203748974",
        "updated_at": ""
      },
      {
        "id": "d556e01d-422c-4022-a4ce-72e50630590f",
        "pokemon_name": "azez",
        "pokemon_picture": "https://img.pokemondb.net/artwork/large/girafarig.jpg",
        "number": 11,
        "user_id": "1",
        "created_at": "2023-07-19 17:10:56.591475311 +0700 WIB m=+1760.184533697",
        "updated_at": ""
      }
    ]
  }
  ```
  
  _Response (500 - Internal Server Error)_
  
  ```json
  {
    "message": "internal server error",
  }
  ```
  
  Also this rest api can search by pokemon name, example

  ```jso
  https://pokemaster-af3f8e6d4ac3.herokuapp.com/api/v1/pokemon/list-pokemon?pokemon_name=wobuffet
  ```
  
  
  | Parameter    | Data Type |     Location |
  | :----------- | :-------: | -----------: |
  | pokemon_name |  String   | query header |
  
  
  ---
  
- `POST /api/v1/pokemon/release-pokemon`
  REST API to release pokemon. This API should return a prime number, if the number returned by the API is not a prime number, then release will fail and vice versa.

  > Release Pokemon

  

  note prime number is 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, etc

  prime number input prime number except the numbers above, wll be error (not insert)
  
  
  
  _Request Body_
  
  ```json
  {
  	"pokemon_name": "azez",
  	"pokemon_picture": "link",
  	"number": 3, // for prime number
  	"user_id": "1" // in my case user id sent from core not from user
  }
  ```
  
  _Response (201)_
  
  ```json
  {
    "status": true,
    "message": "OK",
    "code": 200,
    "data": {
      "result": "success",
      "released": true,
      "prime_number": 3
    }
  }
  ```
  
  _Response (500 - Internal Server Error)_
  
  ```json
  {
    "message": "internal server error"
  }
  ```
  
  Bad Request (400 - bad request) - if user input not prime number
  
  ```json
  {
    "message": "number must be a prime number"
  }
  ```
  
  ---
  
  
  
- `PUT /api/v1/pokemon/rename-pokemon/:id`
  REST API to rename pokemon. This function should return a combination of first name assigned combined
  with Fibonacci sequence.
  
  > Rename pokemon
  
  
  
  Change name return a combination of first name assigned combined with Fibonacci sequence
  
  Fibonancci number:
  
   **0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181**.
  
  
  
  by example:
  
  • First name assigned is “Mighty Pikachu”, first time renamed should be: “Mighty Pikachu-0”
  • Second time renamed should be: “Mighty Pikachu-1”
  • Third time renamed should be: “Mighty Pikachu-1”
  • Fourth time renamed should be: “Mighty Pikachu-2”, and so on.
  
  
  
  _Path Example_
  
  ```
  https://pokemaster-af3f8e6d4ac3.herokuapp.com/api/v1/pokemon/rename-pokemon/d556e01d-422c-4022-a4ce-72e50630590f
  ```
  
  _Request Body_
  
  ```json
  {
  	"pokemon_name": "Mighty Pikachu"
  }
  ```
  
  _Response (200)_
  
  ```json
  {
    "status": true,
    "message": "OK",
    "code": 200,
    "data": {
      "id": "f7e11c11-2b9a-4c45-9a16-853a1b1bd367",
      "pokemon_name": "Mighty Pikachu-0",
      "created_at": "2023-07-19 14:08:07.196095389 +0700 WIB m=+5331.265260785",
      "updated_at": "2023-07-19 14:08:07.196117668 +0700 WIB m=+5331.265283064"
    }
  }
  ```
  
  _Response (500 - Internal Server Error)_
  
  ```json
  {
    "message": "internal server error"
  }
  ```
  
   _Bad Request (400 - bad request)_
  
  ```json
  {
    "status": false,
    "message": "Bad Request",
    "code": 400,
    "error_validation": {
      "pokemon_name": "the name of pokemon is required"
    }
  }
  ```
  
  ---
  
  



  

