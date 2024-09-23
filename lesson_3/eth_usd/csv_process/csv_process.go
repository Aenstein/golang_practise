package csv_process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type EthereumPrice struct {
	Price float64
	Date  time.Time
}

func LoadDateFromCSV(path string) ([]EthereumPrice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error openingprices file: %w", err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	var pricePair []EthereumPrice

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading from csv file: %w", err)
		}

		parsedPrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse price: %w", err)
		}

		parsedTime, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, fmt.Errorf("cannot parse time: %w", err)
		}
		
		pricePair = append(pricePair, EthereumPrice{
			Price: parsedPrice,
			Date: parsedTime,
		})
	}

	return pricePair, nil
}
