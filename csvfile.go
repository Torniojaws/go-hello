package main

import (
    "bytes"
    "io"
    "os"
)

type CSVFile struct {
    path, name, separator string
}

func (f CSVFile) GetRowCount() (int, error) {
    file, _ := os.Open(f.GetAbsolutePath())
    r := io.Reader(file)
    bufferSize := 32*1024
    buffer := make([]byte, bufferSize)
    count := 0
    lineSep := []byte{'\n'}
    for {
        char, err := r.Read(buffer)
        count += bytes.Count(buffer[:char], lineSep)
        switch {
        case err == io.EOF:
            return count, nil
        case err != nil:
            return count, err
        }
    }
}

func (f CSVFile) GetAbsolutePath() string {
    var buffer bytes.Buffer
    buffer.WriteString(f.path)
    // Since os.Getwd() does not append "/" to the end
    buffer.WriteString("/")
    buffer.WriteString(f.name)
    return buffer.String()
}
