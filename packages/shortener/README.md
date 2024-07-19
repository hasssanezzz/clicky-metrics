# Shortener service

This API provides functionality to shorten and manage long URLs, it depeneds on the auth service.
## Usage
1. install dependencies
    ```
    pip install -r requirements.txt
    ```
1. Configure your user data storage mechanism (e.g., database) in the relevant code.
1. Set the required environment variables.
    ```
    POSTGRES_HOST=...
    POSTGRES_PORT=...
    POSTGRES_USER=...
    POSTGRES_PASSWORD=...
    POSTGRES_DB=...
    AUTH_SERVICE_HOST=...
    AUTH_SERVICE_PORT=...
    ```
1. Run the main application to start the API server.
    ```
    flask run
    ```

### Authentication

Some endpoints require a valid authorization token to be included in the request header under the key `AUTH_TOKEN_NAME` (defined in the constants file). The `AuthService` class is responsible for validating the token and retrieving the associated username.

### Endpoints Table

| Method | URI           | Description                                                 | Authentication                    |
| ------ | ------------- | ----------------------------------------------------------- | --------------------------------- |
| GET    | `/v1`         | Retrieves all shortened URLs for the authenticated user.    | Required                          |
| POST   | `/v1`         | Creates a new shortened URL for a provided long URL.        | Not Required                      |
| GET    | `/v1/<short>` | Redirects to the long URL associated with the short code.   | Not Required                      |
| PUT    | `/v1/<short>` | Updates the long URL associated with a specific short code. | Required (ownership verification) |
| DELETE | `/v1/<short>` | Deletes a shortened URL.                                    | Required (ownership verification) |

**Note:**

- `<short>` represents the unique short code generated for the long URL.
- Ownership verification in PUT and DELETE requests ensures that a user can only modify or delete URLs they have created.

## Endpoints

* Get all user's shortened URLs
    - Path `GET /v1`
    - Authentication: **REQUIRED**
    - Success response:
        ```json
       {
            "data": [
                {
                    "id": ...,
                    "user_username": ...,
                    "short": "Voopxb",
                    "long": "http://colorhunt.co"
                },
                {
                    "id": ...,
                    "user_username": ...,
                    "short": "JcDwqe",
                    "long": "http://x.com"
                },
                {
                    "id": ...,
                    "user_username": ...,
                    "short": "QQIcuh",
                    "long": "http://google.com"
                },
                {
                    "id": ...,
                    "user_username": ...,
                    "short": "ybtp1g",
                    "long": "http://stackoverflow.com"
                }
            ],
            "meta": null
       }
        ```
    - Errors:
        * 401 - unauthorized
        ```json
        {
            "error": {
                "root": ["unauthorized"]
            },
            "meta": null
        }
        ```
* Redirect
    - Path `GET /v1/{short}`
    - Authentication: **NOT REQUIRED**
    - Errors:
        * 404 - not found
        ```json
        {
            "error": {
                "short": ["short not found"]
            },
            "meta": null
        }
        ```
* Shorten a URL
    - Path `POST /v1`
    - Authentication: **NOT REQUIRED**
    - Request body:
        ```json
        { "url": "<a valid URL>" }
        ```
    - Success response:
        ```json
        {
            "data": {
                "id": ...,
                "user_username": ...,
                "short": "KfkfAw",
                "long": "http://colorhunt.co"
            },
            "meta": null
        }
        ```
    - Errors:
        * 500 internal server error
        ```json
        {
            "error": {
                "root": ["internal server error"]
            },
            "meta": null
        }
        ```
- Update a URL
    - Path `PUT /v1/{short}`
    - Authentication: **REQUIRED**
    - Request body:
        ```json
        { "url": "<a valid URL>" }
        ```
    - Success response:
        ```json
        {
            "data": {
                "id": ...,
                "user_username": ...,
                "short": "KfkfAw",
                "long": "http://colorhunt.co"
            },
            "meta": null
        }
        ```
    - Errors:
        * 401 - unauthorized
        ```json
        {
            "error": {
                "root": ["unauthorized"]
            },
            "meta": null
        }
        ```
        * 404 - not found
        ```json
        {
            "error": {
                "short": ["short not found"]
            },
            "meta": null
        }
        ```
- Delete a URL
    - Path `PUT /v1/{short}`
    - Authentication: **REQUIRED**
    - Success response:
        ```json
        {
            "data": null,
            "meta": null
        }
        ```
    - Errors:
        * 401 - unauthorized
        ```json
        {
            "error": {
                "root": ["unauthorized"]
            },
            "meta": null
        }
        ```
        * 404 - not found
        ```json
        {
            "error": {
                "short": ["short not found"]
            },
            "meta": null
        }
        ```
