package main

type Post struct {
	ID        string `db:"id" json:"id"`
	Body      string `db:"body" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type User struct {
	ID              string `db:"id" json:"id"`
	Name            string `db:"username" json:"name"`
	CreatedAt       string `db:"created_at" json:"created_at"`
	UpdatedAt       string `db:"updated_at" json:"updated_at"`
	ProfileImageURL string `db:"profile_image_url" json:"profile_image_url"`
	Bio             string `db:"bio" json:"bio"`
}
