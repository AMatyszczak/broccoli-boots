package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// func (a *App) SelectFile() ([]byte, error) {
func (a *App) SelectFile() ([]map[string]string, error) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})

	CSVFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
	}
	defer CSVFile.Close()

	fmt.Print("_______________________________SelectFile 1, csvData:", CSVFile)
	reader := csv.NewReader(CSVFile)

	var jsonData []map[string]string
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV headers")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("error reading CSV record")
		}

		data := make(map[string]string)
		for i, value := range record {
			data[headers[i]] = value
		}

		jsonData = append(jsonData, data)
	}

	// jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	// if err != nil {
	// 	return nil, fmt.Errorf("error marshaling JSON")
	// }
	// fmt.Print("_______________________________SelectFile 2, jsonData:", jsonData)

	return jsonData, nil
}
