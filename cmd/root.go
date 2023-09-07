package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var silent bool
var countLoops int

var rootCmd = &cobra.Command{
	Use:   "smash [target]",
	Short: "Smash securely deletes files or directories",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return cmd.Help()
		}

		target := args[0]
		absPath, err := filepath.Abs(target) // Get the absolute path
		if err != nil {
			return err
		}

		info, err := os.Stat(target)
		if err != nil {
			return err
		}

		if !silent {
			if info.IsDir() {
				fmt.Printf("Are you sure you want to permanently delete the folder and all its contents at '%s'? [y/N]: ", absPath)
			} else {
				fmt.Printf("Are you sure you want to permanently delete the file at '%s'? [y/N]: ", absPath)
			}

			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			if strings.TrimSpace(strings.ToLower(response)) != "y" {
				fmt.Println("Operation canceled.")
				return nil
			}
		}

		if info.IsDir() {
			return smashDir(target)
		}

		return smashFile(target)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&silent, "silent", "s", false, "Run silently without confirmation")
	rootCmd.Flags().IntVarP(&countLoops, "count", "c", 5, "Number of overwriting loops")
}

func Execute() error {
	return rootCmd.Execute()
}

func smashFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	size := fileInfo.Size()
	buffer := make([]byte, size)

	for i := 0; i < countLoops; i++ {
		for _, pattern := range []byte{0x00, 0xFF, 0xAA} {
			for j := int64(0); j < size; j++ {
				buffer[j] = pattern
			}

			_, err = file.Seek(0, 0)
			if err != nil {
				return err
			}

			_, err = file.Write(buffer)
			if err != nil {
				return err
			}
		}
	}

	if err := file.Close(); err != nil {
		return err
	}

	return os.Remove(filePath)
}

func smashDir(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if file.IsDir() {
			if err := smashDir(filePath); err != nil {
				return err
			}
		} else {
			if err := smashFile(filePath); err != nil {
				return err
			}
		}
	}

	return os.RemoveAll(dirPath)
}
