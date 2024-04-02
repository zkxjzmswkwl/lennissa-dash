package api

type RunescapeAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Error struct {
	Code    int
	Message string
}
