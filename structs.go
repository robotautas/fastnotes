package main

type Note struct {
	Title string
	Body  string
	Tags  []Tag
}

type Tag struct {
	Title string
	Notes []Note
}

type Page struct {
	Notes []Note
}

type Login struct {
	Email    string
	Password string
}
