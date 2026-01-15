# YouTube Bingo Web Application

## Overview
This project is a web application that allows users to play Youtube Bingo by randomly selecting items from a predefined list. The application is built using Go and serves a simple web interface for user interaction.

## Project Structure
```
youtube-bingo-web
├── cmd
│   └── server
│       └── main.go          # Entry point of the web application
├── internal
│   ├── handlers
│   │   └── bingo.go         # HTTP handler functions for bingo functionality
│   └── bingo
│       └── bingo.go         # Core logic for the bingo game
├── web
│   ├── static
│   │   └── style.css        # CSS styles for the web application
│   └── templates
│       └── index.html       # Main HTML template for user interface
├── list.txt                 # List of items for the bingo game
├── go.mod                   # Module definition for the Go project
└── README.md                # Documentation for the project
```

## Setup Instructions
1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd youtube-bingo-web
   ```

2. **Install Dependencies**
   Ensure you have Go installed. Run the following command to download the necessary dependencies:
   ```bash
   go mod tidy
   ```

3. **Create the List File**
   Populate `list.txt` with the items you want to use for the bingo game, each item on a new line.

4. **Run the Application**
   Navigate to the `cmd/server` directory and run:
   ```bash
   go run main.go
   ```

5. **Access the Application**
   Open your web browser and go to `http://localhost:8080` to access the bingo game.

## Usage
- Click the button on the web interface to randomly select an item from the list.
- The selected item will be displayed on the screen.

## Contributing
Feel free to submit issues or pull requests for improvements or bug fixes.