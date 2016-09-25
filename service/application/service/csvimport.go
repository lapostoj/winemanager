package service

import (
	"bufio"
	"errors"
	"mime/multipart"
	"strings"
	"time"

	"github.com/lapostoj/winemanager/service/domain/model/wine"
	"github.com/lapostoj/winemanager/service/infrastructure/persistence/datastore"

	"google.golang.org/appengine/log"

	"golang.org/x/net/context"
)

// CsvImport implements a service to parse a CSV file and import in in the db.
type CsvImport struct {
}

// ExecuteCsvImport executes the service.
func ExecuteCsvImport(ctx context.Context, file multipart.File) error {
	lines, err := readLineByLine(ctx, file)
	if err != nil {
		return errors.New("ExecuteCsvImport: " + err.Error())
	}
	log.Infof(ctx, "%d data lines read", lines)
	return nil
}

func readLineByLine(ctx context.Context, file multipart.File) (int, error) {
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	lines := 0
	for scanner.Scan() {
		if lines == 0 {
			if !validateData(scanner.Text(), referenceHeaders) {
				return 0, errors.New("Invalid data.")
			}
		} else {
			persistDataLine(ctx, scanner.Text())
		}
		lines++
	}
	return lines - 1, nil
}

func validateData(header string, referenceHeaders []string) bool {
	headers := parseLine(header)
	return equalsStringSlices(referenceHeaders, headers)
}

func persistDataLine(ctx context.Context, line string) {
	data := parseLine(line)
	wine := dataToWine(data)
	persistence.SaveWine(ctx, wine)
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
		Year:        stringToInt(data[3]),
		Country:     "France",
		Region:      data[4],
		Color:       wine.StringToColor(data[5]),
		Type:        wine.StringToType(data[6]),
		Quantity:    stringToInt(data[7]),
		Producer:    data[8],
		Size:        wine.IntToSize(stringToInt(data[10])),
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
