# Travel Destinations API

The **Travel Destinations API** is a full-stack application designed to provide information about various travel destinations. It includes a backend built with Go and a frontend built with Vue.js. The API allows users to explore destinations by continent, retrieve details by ID, find nearby destinations based on geolocation, and receive real-time notifications using WebSocket.

---

## Usage

### Backend
The backend serves as the API provider, handling requests and sending data. It runs on **`http://localhost:8080`**.

### Frontend
The frontend provides a user interface for interacting with the API, such as searching destinations and viewing real-time notifications. It runs on **`http://localhost:8090`**.

### Available Features
1. **Explore Destinations**:
   - Search destinations by continent.
   - View details about a destination, including highlights and a ticket link.
<!-- 2. **Nearby Destinations**:
   - Find destinations within a specified radius of your current location. -->
2. **Real-time Notifications**:
   - Receive alerts about nearby destinations as you move using WebSocket.

### Backend API Endpoints
#### Example Endpoints:
- `GET /destinations/continent/{continent}` - Fetch destinations by continent.
- `GET /destinations/{id}` - Fetch destination details by ID.
- `GET /destinations/nearby?latitude={lat}&longitude={lon}&radius={radius}` - Get nearby destinations.
- WebSocket: `ws://localhost:8080/geo-tracking` - Real-time location notifications.

#### Frontend Access:
- Open your browser and navigate to **`http://localhost:8090`**.
- Use the UI to explore features such as searching destinations and receiving notifications.

---

## Tools Used

### Backend:
- **Go**:
  - RESTful API design
  - WebSocket integration for real-time communication
- **Gorilla Mux**:
  - Routing and HTTP request handling
- **WebSocket**:
  - Real-time updates based on geolocation

### Frontend:
- **Vue.js**:
  - Reactive UI for interacting with the API
- **Axios**:
  - Handling HTTP requests to communicate with the backend
- **CSS**:
  - Styling the application for a clean user experience

### Other Tools:
- **Curl**: Testing API endpoints
- **Git/GitHub**: Version control and collaboration

---

## Technicalities Involved

1. **Geolocation Handling**:
   - The backend calculates the distance between the user's location and destinations using the Haversine formula.
   - Real-time updates are sent using WebSocket when the user is near a destination.

2. **Frontend Communication**:
   - The Vue.js frontend interacts with the backend via REST API calls (using Axios).
   - WebSocket integration provides instant notifications in the UI.

3. **Cross-Origin Resource Sharing (CORS)**:
   - Configured on the backend to allow the frontend to access API endpoints.

4. **Scalability**:
   - Backend designed with a modular structure, enabling future integration with a database for dynamic content.

5. **User Experience**:
   - Responsive frontend interface to enhance usability.
   - Real-time alerts for improved engagement.

---
### Clone repository

1. Clone
```bash
git clone https://github.com/tuoreste/API-Travel-Destinations.git 
```
2. Navigate to project directory
```bash
cd API-Travel-Destinations
```

### How to Run

#### Backend:
1. Navigate to the backend directory.
2. Run:
   ```bash
   go run main.go
   ```
3. Backend will be available at ```http://localhost:8080.```

#### Frontend:

1. Navigate to the frontend directory.
```bash
cd ../frontend
```
2. Install dependencies:
```bash
npm install
```
3. Run:
```bash
npm run dev
```
4. Frontend will be available at ```http://localhost:8090```

## Next Updates

Here are the proposed updates to enhance the **Travel Destinations API**:

1. **Move the Current Raw Data to a Database**  
   - Transition the hardcoded data to a relational or NoSQL database (e.g., Mariadb).  
   - Implement models and data access layers to handle CRUD operations efficiently.  
   - Improve scalability and dynamic data updates.

2. **Automate Database Population**  
   - Develop a scraper or integrate APIs to gather travel destination data from popular websites (e.g., TripAdvisor, Google Places).  
   - Automate periodic updates to keep the database current with trending destinations and events.

3. **Enhance the Frontend with Maps and Events**   
   - Highlight details for each destination, such as:  
     - Type of events happening.  
     - Distance from the user.  
     - Directions to the destination.  
   - Use dynamic styling to make the frontend more engaging and user-friendly.

4. **Containerize the Project**  
   - Split the project into three microservices:  
     - **Frontend**: Vue.js-based service for the user interface.  
     - **Backend**: Go-based service for APIs and real-time notifications.  
     - **Database**: Service hosting the chosen database (Mariadb).  
   - Use Docker to containerize each microservice for isolated and consistent deployment.

5. **Deploy on an Orchestrator**  
   - Deploy the containerized microservices into a Kubernetes cluster.  
   - Implement horizontal scaling for the backend and database services.  
   - Use tools like Helm charts for configuration management.  
   - Leverage monitoring solutions (e.g., Prometheus, Grafana) for operational insights.  

These updates will improve the scalability, usability, and maintainability of the project while making it ready for production-grade deployment.

## END
