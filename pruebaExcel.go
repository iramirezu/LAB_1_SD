package main
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)
func main() {
	rows := readSample()
	for i := 1; i < len(rows); i++{
		
		fmt.Printf("Row: %s \n", rows[i][0])

	}
}

func readSample() [][]string {
    f, err := os.Open("registroCamion.csv")
    if err != nil {
        log.Fatal(err)
    }
    rows, err := csv.NewReader(f).ReadAll()
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
    return rows
}


func writeChanges(rows [][]string) {
    f, err := os.Create("registroCamion.csv")
    if err != nil {
        log.Fatal(err)
    }
    err = csv.NewWriter(f).WriteAll(rows)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }
}