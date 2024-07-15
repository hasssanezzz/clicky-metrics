# Multi-Service Shortner - Auth API

This document describes the API for user registration, login, and token validation for a multi-service shortner application.

## Usage
1. install dependencies
    ```
    go mod download
    ```
1. Configure your user data storage mechanism (e.g., database) in the relevant code.
1. Set the required environment variables.
    ```
    POSTGRES_HOST=...
    POSTGRES_PORT=...
    POSTGRES_USER=...
    POSTGRES_PASSWORD=...
    POSTGRES_DB=...

    JWT_SECRET=...
    ```
1. Run the main application to start the API server.
    ```
    go run .
    ```

## Functionalities

This API offers the following functionalities:

* User Registration: Allows users to create a new account with a username, email, and password.
* User Login: Enables users to log in with their username and password combination.
* Token Validation: Validates a provided JWT token to verify user identity.

## Endpoints table

| Method | Path                 | Description                                          |
| ------- | -------------------- | --------------------------------------------------- |
| GET    | /v1/validate         | Validates a JWT token in the Authorization header     |
| POST   | /v1/login            | Logs in a user and returns a JWT token on success     |
| POST   | /v1/register         | Registers a new user and returns the created user data |

## Authentication

This API uses JWT (JSON Web Token) for authentication. A successful login attempt will return a JWT token in the Authorization header of the response. Subsequent requests to protected resources should include this token in the Authorization header.

## Endpoints

### Login 
* Path: `POST /v1/login`
* Request body:
    ```json
    {
        "username": "...",
        "password": "..."
    }
    ```
* Response Example:
    ```json
    {
        "data": {
            "id": 6,
            "username": "...",
            "email": "...",
            "password": "...",
            "created_at": "2024-07-15T17:45:26.771989+03:00",
            "updated_at": "2024-07-15T17:45:26.771989+03:00",
            "urls": null
        },
        "meta": null
    }
    ```
* Response headers:
    ```
    Authorization: <token>
    ```
* Potential errors:
    1. Username does not exist
        ```json
        {
            "errors": {
                "username": [
                    "username does not exist."
                ]
            },
            "meta": null
        }
        ```
    1. Wrong password
        ```json
        {
            "errors": {
                "password": [
                    "incorrect username/password combination."
                ]
            },
            "meta": null
        }
        ```
    1. Validation errors
        ```json
            {
                "errors": {
                    "password": [
                        "Password length should be from 8 to 72 characters."
                    ],
                    "username": [
                        "Usernames can only consist of letters, digits and _, length should be from 4 to 20 characters."
                    ]
                },
                "meta": null
            }
        ```
    1. Missing fields
        ```json
            {
                "errors": {
                    "password": [
                        "password is requried."
                    ],
                    "username": [
                        "username is requried."
                    ]
                },
                "meta": null
            }
        ```

### Register
* Path: `POST /v1/register`
* Request body:
    ```json
    {
        "email": "...",
        "username": "...",
        "password": "..."
    }
    ```
* Response Example:
    ```json
    {
        "data": {
            "id": 6,
            "username": "...",
            "email": "...",
            "password": "...",
            "created_at": "2024-07-15T17:45:26.771989+03:00",
            "updated_at": "2024-07-15T17:45:26.771989+03:00",
            "urls": null
        },
        "meta": null
    }
    ```
* Response headers:
    ```
    Authorization: <token>
    ```
* Potential errors:
    1. Username/email is in use.
        ```json
        {
            "errors": {
                "email": [
                "E-mail is in use."
                ],
                "username": [
                "Username is in use."
                ]
            },
            "meta": null
        }
    1. Validation errors
        ```json
        {
            "errors": {
                "email": [
                "Invalid E-mail format."
                ],
                "password": [
                    "Password length should be from 8 to 72 characters."
                ],
                "username": [
                    "Usernames can only consist of letters, digits and _, length should be from 4 to 20 characters."
                ]
            },
            "meta": null
        }
        ```
    1. Missing fields
        ```json
        {
            "errors": {
                "email": [
                    "email is required."
                ],
                "password": [
                    "password is requried."
                ],
                "username": [
                    "username is requried."
                ]
            },
            "meta": null
        }
        ```

### Validate request

* Path: `GET /v1/validate`
* Required request header:
    ```
    Authorization: <token>
    ```
* Response body:
    ```json
    {
        "data": "<username extracted from the token>",
        "meta": null
    }
    ```
* Potential error:
    ```json
    {
        "errors": {
            "root": [
                "can not extract data from token"
            ]
        },
        "meta": null
    }
    ```