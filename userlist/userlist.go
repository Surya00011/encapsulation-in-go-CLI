package userlist
import("encapsulation/user")

var users []user.User

func AddUser(newUser user.User){
	users = append(users,newUser)
}
func GetAllUsers() []user.User{
	return users
}
func DeleteUser(email string) string{

	for index,user := range users{
		if user.GetEmail() == email {
	         users = append(users[:index],users[index+1:] ...)		
                 return "User deleted"
		}
        }
	return "User Notfound"
}

