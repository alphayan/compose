package main

func create(u *User) {
	db.Create(u)
}
func update(u *User) {
	db.Updates(u)
}
func read(u *User) {
	db.First(u)
}
func reads() []User {
	us := []User{}
	db.Find(&us)
	return us
}
func delete(u *User) {
	db.Delete(u)
}
