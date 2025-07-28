package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("file path is required")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	// data -> Encode(data)
	packed := "" + string(data) // TODO: remove
	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}
	// 0644 - права для файла
	// 6 - текущий пользователь может читать и писать
	// 4 - группа может только читать
	// 4 - все остальные пользователи могут только читать
}

func packedFileName(path string) string {
	// path/to/file.txt -> file.vlc

	fileName := filepath.Base(path) // file.vlc
	// ext := filepath.Ext(fileName)					// .txt
	// baseName := strings.TrimSuffix(fileName, ext)	// file

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
