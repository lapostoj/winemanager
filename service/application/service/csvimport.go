package service

import (
	"bufio"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"
	"google.golang.org/appengine/log"
)

// CsvImport implements a service to parse a CSV file and import in in the db.
type CsvImport struct {
}

// ExecuteCsvImport executes the service.
func ExecuteCsvImport(ctx context.Context, reader *bufio.Reader) ([]wine.Wine, error) {
	wines, err := readLineByLine(ctx, reader)
	if err != nil {
		return *new([]wine.Wine), errors.New("ExecuteCsvImport - " + err.Error())
	}
	log.Infof(ctx, "ExecuteCsvImport - %d data lines read", len(wines))
	return wines, nil
}

func readLineByLine(ctx context.Context, reader *bufio.Reader) ([]wine.Wine, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	linesRead := 0
	var wines []wine.Wine
	for scanner.Scan() {
		line := scanner.Text()
		if linesRead == 0 {
			if !validateData(line, referenceHeaders) {
				return wines, errors.New("Invalid data")
			}
		} else {
			wine := persistDataLine(ctx, line)
			wines = append(wines, *wine)
		}
		linesRead++
	}
	return wines, nil
}

func validateData(header string, referenceHeaders []string) bool {
	headers := parseLine(header)
	return EqualsStringSlices(referenceHeaders, headers)
}

func persistDataLine(ctx context.Context, line string) *wine.Wine {
	data := parseLine(line)
	wine := dataToWine(data)
	persistence.SaveWine(ctx, wine)
	return wine
}

func parseLine(line string) []string {
	cleanedLine := strings.TrimSuffix(line, ",")
	return strings.Split(cleanedLine, ",")
}

func dataToWine(data []string) *wine.Wine {
	return &wine.Wine{
		Name:        data[0],
		Designation: data[1],
		Growth:      data[2],
		Year:        StringToInt(data[3]),
		Country:     "France",
		Region:      data[4],
		Color:       wine.StringToColor(data[5]),
		Type:        wine.StringToType(data[6]),
		Quantity:    StringToInt(data[7]),
		Producer:    data[8],
		Size:        wine.IntToSize(StringToInt(data[10])),
		StorageLocation: wine.StorageLocation{
			Cellar: "Moiré",
		},
		CreationTime: time.Now().UTC(),
	}
}

var referenceHeaders = []string{
	"Nom",
	"Appellation",
	"Cru",
	"Millésime",
	"Région",
	"Couleur",
	"Type",
	"Stock",
	"Producteur",
	"Origine",
	"Format",
}
