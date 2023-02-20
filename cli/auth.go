package cli

import "fmt"

var UserAuth *User

func SignIn() error {
	fmt.Println("Please insert your email")
	Scanner.Scan()
	fmt.Println("Please insert your password")
	Scanner.Scan()
	if true {
		UserAuth = &User{
			Name:     "Aliz",
			Email:    "ar.mokhtaro.g@gmail.com",
			Password: "1212",
		}
		return nil
	}
	UserAuth = &User{}
	return fmt.Errorf("sign-in failed")
}

func SignUp() error {
	if true {
		return nil
	}
	return fmt.Errorf("registration failed")
}
