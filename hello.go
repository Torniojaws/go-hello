package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("Hello there\n")

    // Read a CSV file
    path, err := os.Getwd()
    if err != nil {
        fmt.Println("Something exploded! ", err)
    }
    filename := "test.csv"
    separator := ","
    f := CSVFile{path, filename, separator}
    fmt.Printf("Current file is %s\n", f.GetAbsolutePath())

    // Get the row count from the file
    rowCount, err := f.GetRowCount()
    if err != nil {
        fmt.Printf("Could not read file! ", err)
    }
    fmt.Printf("The file has %d rows\n", rowCount)

    // Print the data from the 3rd row of the file
    row, err := f.ReadRow(3)
    if err != nil {
        fmt.Printf("Couldn't read the row - does the file have so many rows?", err)
    }
    fmt.Printf("")
}
