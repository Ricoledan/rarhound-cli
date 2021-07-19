/*
Copyright © 2021 Ricardo Ledan <ricoledan@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func unzip(src string, dest string) ([]string, error) {
	fmt.Println("🐕 sniff...sniff...Woof! ", src)

	var filenames []string
	r, err := zip.OpenReader(src)

	if err != nil {
		log.Fatal(err)
		return filenames, err
	}

	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)
		println("filename: ", filenames)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Woof! Woof! 🐕",
	Long:  `Unpacks RAR files recursively within a specified directory so you can concentrate on enjoying your content.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		unzip(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
