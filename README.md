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
./sertile <port> <offline-map-folder-path>
```

The server will start and listen on the specified port, serving tile and terrain data.

## API Endpoints

The Offline Maps Server exposes the following API endpoints:

- `GET /{tilePattern}`: Retrieves tile data based on the given pattern.
- `GET /{terrainPattern}`: Retrieves terrain data based on the given pattern.
- `GET /terrain/layer.json`: Retrieves the JSON file containing layer information.
