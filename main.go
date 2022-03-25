package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type TimeSeriesData struct {
	Id    uint64 `json:"id"`
	Value uint64 `json:"value"`
}

func main() {

	sh := shell.NewShell("localhost:5001")

	tsd := &TimeSeriesData{
		Id:    1,
		Value: 23}

	tsdBin, _ := json.Marshal(tsd)
	reader := bytes.NewReader(tsdBin)

	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %s", err)
		os.Exit(1)
	}
	fmt.Printf("Added %s\n", cid)

	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%s", err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()

	res := &TimeSeriesData{}
	json.Unmarshal([]byte(newStr), &res)
	fmt.Println(res)
}
