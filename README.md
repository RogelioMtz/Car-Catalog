# Go Microservice

This project demonstrates a simple web application that interacts with a Go microservice. The application consists of an HTML front-end that makes HTTP requests to a Go back-end service.

## Project Structure

```
html-go-microservice/
├── go-microservice/
│   ├── main.go          # Go server code
│   ├── car_catalog.db   # SQLite database file
├── src/
│   ├── index.html       # Frontend HTML
│   ├── scripts/
│   │   └── app.js       # Frontend JavaScript
│   └── styles.css       # Optional CSS file
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**

   ```
   git clone https://github.com/RogelioMtz/Go-Microservice
   cd html-go-microservice
   ```

2. **Set up the Go microservice:**

   - Navigate to the `go-microservice` directory.
   - Ensure you have Go installed on your machine.
   - Run the microservice:
     ```
     go run main.go
     ```

3. **Open the HTML page:**

   - Open `localhost:8080` or `0.0.0.0:8080` in your web browser.

## Usage

- The HTML page contains a form/button that triggers a call to the Go microservice.
- The JavaScript file (`app.js`) handles the HTTP requests and processes the responses from the microservice.

## Additional Information

- Ensure that the Go microservice is running before interacting with the HTML page.
- Modify the Go handlers in `handler.go` to customize the behavior of the microservice as needed.

## License

This project is licensed under the MIT License.
