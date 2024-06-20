package model

type JWT struct {
	AccessToken  string
	RefreshToken string
	ExpiredIn    int
}
