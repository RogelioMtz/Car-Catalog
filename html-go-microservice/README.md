# HTML Go Microservice

This project demonstrates a simple web application that interacts with a Go microservice. The application consists of an HTML front-end that makes HTTP requests to a Go back-end service.

## Project Structure

```
html-go-microservice
├── src
│   ├── index.html        # Main HTML page
│   └── scripts
│       └── app.js       # JavaScript for handling interactions
├── go-microservice
│   ├── main.go          # Entry point for the Go microservice
│   └── handlers
│       └── handler.go   # Request handlers for the microservice
├── README.md            # Project documentation
└── .gitignore           # Files to ignore in version control
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
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
   - Navigate to the `src` directory.
   - Open `index.html` in your web browser.

## Usage

- The HTML page contains a form/button that triggers a call to the Go microservice.
- The JavaScript file (`app.js`) handles the HTTP requests and processes the responses from the microservice.

## Additional Information

- Ensure that the Go microservice is running before interacting with the HTML page.
- Modify the Go handlers in `handler.go` to customize the behavior of the microservice as needed.

## License

This project is licensed under the MIT License.