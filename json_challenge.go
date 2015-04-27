package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Name      string `json:"name"`
	Credit    string `json:"creditcard"`
	City      string `json:"city"`
	Email     string `json:"email"`
	Mac       string `json:"mac"`
	Timestamp string `json:"timestamp"`
}

func GetJson(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func BuildData(data []byte) []Data {
	var d []Data
	err := json.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}

	return d
}

func FilterData(data []Data) []Data {
	result := []Data{}

	for _, line := range data {
		if line.Credit != "" {
			result = append(result, line)
		}
	}

	return result
}

func ConvertToPairs(data []Data) [][]string {
	data = FilterData(data)
	result := make([][]string, len(data))

	for i, line := range data {
		result[i] = []string{line.Name, line.Credit}
	}

	return result
}

func GenerateFilename() string {
	now := time.Now()
	return fmt.Sprintf("%d%s%d.csv", now.Year(), now.Month(), now.Day())
}

func BuildCSV(data []Data) error {
	filename := GenerateFilename()

	csvfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer csvfile.Close()

	records := ConvertToPairs(data)

	writer := csv.NewWriter(csvfile)

	for _, record := range records {
		err = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

	return err
}

func main() {
	bodyJson := GetJson("https://gist.githubusercontent.com/jorin-vogel/7f19ce95a9a842956358/raw/e319340c2f6691f9cc8d8cc57ed532b5093e3619/data.json")

	d := BuildData(bodyJson)

	if err := BuildCSV(d); err != nil {
		panic(err)
	}

	fmt.Printf("Data: %v\n", d[0:5])
}
