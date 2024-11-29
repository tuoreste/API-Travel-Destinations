package models

// let's define destination struct from here
type Destination struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Country     string   `json:"country"`
    Description string   `json:"description"`
    Activities  []string `json:"activities"`
    BestSeason  string   `json:"best_season"`
    ImageURL    string   `json:"image_url"`
    Rating      float32  `json:"rating"`
    ReviewCount int      `json:"review_count"`
}
