package strunctInit

type EsQuery struct {
	Query string `json:"query"`
}

type EsData struct {
	Columns []Columns       `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}
type Columns struct {
	Name string `json:"name"`
}
