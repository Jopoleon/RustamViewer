package models

type Session struct {
	LoggedIn     bool
	UserID       int
	SessionToken string
}
