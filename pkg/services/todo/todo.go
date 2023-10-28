package todo

type Todo struct {
	Id   string
	Text string
	Done string
	User User
}

type User struct {
	Id   string
	Name string
}
