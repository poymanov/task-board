package task

type NewTask struct {
	Title string

	Description string

	Assignee string

	ColumnId int
}

func NewNewTask(title, description, assignee string, columnId int) NewTask {
	return NewTask{
		Title:       title,
		Description: description,
		Assignee:    assignee,
		ColumnId:    columnId,
	}
}

type GetAllFilter struct {
	ColumnIds []int
}

type GetAllSort struct {
	SortByColumnId string

	SortByPosition string
}

func NewGetAllFilter(columnIds []int) GetAllFilter {
	return GetAllFilter{
		ColumnIds: columnIds,
	}
}

func NewGetAllSort(sortByColumnId, sortByPosition string) GetAllSort {
	return GetAllSort{
		SortByColumnId: sortByColumnId,
		SortByPosition: sortByPosition,
	}
}
