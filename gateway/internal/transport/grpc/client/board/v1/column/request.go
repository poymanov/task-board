package column

type CreateColumnRequest struct {
	Name string

	BoardId int
}

type DeleteColumnRequest struct {
	Id int
}

type UpdatePositionColumnRequest struct {
	Id int

	LeftPosition float64

	RightPosition float64
}
