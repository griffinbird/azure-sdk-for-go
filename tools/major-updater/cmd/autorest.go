package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var autorestCmd = &cobra.Command{
	Use:   "autorest <SDK dir> <specs dir>",
	Short: "Execute autorest on specs, saving generated SDK code into SDK dir",
	Long: `This command will execute autorest on the specs dir, 
	saving the generated SDK code into SDK dir, then runs some after-scripts`,
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.ExactArgs(2)(cmd, args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		sdk := args[0]
		spec := args[1]
		err := theAutorestCommand(sdk, spec)
		return err
	},
}

func init() {
	rootCmd.AddCommand(autorestCmd)
}

func theAutorestCommand(sdk, spec string) error {
	printf("Executing autorest (%d threads)\n", thread)
	err := os.Setenv("NODE_OPTIONS", "--max-old-space-size=8192")
	if err != nil {
		return fmt.Errorf("failed to set environment variable: %v", err)
	}
	// get absolute path
	absolutePathOfSDK, err := filepath.Abs(sdk)
	if err != nil {
		return fmt.Errorf("failed to get the directory of SDK: %v", err)
	}
	absolutePathOfSpec, err := filepath.Abs(spec)
	if err != nil {
		return fmt.Errorf("failed to get the directory of specs: %v", err)
	}
	// get every single readme.md file in the directory
	files, err := selectFilesWithName(absolutePathOfSpec, readme)
	vprintf("Found %d readme.md files\n", len(files))
	jobs := make(chan work, 1000)
	results := make(chan error, 1000)
	for i := 0; i < thread; i++ {
		go worker(i, jobs, results)
	}
	for _, file := range files {
		jobs <- work{filename: file, sdkFolder: absolutePathOfSDK}
	}
	close(jobs)
	for range files {
		<-results
	}
	vprintln("autorest finished")
	return nil
}
