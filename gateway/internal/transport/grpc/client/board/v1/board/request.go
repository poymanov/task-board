package board

type CreateBoardRequest struct {
	Name string

	Description string

	OwnerId int
}

type GetBoardRequest struct {
	Id int
}
