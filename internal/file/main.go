package file

import (
	"log"
	"os"
	"path"
)

type FileCopy struct {
	SourcePath      string
	DestinationPath string
	Mode            os.FileMode
}

func NewFile(source string, destination string) (FileCopy, error) {

	sourceFileInfo, err := os.Stat(source)
	if err != nil {
		return FileCopy{}, err
	}

	var destinationFilePath string
	destinationFileInfo, err := os.Stat(destination)
	if err != nil {
		if !os.IsNotExist(err) {
			return FileCopy{}, err
		} else {
			destinationFilePath = destination
		}
	} else {
		if destinationFileInfo.IsDir() {
			destinationFilePath = path.Join(destination, sourceFileInfo.Name())
		}
	}

	return FileCopy{
		SourcePath:      source,
		DestinationPath: destinationFilePath,
		Mode:            sourceFileInfo.Mode(),
	}, nil
}

func (fc *FileCopy) DoCopy(onSuccess func(source string, destination string)) error {
	bytesRead, err := os.ReadFile(fc.SourcePath)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fc.DestinationPath, bytesRead, fc.Mode)

	if err != nil {
		log.Fatal(err)
	}

	onSuccess(fc.SourcePath, fc.DestinationPath)

	return nil
}
