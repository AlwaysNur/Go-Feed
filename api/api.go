package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var get string = ""

func HandlePost(w http.ResponseWriter, r *http.Request) {
	// Check for the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	// Extract the token from the Authorization header (Bearer <token>)
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}

	token := parts[1]
	// Validate the token (for example, check it against a known valid token)
	if token != "your-token" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		fmt.Println("** Invalid token **")
		return
	}

	// Token is valid, proceed with processing the request...
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	get += "<div class='m'>" + string(body) + "</div><br>"
	// received data
	fmt.Fprintf(w, "Received POST request with body: %s\n", body)

	// Convert ASCII codes to string and clean up
	var asciiString string
	for _, code := range body {
		asciiString += string(rune(code))
	}

	// Remove newline and spaces
	asciiString = strings.Replace(asciiString, "\n", "", -1)
	asciiString = strings.Replace(asciiString, " ", "", -1)

	// Output the cleaned up ASCII string
	// fmt.Println(asciiString)
}

func Response(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<head><meta charset=UTF-8><meta content='width=device-width,initial-scale=1'name=viewport><title>Feed</title><link href=data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48IS0tIFVwbG9hZGVkIHRvOiBTVkcgUmVwbywgd3d3LnN2Z3JlcG8uY29tLCBHZW5lcmF0b3I6IFNWRyBSZXBvIE1peGVyIFRvb2xzIC0tPgo8c3ZnIHdpZHRoPSI4MDBweCIgaGVpZ2h0PSI4MDBweCIgdmlld0JveD0iMCAwIDI0IDI0IiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGlkPSJjaGF0LWFsdC0yIiBjbGFzcz0iaWNvbiBnbHlwaCI+PHBhdGggZD0iTTIwLDJINEEyLDIsMCwwLDAsMiw0VjE2YTIsMiwwLDAsMCwyLDJIN3YzYTEsMSwwLDAsMCwuNTcuOUEuOTEuOTEsMCwwLDAsOCwyMmExLDEsMCwwLDAsLjYyLS4yMkwxMy4zNSwxOEgyMGEyLDIsMCwwLDAsMi0yVjRBMiwyLDAsMCwwLDIwLDJaIiBzdHlsZT0iZmlsbDojMjMxZjIwIj48L3BhdGg+PC9zdmc+ rel='shortcut icon'><link rel='stylesheet' href='/css'></head><body>%s</body>", get)
}
