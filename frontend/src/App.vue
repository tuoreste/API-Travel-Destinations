<template>
  <div id="app">
    <h1>Global Travel Destinations</h1>
    
    <!-- Continent Selector -->
    <div class="continent-selector">
      <button 
        v-for="continent in continents" 
        :key="continent"
        @click="selectContinent(continent)"
        :class="{ active: selectedContinent === continent }"
      >
        {{ continent }}
      </button>
    </div>
    
    <!-- Destinations Display -->
    <div v-if="loading" class="loading">
      Loading destinations...
    </div>
    
    <div v-else-if="destinations.length" class="destinations-grid">
      <div 
        v-for="destination in destinations" 
        :key="destination.id" 
        class="destination-card"
      >
        <img :src="destination.image_url" :alt="destination.name">
        <div class="destination-info">
          <h2>{{ destination.name }}</h2>
          <p>{{ destination.description }}</p>
          <div class="destination-details">
            <p><strong>Country:</strong> {{ destination.country }}</p>
            <div class="highlights">
              <h3>Highlights:</h3>
              <ul>
                <li v-for="(highlight, index) in destination.highlights" :key="index">
                  {{ highlight }}
                </li>
              </ul>
            </div>
            <div class="actions">
              <a 
                :href="destination.ticket_link" 
                target="_blank" 
                class="ticket-link"
              >
                Book Tickets
              </a>
              <a 
                :href="`https://maps.google.com/?q=${destination.location.latitude},${destination.location.longitude}`" 
                target="_blank" 
                class="map-link"
              >
                View on Map
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div v-else class="no-destinations">
      No destinations found for this continent.
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      continents: [],
      destinations: [],
      selectedContinent: null,
      loading: false
    }
  },
  methods: {
    async fetchContinents() {
      try {
        const response = await axios.get('http://localhost:8090/continents')
        this.continents = response.data
      } catch (error) {
        console.error('Error fetching continents:', error)
      }
    },
    async selectContinent(continent) {
      this.selectedContinent = continent
      this.loading = true
      try {
        const response = await axios.get(`http://localhost:8090/destinations/${continent}`)
        this.destinations = response.data
      } catch (error) {
        console.error('Error fetching destinations:', error)
        this.destinations = []
      } finally {
        this.loading = false
      }
    }
  },
  created() {
    this.fetchContinents()
  }
}
</script>

<style scoped>
.continent-selector {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.continent-selector button {
  margin: 0 10px;
  padding: 10px 15px;
  background-color: #f0f0f0;
  border: none;
  cursor: pointer;
}

.continent-selector button.active {
  background-color: #007bff;
  color: white;
}

.destinations-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 20px;
  max-width: 600px;
  margin: 0 auto;
}

.destination-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
}

.destination-card img {
  width: 100%;
  height: 250px;
  object-fit: cover;
}

.destination-info {
  padding: 15px;
}

.actions {
  display: flex;
  justify-content: space-between;
  margin-top: 15px;
}

.ticket-link, .map-link {
  text-decoration: none;
  padding: 8px 12px;
  background-color: #007bff;
  color: white;
  border-radius: 4px;
}
</style>