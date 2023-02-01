
# Golang Account Microservice

A simple microservice for user registration and sign up using Golang.





## Authors

- [Juan Esteban Oviedo](https://www.github.com/juanoviedog)


## Tech Stack

**Server:** Golang, Gin for REST API, Gin Middleware for Authentication Requirement.

**Database Service:** SQLite using Gorm and Gorm Driver

**JWT Token**: Go-JWT

**Data Encryption**: Bcrypt


## Features

- User registration.
- User login.
- JWT Token issuance and validation for users.
- Authorization middleware for future routes. 


## Screenshots

### User Sign Up
![App Screenshot](https://raw.githubusercontent.com/juanoviedog/golang-ms/main/screenshots/SignUp.png)

### User Log In: 
Upon login user is assigned a JWT Token with a given expiration (30 days). The token is stored inside a cookie in the browser under the 'Authorization' name. Both the duration and the name of the cookie can be changed at will. 

![App Screenshot](https://raw.githubusercontent.com/juanoviedog/golang-ms/main/screenshots/LoginCookie.png)

## Token Validation:
The service allows for the validation of the token using Gin middleware, if the token is valid the middleware returns the user data, otherwise it will return an error as seen on the second image.

![App Screenshot](https://raw.githubusercontent.com/juanoviedog/golang-ms/main/screenshots/TokenValidation.png)

If the token is invalid or there is no cookie set with the given name the user receives a 401 Unauthorized code. 
![App Screenshot](https://raw.githubusercontent.com/juanoviedog/golang-ms/main/screenshots/TokenValidationError.png)

## Deployment

To run this project you need to have installed Go 1.19 as well as the required packages.

```bash
  go run main.go
```
