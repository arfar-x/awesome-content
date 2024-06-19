package mysql

type ContentModel struct {
	BaseModel
	Name  string
	Title string
	Rate  int
	Text  string
}
