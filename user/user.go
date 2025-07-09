package user

type User struct{
	name string
	email string
}
//constructor
func userConstructor(name,email string) User{
	return User{name:name,email:email}
}
//getName
func (user User) getName() string{
	return user.name
}
//getEmail
func (user User) getEmail() string{
	return user.email
}
//setName
func (user *User) setName(name string){
	user.name=name
}
//setEmail
func (user *User) setEmail(email string){
	user.email=email
}

