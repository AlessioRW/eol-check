package output

import (
	"bytes"
	"eol-checker/internal/api"
	"fmt"
	"log/slog"
	"os"
)

type Client struct {
}

func Output(outputClient Client, checks []api.EolCheck) error {
	outBuffer := bytes.Buffer{}
	headLine := "id,product,version,is_eol,eol_date,maintained,lts\n"
	_, err := outBuffer.Write([]byte(headLine))
	if err != nil {
		slog.Error("error writing csv heading line", "error", err)
		return err
	}

	for _, check := range checks {
		fmt.Printf("%+v\n", check)
		checkLine := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v\n",
			check.CheckId,
			check.Product,
			check.Current.VersionNum,
			check.Current.IsEol,
			check.Current.EolDate,
			check.Current.IsMaintained,
			check.Current.IsLTS,
		)

		_, err = outBuffer.Write([]byte(checkLine))
		if err != nil {
			slog.Error("error writing check line into csv", "error", err)
			return err
		}
	}
	return outputClient.Write(outBuffer.Bytes())
}

func (client Client) Write(outputBuffer []byte) error {
	file, err := os.Create("eol-output.csv")
	if err != nil {
		slog.Error("error creating output file", "error", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(outputBuffer)
	if err != nil {
		slog.Error("error writing check line into csv", "error", err)
		return err
	}

	return nil

}
