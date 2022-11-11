package entity

type Favorite struct {
	ID       string `json:"id"`
	IDHash   string `json:"idHash"`
	Name     string `json:"name"`
	IsPublic bool   `json:"isPublicFavorite"`
}
