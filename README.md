# HTTP FROM SCRATCH IN GOLANG

### 1. **GET / (Main Page)**

- **Handler**: `mainPageHandler`
- **Response**: It returns a simple HTML page with a "Hello World" message.
- **Command**: `curl http://localhost:6969/`
- **Output**: `<h1>Hello World</h1>`

### 2. **GET /echo**

- **Handler**: `echoHandler`
- **Response**: It returns whatever is in the path after `/echo/`. If you send `/echo/anything`, it will return `"anything"`. If you don't provide anything after `/echo/`, it returns a 404 Not Found response.
- **Command**: `curl http://localhost:6969/echo/some-text`
- **Output**: `some-text`

### 3. **GET /user-agent**

- **Handler**: `userAgentHandler`
- **Response**: It returns the `User-Agent` header value from the request. If the `User-Agent` header is not found, it returns a 400 error.
- **Command**: `curl -H "User-Agent: myAgent" http://localhost:6969/user-agent`
- **Output**: `myAgent`

### 4. **GET /files**

- **Handler**: `filesGetHandler`
- **Response**: It returns the contents of a file from the server. You must specify the file name after `/files/`. The file's directory path is determined by the second command-line argument passed when starting the server. If the file doesn't exist, it returns a 404 error.
- **Command**: `curl http://localhost:6969/files/filename.txt`
- **Output**: The contents of `filename.txt` from the server (if it exists).

### 5. **POST /files**

- **Handler**: `filesPostHandler`
- **Response**: It saves the request body to a file on the server. The file name is specified after `/files/` in the path. The directory path is again based on the second command-line argument.
- **Command**: `curl -X POST -d "Some content" http://localhost:6969/files/filename.txt`
- **Output**: `saved` if the file is successfully created or updated.

### Possible HTTP Methods and Paths:

- **GET /**: Main page with a "Hello World" message.
- **GET /echo/{text}**: Echoes the `{text}` provided in the URL path.
- **GET /user-agent**: Returns the `User-Agent` header from the request.
- **GET /files/{filename}**: Retrieves the contents of `{filename}` from the file system.
- **POST /files/{filename}**: Saves the body content to a file named `{filename}` on the server.

### How to Run the Server:

1. You need to run this Go program, making sure to specify a directory path for the `files` handler if needed:
    
    ```bash
    go run main.go /path/to/files
    
    ```
    
2. Then, you can use `curl` commands as outlined above to interact with the server.
