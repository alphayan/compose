package main

func create(u *User) {
	engine.InsertOne(u)
}
func creates(us []User) {
	engine.Insert(us)
}
func update(u *User) {
	engine.Update(u)
}
func read(u *User) {
	engine.Get(u)
}
func reads(u *User) {
	engine.Find(u)
}
func delete(u *User) {
	engine.Delete(u)
}
