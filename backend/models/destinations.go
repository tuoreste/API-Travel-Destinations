package models

type Destination struct {
	ID				string `json:"id"`
	Name			string `json:"name"`
	Continent		string `json:"continent"`
	Country			string `json:"country"`
	Description		string `json:"description"`
	Location		struct {
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Ticket_Link		string `json:"ticket_link"`
	Image_URL		string `json:"image_url"`
	Highlights		[]string `json:"highlights"`
	Notification	string `json:"notification"`
}

//raw data ai generated
var Destinations = []Destination{
	//north america
	{
		ID: "1",
		Name: "New York City",
		Continent: "North America",
		Country: "United States",
		Description: "The city that never sleeps, known for its iconic skyline and cultural diversity.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 40.7128,
			Longitude: -74.0060,
		},
		Ticket_Link: "https://www.example.com/nyc-tickets",
		Image_URL: "https://st.depositphotos.com/2578749/2938/i/450/depositphotos_29380709-stock-photo-the-statue-of-liberty-and.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "2",
		Name: "Toronto",
		Continent: "North America",
		Country: "Canada",
		Description: "A vibrant city with diverse cultures and a thriving arts scene.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			}{
				Latitude: 43.6511,
				Longitude: -79.3835,
		},
		Ticket_Link: "https://www.example.com/toronto-tickets",
		Image_URL: "https://images.pexels.com/photos/1519088/pexels-photo-1519088.jpeg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "3",
		Name: "Mexico City",
		Continent: "North America",
		Country: "Mexico",
		Description: "A bustling metropolis with a rich cultural heritage.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 19.4326,
			Longitude: -99.1332,
		},
		Ticket_Link: "https://www.example.com/mexico-city-tickets",
		Image_URL: "https://media.istockphoto.com/id/539002142/photo/mexico-citys-downtown-at-twilight.jpg?s=612x612&w=0&k=20&c=ni3IHGFkUzGVRGXoytCdoe8BdUNyPI1qIKqTkzrG00c=",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "4",
		Name: "Chicago",
		Continent: "North America",
		Country: "United States",
		Description: "Known for its bold architecture and deep-dish pizza.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 41.8781,
			Longitude: -87.6298,
		},
		Ticket_Link: "https://www.example.com/chicago-tickets",
		Image_URL: "https://www.americanholidays.com/wp-content/uploads/2017/02/what-to-do-in-chicago-in-one-day.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "5",
		Name: "Havana",
		Continent: "North America",
		Country: "Cuba",
		Description: "Famous for its vintage cars and rich Cuban culture.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 23.1136,
			Longitude: -82.3666,
		},
		Ticket_Link: "https://www.example.com/havana-tickets",
		Image_URL: "https://as2.ftcdn.net/v2/jpg/01/91/62/59/1000_F_191625972_8WoOVV9YpIJOS1mZz3mCrMAULyNcrAwq.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	//south america
	{
		ID: "6",
		Name: "Rio de Janeiro",
		Continent: "South America",
		Country: "Brazil",
		Description: "Famous for its Carnival, beaches, and the Christ the Redeemer statue.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -22.9068,
			Longitude: -43.1729,
		},
		Ticket_Link: "https://www.example.com/rio-tickets",
		Image_URL: "https://st4.depositphotos.com/2801443/23627/i/450/depositphotos_236273866-stock-photo-aerial-view-christ-redeemer-sugarloaf.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "7",
		Name: "Buenos Aires",
		Continent: "South America",
		Country: "Argentina",
		Description: "A city blending European elegance and Latin passion.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -34.6037,
			Longitude: -58.3816,
		},
		Ticket_Link: "https://www.example.com/buenos-aires-tickets",
		Image_URL: "https://www.shutterstock.com/shutterstock/videos/3522937179/thumb/1.jpg?ip=x480",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "8",
		Name: "Lima",
		Continent: "South America",
		Country: "Peru",
		Description: "A city with stunning colonial architecture and rich history.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -12.0464,
			Longitude: -77.0428,
		},
		Ticket_Link: "https://www.example.com/lima-tickets",
		Image_URL: "https://thumbs.dreamstime.com/b/plaza-de-armas-de-lima-peru-plaza-de-armas-lima-peru-classical-building-lima-town-116999674.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "9",
		Name: "Santiago",
		Continent: "South America",
		Country: "Chile",
		Description: "The vibrant capital surrounded by the Andes Mountains.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -33.4489,
			Longitude: -70.6693,
		},
		Ticket_Link: "https://www.example.com/santiago-tickets",
		Image_URL: "https://www.shutterstock.com/image-photo/santiago-white-cityscape-600nw-1116349802.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "10",
		Name: "Bogotá",
		Continent: "South America",
		Country: "Colombia",
		Description: "The cultural heart of Colombia with vibrant street art.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 4.7110,
			Longitude: -74.0721,
		},
		Ticket_Link: "https://www.example.com/bogota-tickets",
		Image_URL: "https://media.istockphoto.com/id/1453256961/photo/aerial-view-of-modern-bogota-cityscape-in-colombia-in-the-afternoon.jpg?s=612x612&w=0&k=20&c=9rb1gHHo2z8B7nRWlTSiBfJ7vLPziF99qkNz04nzelM=",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	//europe
	{
		ID: "11",
		Name: "Berlin",
		Continent: "Europe",
		Country: "Germany",
		Description: "A vibrant city with a rich history and modern culture.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 52.5200,
			Longitude: 13.4050,
		},
		Ticket_Link: "https://www.example.com/berlin-tickets",
		Image_URL: "https://media.istockphoto.com/id/489776362/photo/berlin-skyline-panorama-with-tv-tower-at-sunset-germany.jpg?s=612x612&w=0&k=20&c=Ng6cNc_9FEtoA1Go3P86ltVsy_-yoAZWcANCQr2ftGs=",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "stuttgart-001",
		Name: "Stuttgart",
		Continent: "Europe",
		Country: "Germany",
		Description: "Stuttgart, the capital of the state of Baden-Württemberg, is known for its rich cultural heritage, modern architecture, and as the home of famous automotive brands like Porsche and Mercedes-Benz.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
		  Latitude: 48.7758,
		  Longitude: 9.1829,
		},
		Ticket_Link: "https://www.stuttgart-tourist.de/en/tickets",
		Image_URL: "https://cdn.pixabay.com/photo/2017/03/01/23/50/stuttgart-2109990_640.jpg",
		Highlights: []string{"Mercedes-Benz Museum","Porsche Museum","Stuttgart State Gallery","Wilhelma Zoological and Botanical Garden","Schlossplatz","Königstrasse (shopping street)"},
		Notification: "You are near Stuttgart. Explore the Mercedes-Benz Museum and other cultural landmarks!",
	},
	{
		ID: "12",
		Name: "Paris",
		Continent: "Europe",
		Country: "France",
		Description: "The city of love and a global hub for art, fashion, and gastronomy.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 48.8566,
			Longitude: 2.3522,
		},
		Ticket_Link: "https://www.example.com/paris-tickets",
		Image_URL: "https://cdn.stocksnap.io/img-thumbs/280h/eiffeltower-france_6UZ3XEKQVA.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "13",
		Name: "Rome",
		Continent: "Europe",
		Country: "Italy",
		Description: "A city steeped in history, home to the Colosseum and Vatican City.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 41.9028,
			Longitude: 12.4964,
		},
		Ticket_Link: "https://www.example.com/rome-tickets",
		Image_URL: "https://cdn.pixabay.com/photo/2019/10/06/08/57/tiber-river-4529605_640.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "14",
		Name: "Amsterdam",
		Continent: "Europe",
		Country: "Netherlands",
		Description: "Known for its canals, museums, and vibrant nightlife.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 52.3676,
			Longitude: 4.9041,
		},
		Ticket_Link: "https://www.example.com/amsterdam-tickets",
		Image_URL: "https://cdn.pixabay.com/photo/2016/01/19/19/26/amsterdam-1150319_640.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "15",
		Name: "Madrid",
		Continent: "Europe",
		Country: "Spain",
		Description: "The lively Spanish capital known for its art and culture.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 40.4168,
			Longitude: -3.7038,
		},
		Ticket_Link: "https://www.example.com/madrid-tickets",
		Image_URL: "https://d2bgjx2gb489de.cloudfront.net/gbb-blogs/wp-content/uploads/2018/08/29184041/Madrid-Spain-Royal-Palace-and-Cathedral-76506930.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	//asia
	{
		ID: "1",
		Name: "Tokyo",
		Continent: "Asia",
		Country: "Japan",
		Description: "The bustling capital of Japan, blending ultramodern skyscrapers with traditional temples.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 35.6895,
			Longitude: 139.6917,
		},
		Ticket_Link: "https://www.example.com/tokyo-tickets",
		Image_URL: "https://www.example.com/images/tokyo.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "2",
		Name: "Bangkok",
		Continent: "Asia",
		Country: "Thailand",
		Description: "Thailand's capital, famous for its vibrant street life and cultural landmarks.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 13.7563,
			Longitude: 100.5018,
		},
		Ticket_Link: "https://www.example.com/bangkok-tickets",
		Image_URL: "https://www.example.com/images/bangkok.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "3",
		Name: "Beijing",
		Continent: "Asia",
		Country: "China",
		Description: "The ancient capital of China, home to the Great Wall and the Forbidden City.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 39.9042,
			Longitude: 116.4074,
		},
		Ticket_Link: "https://www.example.com/beijing-tickets",
		Image_URL: "https://www.example.com/images/beijing.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "4",
		Name: "Dubai",
		Continent: "Asia",
		Country: "United Arab Emirates",
		Description: "A modern city in the UAE, known for its luxury shopping and ultramodern architecture.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 25.276987,
			Longitude: 55.296249,
		},
		Ticket_Link: "https://www.example.com/dubai-tickets",
		Image_URL: "https://www.example.com/images/dubai.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "5",
		Name: "Mumbai",
		Continent: "Asia",
		Country: "India",
		Description: "India's largest city and financial hub, known for Bollywood and cultural diversity.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 19.076,
			Longitude: 72.8777,
		},
		Ticket_Link: "https://www.example.com/mumbai-tickets",
		Image_URL: "https://www.example.com/images/mumbai.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	//africa
	{
		ID: "1",
		Name: "Cairo",
		Continent: "Africa",
		Country: "Egypt",
		Description: "The sprawling capital of Egypt, known for its ancient pyramids and rich history.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 30.0444,
			Longitude: 31.2357,
		},
		Ticket_Link: "https://www.example.com/cairo-tickets",
		Image_URL: "https://www.example.com/images/cairo.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "2",
		Name: "Cape Town",
		Continent: "Africa",
		Country: "South Africa",
		Description: "A port city on South Africa's southwest coast, beneath the iconic Table Mountain.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -33.9249,
			Longitude: 18.4241,
		},
		Ticket_Link: "https://www.example.com/cape-town-tickets",
		Image_URL: "https://www.example.com/images/cape-town.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "3",
		Name: "Nairobi",
		Continent: "Africa",
		Country: "Kenya",
		Description: "Kenya's capital city and a gateway to safari adventures.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -1.2864,
			Longitude: 36.8172,
		},
		Ticket_Link: "https://www.example.com/nairobi-tickets",
		Image_URL: "https://www.example.com/images/nairobi.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "4",
		Name: "Kigali",
		Continent: "Africa",
		Country: "Rwanda",
		Description: "The clean and modern capital city of Rwanda, known for its resilience and growth.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -1.9441,
			Longitude: 30.0619,
		},
		Ticket_Link: "https://www.example.com/kigali-tickets",
		Image_URL: "https://www.example.com/images/kigali.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	{
		ID: "5",
		Name: "Marrakech",
		Continent: "Africa",
		Country: "Morocco",
		Description: "A vibrant city in Morocco known for its historic medina and bustling souks.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: 31.6295,
			Longitude: -7.9811,
		},
		Ticket_Link: "https://www.example.com/marrakech-tickets",
		Image_URL: "https://www.example.com/images/marrakech.jpg",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
}

// model1: function to get destinations by continent
func	GetDestinationsByContinent(continent string) []Destination {
	var result	[]Destination
	for _, dest := range Destinations {
		if dest.Continent == continent {
			result = append(result, dest)
		}
	}
	return result
}

// model1: function to get destinations by ID
func	GetDestinationByID(id string) (Destination, bool) {
	for _, dest := range Destinations {
		if dest.ID == id {
			return dest, true
		}
	}
	return Destination{}, false
}