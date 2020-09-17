package es

import (
	"bytes"
	"context"
	"encoding/json"
	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	jsoniter "github.com/json-iterator/go"
	"beegoTool/strunctInit"
	"strings"
)

func Select(sql string) []map[string]interface{} {
	returnDatas := []map[string]interface{}{}
	sli := strings.Split(sql, " ")
	for _, v := range sli {
		v = strings.ToUpper(v)
		if v == `UPDATE` || v == "DELETE" {
			return returnDatas
		}
	}
	esConfig := elasticsearch6.Config{
		Username:  "xxx",
		Password:  "xxxx",
		Addresses: []string{"xxxxxx:9200"},
	}
	c, err := elasticsearch6.NewClient(esConfig)
	if err != nil {
		return returnDatas
	}
	query := map[string]interface{}{
		"query": sql,
	}
	jsonBody, _ := json.Marshal(query)
	req := esapi.XPackSQLQueryRequest{
		Body: bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), c)
	if err != nil || res.IsError() {
		return returnDatas
	}
	str := res.String()
	jsonStr := str[9:len(str)]
	data := strunctInit.EsData{}
	err = jsoniter.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return returnDatas
	}

	for _, v := range data.Rows {
		datas := map[string]interface{}{}
		i := 0
		for _, key := range data.Columns {
			datas[key.Name] = v[i]
			i = i + 1
		}
		returnDatas = append(returnDatas, datas)
	}
	return returnDatas
}
