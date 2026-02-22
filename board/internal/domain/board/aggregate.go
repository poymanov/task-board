package board

type NewBoard struct {
	Name string

	Description string

	OwnerId int
}

type GetAllFilter struct {
	OwnerId int
}

func NewNewBoard(name, description string, ownerId int) NewBoard {
	return NewBoard{
		Name:        name,
		Description: description,
		OwnerId:     ownerId,
	}
}

func NewGetAllFilter(ownerId int) GetAllFilter {
	return GetAllFilter{
		OwnerId: ownerId,
	}
}
