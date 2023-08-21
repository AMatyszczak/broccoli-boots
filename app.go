package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/dgraph-io/badger/v3"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	runtime *wails.Runtime
}

type Filter struct {
	operationType string
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

func (a *App) FetchOperationsForFilter(filter Filter) ([]map[string]string, error) {
	operations, err := a.readValueByKey("operations")
	if err != nil {
		return nil, fmt.Errorf("Unable to read operations from db")
	}

	return operations, nil
}

func (a *App) saveKeyValue(key string, value []map[string]string) error {
	dataBytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON")
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(dataBytes))
		return err
	})

	if err != nil {
		return fmt.Errorf("error updating database")
	}
	return nil
}

func (a *App) readValueByKey(key string) ([]map[string]string, error) {

	var m []map[string]string
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {

			if err := json.Unmarshal(val, &m); err != nil {
				return fmt.Errorf("Error Unmarshaling json, %v", err)
			}

			return nil
		})
		return err
	})
	return m, err
}

func (a *App) SelectFile() error {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})

	CSVFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
	}
	defer CSVFile.Close()

	reader := csv.NewReader(CSVFile)

	var jsonData []map[string]string
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV headers")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("error reading CSV record")
		}

		data := make(map[string]string)
		for i, value := range record {
			data[headers[i]] = value
		}

		jsonData = append(jsonData, data)
	}

	a.saveKeyValue("operations", jsonData)
	runtime.EventsEmit(a.ctx, "operations-loaded", nil)

	return nil
}
