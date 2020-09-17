package controllers

type Index struct {
	Base
}

func (CON *Index) GetIndex(){
	data := map[string]interface{}{
		"data":"hello world",
	}
	CON.SuccReturn(data)
}
