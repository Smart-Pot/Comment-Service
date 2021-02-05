## Comment Struct

```json
Comment{
    "id": string,
    "userId": string,
    "postId": string,
    "content":string,
    "date":string,
}
```

## GetByUser

- path: `/comment/user/{id}`

- returns:
    ```js
    {
        "comments": Comment[],
        "success": Number,
        "message" : String
    }
    ```


## GetByPost

- path : `/comment/post/{id}`
- method: `DELETE`
- returns:
    ```js
    {
        "comments": Comment[],
        "success": Number,
        "message" : String
    }
    ```


## Add
- path: `/comment/new`
- method: `POST`
- params:
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |

- returns:
    ```js
    {
        "comments": Comment[],
        "success": Number,
        "message" : String
    }

## Delete 
- path: `/comment/{id}`
- method: `DELETE`
- params:
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |
- returns:
    ```js
    {
        "comments": Comment[],
        "success": Number,
        "message" : String
    }

## Vote
- path: `/comment/vote`
- method: `POST`
- params:
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |

- returns:
    ```js
    {
        "comments": Comment[],
        "success": Number,
        "message" : String
    }
