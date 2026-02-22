package column

type NewColumn struct {
	Name string

	BoardID int
}

func NewNewColumn(name string, boardID int) NewColumn {
	return NewColumn{
		Name:    name,
		BoardID: boardID,
	}
}

type GetAllFilter struct {
	BoardId int
}

type GetAllSort struct {
	SortByPosition string
}

func NewGetAllFilter(boardId int) GetAllFilter {
	return GetAllFilter{
		BoardId: boardId,
	}
}

func NewGetAllSort(sortByPosition string) GetAllSort {
	return GetAllSort{
		SortByPosition: sortByPosition,
	}
}
