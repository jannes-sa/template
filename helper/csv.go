package helper

import (
	"encoding/csv"
	"os"
)

// CsvWriter  ...
type CsvWriter struct {
	// mutex     *sync.Mutex
	csvWriter *csv.Writer
}

// NewCsvWriter ...
func NewCsvWriter(fileName string) (*CsvWriter, error) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(csvFile)
	return &CsvWriter{
		csvWriter: w,
		// mutex: &sync.Mutex{}
	}, nil
}

// Write ...
func (w *CsvWriter) Write(row []string) {
	// w.mutex.Lock()
	errWrite := w.csvWriter.Write(row)
	CheckErr("error write csv", errWrite)
	// w.mutex.Unlock()
}

// Flush ...
func (w *CsvWriter) Flush() {
	// w.mutex.Lock()
	w.csvWriter.Flush()
	// w.mutex.Unlock()
}
