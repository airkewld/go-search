package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	fmt.Println("Building with Dagger")
    cleanBuildDir()
	if err := build(context.Background()); err != nil {
		fmt.Println("Error with pipeline: ", err)
	} else {
		fmt.Println("App built successfully!")
	}
}

func cleanBuildDir() {
	fmt.Println("Cleaning up 'build' directory...")
	os.RemoveAll("./build")
}

func build(ctx context.Context) error {

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	buildFor := []string{"linux", "darwin", "windows"}

	// get reference to the local project
	src := client.Host().Directory(".")

	// set images
	builderImg := "golang:1.20"
	goSecImg := "securego/gosec"

	// set builder and goSec containers
	goSec := client.Container().From(goSecImg)
	builder := client.Container().From(builderImg)

	// mount cloned app path into `builder` image
	builder = builder.WithDirectory("/src", src).WithWorkdir("/src")
	// mount cloned app path into `goSec` image
	goSec = goSec.WithDirectory("/src", src).WithWorkdir("/src")

	// run simple fmt
	gofmt, err := builder.WithExec([]string{"go", "fmt", "."}).Stdout(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gofmt)
	}
	// run gosec
	gosec, err := goSec.WithExec([]string{"gosec", "."}).Stdout(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gosec)
	}

	// define the application build command
	for _, os := range buildFor {
		path := fmt.Sprintf("build/%s/", os)
		builder = builder.WithEnvVariable("GOOS", os)
		builder = builder.WithExec([]string{"go", "build", "-o", path})
		// get reference to build output directory in container
		output := builder.Directory(path)
		// write contents of container build/ directory to the host
		_, err = output.Export(ctx, path)
		if err != nil {
			return err
		}
	}
	return nil
}
