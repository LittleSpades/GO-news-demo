package transport

//Search is ...
type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    interface{}
}

//IsLastPage is ...
func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

//CurrentPage is ...
func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

//PreviousPage is ...
func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}
