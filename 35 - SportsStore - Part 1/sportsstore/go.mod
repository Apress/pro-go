module sportsstore

go 1.17

require platform v1.0.0

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
)

replace platform v1.0.0 => ../platform
