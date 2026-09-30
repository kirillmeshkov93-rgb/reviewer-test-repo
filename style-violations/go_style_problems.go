package styleviolations

import "fmt"

func ProcessUser(UserName string, user_age int) {
	user_name := UserName
	fmt.Println("User: " + user_name + ", age: " + fmt.Sprint(user_age))
}
