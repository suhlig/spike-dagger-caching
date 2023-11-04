package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func mainE() error {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))

	if err != nil {
		return err
	}

	defer client.Close()

	filesStream, err := client.Container().
		From("suhlig/b2").
		WithEnvVariable("B2_APPLICATION_KEY_ID", os.Getenv("B2_APPLICATION_KEY_ID")).
		WithSecretVariable("B2_APPLICATION_KEY", client.SetSecret("B2_APPLICATION_KEY", os.Getenv("B2_APPLICATION_KEY"))).
		WithExec([]string{"b2", "ls", "suhlig-transcription-test"}).
		Stdout(ctx)

	if err != nil {
		return err
	}

	fmt.Println(filesStream)

	return nil
}
