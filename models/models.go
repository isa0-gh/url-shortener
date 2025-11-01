package models

type Body struct {
	Url    string `json:"url" binding:"requried"`
	Expire int    `json:"expire" binding:"required"` // Expire seconds if nothing gived it will never expire
}

type NewUrl struct {
	DeleteId string `json:"delete_id"`
	ShortId  string `json:"short_id"`
}
