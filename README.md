Certainly! Below is a sample README file for your Go project based on the provided `main.go` code. The README provides an overview of your project, explains the main components, and guides users on how to set up and run the application.

---

# Offline Maps Server

This is an Offline Maps Server written in Go. The server serves tile and terrain data for offline use in map applications.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Overview

The Offline Maps Server is a Go application that serves tile and terrain data for offline map applications. It exposes HTTP endpoints for retrieving tiles and terrain data based on the given parameters.

## Installation

To run the Offline Maps Server on your local machine, you need to have Go installed. Follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/offline-maps-server.git
   cd offline-maps-server
   ```

2. Install dependencies:

   ```sh
   go get github.com/gorilla/mux
   ```

3. Build the project:

   ```sh
   go build
   ```

## Usage

After building the project, you can run the application:

```sh
./offline-maps-server
```

The server will start and listen on the specified port, serving tile and terrain data.

## Configuration

The application uses command-line arguments for configuration. You can specify options such as the port number to listen on. Here's how you can pass arguments:

```sh
./offline-maps-server --port 8080
```

For more configuration options, refer to the [Configuration](#configuration) section.

## API Endpoints

The Offline Maps Server exposes the following API endpoints:

- `GET /{tilePattern}`: Retrieves tile data based on the given pattern.
- `GET /{terrainPattern}`: Retrieves terrain data based on the given pattern.
- `GET /terrain/layer.json`: Retrieves the JSON file containing layer information.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize the content and structure of the README according to your project's specific details and needs. This template provides a basic outline to help users understand the purpose of your project, how to set it up, and how to use it.
