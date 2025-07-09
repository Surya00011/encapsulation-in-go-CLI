package userlist
import("encapsulation/user")

var users []user.User

func AddUser(newUser user.User){
	users = append(users,newUser)
}
func GetAllUsers() []user.User{
	return users
}
