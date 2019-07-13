//bcrypt example
import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword1 := `password1234`// change to password123 to have this run right.
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOG IN, PASSWORD MALFUNTION")
		return
	}
	fmt.Println("You're loggin in.")