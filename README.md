# Description
Ascii-art-web is a web server application written in Go that allows users to generate ASCII art based on the provided text input and selected banner style. The application provides a user-friendly graphical interface for generating ASCII art.

# Authors
- akurmana
- yilgekba

# Usage: how to run
To run the application, follow these steps:
1. Clone the repository to your local machine.
2. Navigate to the root directory of the project.
3. Run the following command in your terminal:
    ```
    go run server.go
    ```
4. Once the server is running, open your web browser and go to `http://localhost:8080` to access the application.

# Implementation details: algorithm
The application consists of two main handlers:
- `mainPageHandler`: This handler serves the main HTML page to the user. It determines the appropriate template path based on the requested URL and renders the corresponding HTML template.
- `generateASCIIArtHandler`: This handler processes the user input for generating ASCII art. It validates the input text format, checks the selected banner style file for correctness, generates the ASCII art, and renders the result in the main HTML template.

The application uses the following components:
- HTML templates: Located in the `templates` directory, these templates define the structure and content of the web pages.
- Banner style files: Located in the `styles` directory, these files define different styles for ASCII art banners.
- Error pages: Located in the `templates` directory, these HTML files display custom error messages for different HTTP error statuses.

The application handles HTTP GET and POST requests for various functionalities, such as rendering the main page, generating ASCII art, and displaying error pages for different error statuses.

The `utils` package provides utility functions for checking the input text format and validating the correctness of the selected banner style file.

The application follows good practices for error handling, logging, and serving static files, ensuring robustness and reliability in its operation.
