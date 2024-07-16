package mysql

type ContentModel struct {
	BaseModel
	Title string
	Rate  int
	Text  string
}

func (c ContentModel) TableName() string {
	return "contents"
}
