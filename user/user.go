package user

type User struct{
	name string
	email string
}
//constructor
func UserConstructor(name,email string) User{
	return User{name:name,email:email}
}
//getName
func (user User) GetName() string{
	return user.name
}
//getEmail
func (user User) GetEmail() string{
	return user.email
}
//setName
func (user *User) SetName(name string){
	user.name=name
}
//setEmail
func (user *User) SetEmail(email string){
	user.email=email
}

