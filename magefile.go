// +build mage

package main

import (
	"fmt"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const name string = "stingshell"
const buildDir string = "build"

type (
	Build   mg.Namespace
	Test    mg.Namespace
	CI      mg.Namespace
	Install mg.Namespace
)

type BuildOptions struct {
	Static    bool
	Extension string
}

func build(os string, arch string, entry string, opts BuildOptions) error {
	env := map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}

	// Build static binary
	if opts.Static {
		env["CGO_ENABLED"] = "0"
	}

	outputName := fmt.Sprintf(
		"%s-%s-%s", name, os, arch,
	)

	outputPath := path.Join(
		buildDir,
		outputName,
	)

	if opts.Extension != "" {
		outputPath += fmt.Sprintf(".%s", opts.Extension)
	}

	args := []string{
		"build",
		"-o",
		outputPath,
		"-ldflags",
		"-s -w",
		entry,
	}

	fmt.Println(args)

	return sh.RunWith(
		env, "go", args...,
	)
}

/*
Build
*/

// Builds all binaries
func (Build) All() {
	mg.Deps(
		Build.DarwinAmd64,
		Build.LinuxAmd64,
	)
}

// Builds the shell for Linux (amd64)
func (Build) LinuxAmd64() error {
	return build("linux", "amd64", ".", BuildOptions{})
}

// Builds the shell for Darwin/macOS (amd64)
func (Build) DarwinAmd64() error {
	return build("darwin", "amd64", ".", BuildOptions{})
}

/*
Test
*/

// Runs all unit tests
// func (Test) Unit() error {
// 	return sh.RunV("go", "test", "-v", "./...")
// }

/*
CI/CD
*/

// Compresses all binaries into a single tarball
func (CI) CompressBinaries() error {
	return sh.Run("tar", "-cvzf", "binaries.tar.gz", "build")
}

/*
Misc
*/

// Cleans up build directories
func Clean() {
	sh.Rm("build")
}

// Installs the server to /usr/local/bin
// (Linux-only)
func (Install) Server() error {
	return sh.Run(
		"cp",
		"build/stingshell-linux-amd64",
		"/usr/local/bin/stingshell",
	)
}
