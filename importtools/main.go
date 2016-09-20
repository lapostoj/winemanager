package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/lapostoj/winemanager/service/presentation/api/request"
)

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
	"Format"}

var url = "http://localhost:8080/api/wines"

func importDataLine(line string) {
	cleanedLine := strings.TrimSuffix(line, ",")
	data := strings.Split(cleanedLine, ",")
	postWineRequest := request.PostWineRequest{
		Name: data[0],
	}
	request, err := json.Marshal(postWineRequest)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(request))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "dev_appserver_login=\"test@example.com:True:185804764220139124118\"")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Posting: %s\n", err)
	} else {
		if res.StatusCode != http.StatusCreated {
			fmt.Printf("Posting: %s\n", res.Status)
		}
	}
}

func readLineByLine(f *os.File) (int, error) {
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	lines := 0
	for scanner.Scan() {
		if lines == 0 {
			if !validateCsvData(scanner.Text(), referenceHeaders) {
				return 0, errors.New("Invalid data.")
			}
		} else {
			importDataLine(scanner.Text())
		}
		lines++
	}
	return lines - 1, nil
}

// This would work only if there is no restriction in the API as it use the POST request.
func main() {
	args := os.Args
	filePath, err := validateArgs(args)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("Reading: %s\n", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	defer f.Close()

	lines, err := readLineByLine(f)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("%d data lines read\n", lines)
}
