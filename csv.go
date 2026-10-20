package api

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteCSV() {

	file, err := os.Create("buses.csv")
	if err != nil {
		fmt.Println("Error creating CSV:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{"busid", "busname", "from", "to", "rupees"},
		{"1", "Chennai Express", "Chennai", "Bangalore", "500"},
		{"2", "Coimbatore Express", "Chennai", "Coimbatore", "400"},
	}

	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing CSV:", err)
			return
		}
	}

	fmt.Println("CSV file written successfully")
}