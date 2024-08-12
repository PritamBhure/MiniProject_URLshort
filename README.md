
# URL Shortener Project

This project helps to shorten long URLs.

### Initial Setup

Before running this project, ensure you have completed the following steps:

1. **Go Language Setup:** Install the Go programming language for your operating system.
2. **Go Extension in VS Code:** If you are using Visual Studio Code (VS Code), make sure the Go extension is installed.
3. **Go Module Initialization:** If you are creating files outside the Go language path (`src`), initialize your Go module by running the command:

   ```bash
   go mod init <module-name>
   ```

### Steps to Run the Project

**Step 1:**

Run the code using the command line:

```bash
go run main.go
```

After running this command, the server will start on your localhost on port 8000.

**Step 2:**

Send a request to the server using the API endpoint `http://localhost:8000/shorten`. Use the following format:

**GET Request:**

```http
http://localhost:8000/shorten
```

**Request Body:**

```json
{
  "URL": "www.example.com"  // Replace with the URL you want to shorten
}
```

**Step 3:**

You will see the response in the standard output window of your IDE. The response will look like this:

```json
"Short_URL": "xyz"
```

**Step 4:**

You can redirect to the original URL using the shortened URL. Use this command:

```http
http://localhost:8000/redirect/xyz  // Replace "xyz" with the code you received
```

Before running this project, ensure you have completed the following steps:

1. **Go Language Setup:** Install the Go programming language for your operating system.
2. **Go Extension in VS Code:** If you are using Visual Studio Code (VS Code), make sure the Go extension is installed.
3. **Go Module Initialization:** If you are creating files outside the Go language path (`src`), initialize your Go module by running the command:

   ```bash
   go mod init <module-name>
   ```

### Steps to Run the Project

**Step 1:**

Run the code using the command line:

```bash
go run main.go
```

After running this command, the server will start on your localhost on port 8000.

**Step 2:**

Send a request to the server using the API endpoint `http://localhost:8000/shorten`. Use the following format:

**GET Request:**

```http
http://localhost:8000/shorten
```

**Request Body:**

```json
{
  "URL": "www.example.com"  // Replace with the URL you want to shorten
}
```

**Step 3:**

You will see the response in the standard output window of your IDE. The response will look like this:

```json
"Short_URL": "xyz"
```

**Step 4:**

You can redirect to the original URL using the shortened URL. Use this command:

```http
http://localhost:8000/redirect/xyz  // Replace "xyz" with the code you received
```

You will see the response on the standerd output Windown of your IDE
like this: "Short_URL : xyz"

step 4 
You can redirect to that URL origin throut that short URL 
Usine this Cmd http://localhost:8000/redirect/xyz(Code that you get)
