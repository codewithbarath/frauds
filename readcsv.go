package api

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func FetchCSV(writer http.ResponseWriter, request *http.Request) {

	log.Println("FetchCSV(+)")

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusNoContent)
		log.Println("FetchCSV(-)")
		return
	}

	if request.Method != http.MethodGet {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
		log.Println("FetchCSV(-)")
		return
	}

	// Open the CSV file
	lFile, lErr := os.Open("buses.csv")

	if lErr != nil {
		log.Println("Error opening CSV:", lErr)
		http.Error(writer, "Unable to open CSV file", http.StatusInternalServerError)
		log.Println("FetchCSV(-)")
		return
	}

	defer lFile.Close()

	// Read the CSV file
	lReader := csv.NewReader(lFile)

	lRecords, lErr := lReader.ReadAll()

	if lErr != nil {
		log.Println("Error reading CSV:", lErr)
		http.Error(writer, "Unable to read CSV file", http.StatusInternalServerError)
		log.Println("FetchCSV(-)")
		return
	}

	// Convert CSV records into JSON
	lErr = json.NewEncoder(writer).Encode(lRecords)

	if lErr != nil {
		log.Println("Error sending CSV data:", lErr)
		log.Println("FetchCSV(-)")
		return
	}

	log.Println("FetchCSV(-)")
}