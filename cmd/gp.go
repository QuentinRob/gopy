package cmd

import (
	"errors"
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gopy/internal/file"
	"os"
	"strconv"
)

func Copy(_ *cobra.Command, args []string) error {
	p, _ := pterm.DefaultProgressbar.WithTotal(len(args) - 1).WithTitle("Copy files...").Start()
	var files []file.FileCopy
	for _, arg := range args[:len(args)-1] {
		fileToCopy, err := file.NewFile(arg, args[len(args)-1])
		if err != nil {
			return err
		}
		files = append(files, fileToCopy)
	}
	for _, fileToCopy := range files {
		err := fileToCopy.DoCopy(func(source string, destination string) {
			pterm.Success.Printf("Copy:  %s -> %s\n", source, destination)
			p.Increment()
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func Validate(cmd *cobra.Command, args []string) error {
	fileInfo, err := os.Stat(args[0])
	if err != nil {
		return err
	}

	isRecursive, err := strconv.ParseBool(cmd.Flag(RECURSIVE_FLAG).Value.String())
	if err != nil {
		return err
	}

	if fileInfo.IsDir() && !isRecursive {
		return errors.New(fmt.Sprintf("-r not specified; omitting directory '%s'", args[0]))
	}

	return nil
}
