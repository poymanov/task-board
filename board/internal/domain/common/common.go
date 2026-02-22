package common

type Board struct {
	Id int

	Name string

	Description string

	OwnerId int

	Columns []Column
}

type Column struct {
	Id int

	Name string

	Position float64

	Tasks []Task
}

type Task struct {
	Id int

	Title string

	Description string

	Assignee string

	Position float64
}
