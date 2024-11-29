package models

// let's define review struct from here
type Review struct {
    ID            string `json:"id"`
    DestinationID string `json:"destination_id"`
    User          string `json:"user"`
    Rating        int    `json:"rating"`
    Comment       string `json:"comment"`
    Date          string `json:"date"`
}
