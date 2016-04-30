// Package
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:",omitempty"`
	Password  string `json:"-"` // don't show the password
}

type UserJSON struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:",omitempty"`
	FullName  string
}

func main() {
	// u := User{FirstName: "Mark", LastName: "Bates", Password: "password"}
	u := User{LastName: "Bates"}
	// m, _ := json.Marshal(u)
	// fmt.Println(string(m))

	// Better way to handle json
	// json.NewEncoder(os.Stdout).Encode(u)

	json.NewEncoder(os.Stdout).Encode(u)

	// Or event this
	// n := json.NewEncoder(os.Stdout)
	// n.Encode(u)
	// n.Encode(u)
	//
	// // equivalenc to
	// m, _ := json.Marshal(u)
	// fmt.Fprint(os.Stdout, string(m))

	s := `{"first_name": "Mark", "LastName": "Bates"}`
	u2 := User{}

	//json.NewDecoder(strings.NewReader(s)).Decode(&u2)
	json.Unmarshal([]byte(s), &u2)

	log.Printf("u.FirstName: %s\n", u2.FirstName)
	log.Printf("u.LastName: %s\n", u2.LastName)

	u3 := User{FirstName: "Mark", LastName: "Bates"}
	json.NewEncoder(os.Stdout).Encode(u3)
}

// First way
// func (u User) MarshalJSON() ([]byte, error) {
// 	m := map[string]string{
// 		"first_name": u.FirstName,
// 		"last_name":  u.LastName,
// 		"full_name":  fmt.Sprintf("%s %s", u.FirstName, u.LastName),
// 	}
// 	return json.Marshal(m)
// }

func (u User) MarshalJSON() ([]byte, error) {
	uj := UserJSON{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		FullName:  fmt.Sprintf("%s %s", u.FirstName, u.LastName),
	}
	return json.Marshal(uj)
}
