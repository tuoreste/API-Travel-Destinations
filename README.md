# Travel Destinations API

The **Travel Destinations API** is a Go-based RESTful API designed to provide information about travel destinations around the world. It supports fetching destinations by continent, retrieving details by ID, and identifying nearby destinations based on a user's geolocation. Additionally, it includes real-time notifications using WebSocket technology.

---

## Features

- **Retrieve Destinations by Continent**: Fetch all destinations for a specific continent.
- **Get Destination Details by ID**: Retrieve detailed information for a destination using its unique ID.
- **Nearby Destinations**: Calculate and return destinations within a given radius of the user's geolocation.
- **Real-time Notifications**: Use WebSocket for live updates and notifications when near a destination.
- **Built-in Data**: Includes a sample dataset with destinations across North America, Europe, and more.

---

## Endpoints

### 1. **Get Destinations by Continent**
- **`GET /destinations/continent/{continent}`**
- Returns a list of destinations for a specified continent.
- **Example**: `/destinations/continent/North America`

### 2. **Get All Continents**
- **`GET /continents`**
- Returns a list of all available continents.

### 3. **Get Destination by ID**
- **`GET /destinations/{id}`**
- Returns detailed information about a destination by its unique ID.
- **Example**: `/destinations/1`

### 4. **Get Nearby Destinations**
- **`GET /destinations/nearby?latitude={lat}&longitude={lon}&radius={radius}`**
- Returns destinations within a specific radius of the provided geolocation.

### 5. **Real-time Geo-tracking WebSocket**
- **`ws://{server}/geo-tracking`**
- Establishes a WebSocket connection for real-time notifications based on user location.

---

## Tech Stack

- **Go**: The main programming language for API development.
- **Gorilla Mux**: For routing HTTP requests.
- **WebSocket**: Used for real-time communication.
- **JSON**: Used for data serialization.

---

## Installation and Setup

### Prerequisites
- **Go** installed on your system (version 1.19 or later recommended).

### Steps

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/travel-destinations-api.git
   cd travel-destinations-api
   ```
2. Install Dependencies:
   ```bash
   go mod tidy
   ```
3. Run the API:
   ```bash
   go run main.go
   ```
4. Access the API: The API will be running on http://localhost:8080.

   
