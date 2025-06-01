package paging

type Direction int

const (
	ASC Direction = iota
	DESC
)

type SortedField struct {
	Field     string
	Direction Direction
}

type PagedSort struct {
	sorted, paged     bool
	sortBy            []SortedField
	pageSize, pageNum uint
}

func Unpaged() PagedSort {
	return PagedSort{}
}

func (p PagedSort) Paged(pageSize, pageNum uint) PagedSort {
	return PagedSort{
		paged:    true,
		pageSize: pageSize,
		pageNum:  pageNum,
		sorted:   p.sorted,
		sortBy:   p.sortBy,
	}
}

func (p PagedSort) SortBy(field string, dir Direction) PagedSort {
	return PagedSort{
		paged:    p.paged,
		pageSize: p.pageSize,
		pageNum:  p.pageNum,
		sorted:   true,
		sortBy:   append(p.sortBy, SortedField{field, dir}),
	}
}
