package service

import (
	cfg "BEAT/config"
	"BEAT/pkg/model"
	"log"
	"os"
)

type resposeWriter struct {
	resultFile *os.File
}

func NewFileResposeWriter(config *cfg.Config) model.Writer {
	resultFile, err := os.Create(config.OutputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	return &resposeWriter{
		resultFile: resultFile,
	}
}

func (r resposeWriter) Close() error {
	err := r.resultFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r resposeWriter) Write(p []byte) (n int, err error) {
	_, err = r.resultFile.WriteString(string(p))
	if err != nil {
		log.Fatal(err)
	}
	return 1, nil

}
