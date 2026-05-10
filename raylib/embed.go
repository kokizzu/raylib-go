//go:build !cgo && !raylib_no_embed && (windows || linux || darwin) && (amd64 || arm64)
// +build !cgo
// +build !raylib_no_embed
// +build windows linux darwin
// +build amd64 arm64

package rl

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	_ "embed"
	"io"
	"os"
	"path/filepath"
)

//go:embed libs/LICENSE
var embeddedLicense []byte

// extractLib extracts the embedded shared library and returns the path to it
func extractLib() (libname string) {
	if os.Getenv("RAYLIB_NO_EMBED") == "1" {
		return
	}

	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	// becomes "C:\Users\{username}\AppData\Local\github.com\gen2brain\raylib-go\"
	// or "$HOME/.cache/github.com/gen2brain/raylib-go/"
	outDir := filepath.Join(userCacheDir, "github.com", "gen2brain", "raylib-go")

	// ensure directory exists
	if err := os.MkdirAll(outDir, 0755); err != nil {
		panic(err)
	}

	// write license file, if not already exists
	licenseFile := filepath.Join(outDir, "LICENSE")
	if _, err := os.Stat(licenseFile); err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(licenseFile, embeddedLicense, 0644); err != nil {
				panic(err)
			}
		}
	}

	gz, err := gzip.NewReader(bytes.NewReader(embeddedLib))
	if err != nil {
		panic(err)
	}
	defer gz.Close()

	tr := tar.NewReader(gz)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// only one file expected
		if hdr.Typeflag == tar.TypeReg {
			destLib := filepath.Join(outDir, hdr.Name)

			// write library only if file not already exists
			if fileInfo, err := os.Stat(destLib); err != nil {
				if os.IsNotExist(err) {
					if file, err := os.OpenFile(destLib, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
						defer file.Close()
						if _, err := io.Copy(file, tr); err != nil {
							panic(err)
						}
						libname = destLib
					} else {
						panic(err)
					}
				}
			} else {
				if fileInfo != nil && !fileInfo.IsDir() {
					libname = destLib
				}
			}
			break
		}
	}
	return
}
