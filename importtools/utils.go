package main

import (
	"errors"
	"os"
	"strings"
)

func equalsStringSlices(a []string, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func validateArgs(args []string) (string, error) {
	const nbArgs = 2
	if len(args) < nbArgs {
		return "", errors.New("Missing arguments.")
	}
	if len(args) > nbArgs {
		return "", errors.New("Too many arguments.")
	}
	arg := args[1]
	_, err := os.Stat(arg)
	if err != nil {
		return "", err
	}
	if arg == strings.TrimSuffix(arg, ".csv") {
		return "", errors.New("Invalid file extension. CSV expected.")
	}
	return arg, nil
}

func validateCsvData(header string, referenceHeaders []string) bool {
	cleanedHeader := strings.TrimSuffix(header, ",")
	headers := strings.Split(cleanedHeader, ",")

	return equalsStringSlices(referenceHeaders, headers)
}
