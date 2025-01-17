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
		Image_URL: "https://encrypted-tbn3.gstatic.com/licensed-image?q=tbn:ANd9GcReaGmVP7LGNkAzSHREsIWhP9dpzLfgmZ5kwoM9l3xGARLlxd0hAJE6QD_O23T7UxNLUEMSeRexmg20YDLiQ6e5nEpcxy5mK7U0_yZW6A",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcSFOMW4-rz7MwjHa9eyEQawDogLijZ-1qaq9iDA9gKbN50r0uzvC0W3M9XYdOLwIi54jxFaLCj5Akxtwoch3eVgebONm86kGAbf_OMcMXo",
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
		Image_URL: "https://encrypted-tbn1.gstatic.com/licensed-image?q=tbn:ANd9GcSvdVCisbqGzdKJ2O5PWHa_pXaEn5KupdG45CjufswdEmeDfJHe3APSXqN4caGPanrpzgqvG0RaMxNzJIdgQoLNgi2Ixi00eNwH9-tr_w",
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
		Image_URL: "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcS2zyXSmr9B8AyfUmJA0R3q0iDiIeFuC1pTu1NtmxV3rr_MUmOyTc8wy6ygRvAI3ABHJzbSNqdZaX5Y9dA8-L1pxpDyEyK8_Hiy8oF-nA",
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
		Image_URL: "https://lh4.googleusercontent.com/proxy/3f6Pd_TASeyng5Gx5H6HzOKOPSm8bW_LLB-NaR0r_eWTsIRa3dp3vZ37L5xOHPQ8ZJECBmQcQ10MH1J779ZjO6lWwNmvMyWLBER2_8fhNY_C0WDMMScaTf04jlZkCQvv9hCmn3lMtiVB-dYzanPNrhxq7VKAgw=w1080-h624-n-k-no",
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
		Image_URL: "https://lh5.googleusercontent.com/p/AF1QipPUTncdgQzogJ1bTDMPJJX3nqQjs0FZ4iDpf-he=w1080-h624-n-k-no",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcSbP6Fok0EpVm_sXbKoO8FHg5yf_uzSWtKbm9m1aHLMfonp0SAp9QME0epDDi-5wcwpunVSDyR2gmk8Ykh4KXZn8yDJ3CUZVXG3tz9LBw",
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
		Image_URL: "https://encrypted-tbn3.gstatic.com/licensed-image?q=tbn:ANd9GcRYwPWTmdSyuYsgr-7gygWg3faLwUjaqzciKItLNJ8tZVAaqG2wMXb_-IOyVYlLdMOPM4Rro13ef6BKGbWzaQ7WoODTTpJpqPfyNe5FglY",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcTWc824LPV3x4rpBqavB3djXzsgpHv2umvO-dYEQEmZYIfKdv0PlS4msjne9mzhGczvaNWAgifBuOSG1a7bKIRWGIJwWcXGrjQWCb36IQ",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcTO-CRv9VDmuMaf5N2rUgRJkEgpbbfdwWvO-_cyT-hwjGI1fPy5WcX-s7N85RiX0r26FNRUZOfK6saHb1oHLzHRJaBA_p3Sdpz8BMFSvgg",
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
		Image_URL: "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcQrAkEc5BUWWMZldToRjaBuqrsRnJmQyeQXapvHCKyh8bnkc4XY8Xy1gPmxhCtFP4cfHLRD6l8qUsp1FJ-MyGdYXMo24izj0GiGiE5qcw",
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
		Image_URL: "https://lh5.googleusercontent.com/p/AF1QipOlbVTgmhVBe69aEoxlnLNVFd4O94MaKLGi2L1i=w1080-h624-n-k-no",
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
		Image_URL: "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcSGF5DWDGW3Km0LY80llTRIPvnLN20Vh-3God4ik4TLxmWpxNpKMgPFJ97_MamXKvshed-npZC_CkQwaahQrd-gNvkYC1vddLxZwVeX3g",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcSh0YONJqCBmZgw-BwJPqo1nUrxuPVJA4bKZCftN75p4dGBrWIqZxtwGWzWWaBN1KnXDKGB692c6CjQajuh7RWetOGlEik-jbE8yHt5AQ",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcQiofXgjVcndS7w_0GGaZj8ZwrY54nRnrPCYwt60iB3WMNTBzpNxQ5GNdd3XXx1960ev3MV63jNe9Vdm8xtS8F5_iqKQoYbJ2O524Bsfg",
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
		Image_URL: "https://encrypted-tbn1.gstatic.com/licensed-image?q=tbn:ANd9GcSs6dFJiJugupPp7qdD8Ra8iWq9wxcMeRxAusipYIZf0c5XWy5rrGz_ECpujCUO2F7bF4ydBzP8Hl3g61q0MbC6EEoIWSPKIfeoMENy-w",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcQF9sBKWwybUuqYa96UYnDmRCKawmKRhKhxGOubI1O4gTaP78JZTSN1xvtssrKvcGNkWaTldx-jeVpXJzO6ttgeaoPty66adoYBbhUvOg", 
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcRDbCsS-grz_WRffHbrPSh6cBMvXW1I06m5F6Hvj1WrfoSyfFXuzzA9gCwrLRW_FlSq9RdlGiQ-Um1xh521gAqDC6dOqT_vndQyTLutVw",
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
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcT06eOer0paFAdNIsveeqQjH1QmE1EcBeBSTBUtv876i4R4LgGid9yyOT5Rry0rPlSkXc9R2EDEz7tiLFftSmPkm4rtCPJHRp79CR-Tvg",
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
		Image_URL: "https://encrypted-tbn1.gstatic.com/licensed-image?q=tbn:ANd9GcQOun7PTUSZ-ZzzU4ek059lF-i53hZ6Lq7Q1HkW3QBbBi3MzXKTXXvij5nZiNkUA9Jy1Le837kOJi6OkAvHhE5H1pgKV1gijGpW1MhC0w",
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
		Image_URL: "https://encrypted-tbn3.gstatic.com/licensed-image?q=tbn:ANd9GcThSjt5H6dWXQy2iAZYRgOaTVi0SdmgrsT117_u-RPXJZoYjIt62d2jO2DNKuAlg5S6VL_bn2L6aVTbWcgBjxW1SMhodENhETctvS3W8A",
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
		Image_URL: "https://encrypted-tbn1.gstatic.com/licensed-image?q=tbn:ANd9GcQ5RII7DjZwcdEKu3blOYwtXwBy2pI6wdMxCtZLQLFHvYFbHcyr0CBnEBdcFaUJBUlUIpvyy0CztbnkK7Iw9Cn7unM5Z5teHD7z0a3Fwg",
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
		Image_URL: "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcTrIZ6YWnCMd8x_4v284QQxBg1V3BHwZ4rOobnudRBgP16vVWnnUKjHHidlBd1bYa4O8K3PXzbz4Vrr7cK3OC7vnYbTDUKPH76qH5BGDQ",
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
		Image_URL: "https://encrypted-tbn1.gstatic.com/licensed-image?q=tbn:ANd9GcRAWYBxh-h-QU-1PDsjpEoNQMT6Ww4JmivSbFpmEv7aRcr0o4_qhM_muGOXxxVMpkB-3xO9R0gNXz3X4ucnw1QidawEA7uXf4gkz_Yrdg",
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
		Image_URL: "https://lh5.googleusercontent.com/p/AF1QipPYwEsS93zLflrAVzqqmCg0M_xmk5aXtuqvOuji=w1080-h624-n-k-no",
		Highlights:  []string{"Chichen Itza", "Beach Resorts", "Underwater Museums"},
		Notification: "You are near Empire State Building. Check out the highlights: Iconic Views, Historic Architecture",
	},
	// Australia
	{
		ID: "1",
		Name: "Sydney",
		Continent: "Australia",
		Country: "Australia",
		Description: "Australia's largest city, known for its stunning harbor, iconic Opera House, and beautiful beaches.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -33.8688,
			Longitude: 151.2093,
		},
		Ticket_Link: "https://www.example.com/sydney-tickets",
		Image_URL: "https://lh5.googleusercontent.com/p/AF1QipMHftgSCBlvyjxYphi4gLqDC_62WWvZvyy1EBuh=w1080-h624-n-k-no",
		Highlights: []string{"Sydney Opera House", "Sydney Harbour Bridge", "Bondi Beach"},
		Notification: "You are near Sydney Opera House. Check out the highlights: Iconic Views, Stunning Architecture",
	},
	// Antarctica
	{
		ID: "1",
		Name: "McMurdo Station",
		Continent: "Antarctica",
		Country: "United States",
		Description: "The largest research station in Antarctica, serving as a hub for scientific research and exploration.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -77.8419,
			Longitude: 166.6863,
		},
		Ticket_Link: "https://www.example.com/mcmurdo-station-tickets",
		Image_URL: "https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/McMurdo_2.12.13_a-HR.jpg/2880px-McMurdo_2.12.13_a-HR.jpg",
		Highlights: []string{"Research Facilities", "Scientific Exploration", "Penguin Watching"},
		Notification: "You are near McMurdo Station. Check out the highlights: Research Stations, Penguins",
	},
	{
		ID: "2",
		Name: "Vostok Station",
		Continent: "Antarctica",
		Country: "Russia",
		Description: "A remote Russian research station, located near the South Pole, known for its ice cores and extreme cold.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -78.4640,
			Longitude: 106.8330,
		},
		Ticket_Link: "https://www.example.com/vostok-station-tickets",
		Image_URL: "https://upload.wikimedia.org/wikipedia/commons/c/c3/Vostok_Station_2024.png",
		Highlights: []string{"Ice Core Drilling", "Extreme Cold Weather Research", "South Pole Access"},
		Notification: "You are near Vostok Station. Check out the highlights: Ice Research, Cold Weather Testing",
	},
	{
		ID: "3",
		Name: "Amundsen-Scott South Pole Station",
		Continent: "Antarctica",
		Country: "United States",
		Description: "The research station at the geographic South Pole, crucial for scientific studies on climate, astronomy, and geology.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -90.0000,
			Longitude: 0.0000,
		},
		Ticket_Link: "https://www.example.com/amundsen-scott-station-tickets",
		Image_URL: "https://upload.wikimedia.org/wikipedia/commons/8/88/South_Pole_Dome_From_Station.JPG",
		Highlights: []string{"South Pole Studies", "Astronomical Observatories", "Climate Research"},
		Notification: "You are near Amundsen-Scott South Pole Station. Check out the highlights: South Pole Studies, Extreme Conditions",
	},
	{
		ID: "4",
		Name: "Palmer Station",
		Continent: "Antarctica",
		Country: "United States",
		Description: "A U.S. research station on the Antarctic Peninsula, focusing on marine biology and glaciology.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -64.7732,
			Longitude: -64.0506,
		},
		Ticket_Link: "https://www.example.com/palmer-station-tickets",
		Image_URL: "https://upload.wikimedia.org/wikipedia/commons/3/36/PalmerFromGlaciar.JPG",
		Highlights: []string{"Marine Biology Research", "Glaciology Studies", "Penguin and Seal Observation"},
		Notification: "You are near Palmer Station. Check out the highlights: Marine Biology, Wildlife Observation",
	},
	{
		ID: "5",
		Name: "Casey Station",
		Continent: "Antarctica",
		Country: "Australia",
		Description: "An Australian research station, focusing on environmental science, biology, and climate change studies.",
		Location: struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}{
			Latitude: -66.2818,
			Longitude: 110.5272,
		},
		Ticket_Link: "https://www.example.com/casey-station-tickets",
		Image_URL: "https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcT2iYicQWxdf0UXpY_QeTZZpJX2-ZUkd6TTpQMq94Y3myqh38uYNF4wqyEGMS8YLfcxIUrWzYj1t_aoWHyiFwHqUgLHwSdI0a1y177uYw",
		Highlights: []string{"Environmental Science", "Climate Change Research", "Antarctic Ecosystem Studies"},
		Notification: "You are near Casey Station. Check out the highlights: Environmental Studies, Climate Research",
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