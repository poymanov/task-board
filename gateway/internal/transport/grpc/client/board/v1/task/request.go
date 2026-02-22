package task

type CreateTaskRequest struct {
	Title string

	Description string

	Assignee string

	ColumnId int
}

type DeleteTaskRequest struct {
	Id int
}

type UpdatePositionTaskRequest struct {
	Id int

	LeftPosition float64

	RightPosition float64
}
