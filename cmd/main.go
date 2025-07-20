package main

import (
	"log"
	"raytracer/internal"

	"github.com/spf13/cobra"
)

func main() {
	var outputFile string

	var rootCmd = &cobra.Command{
		Use:   "raytracer",
		Short: "A 3D raytracer that renders scenes to PNG images",
		Long:  "A Go-based raytracer that renders 3D scenes using ray tracing algorithms and outputs PNG images.",
		Run: func(cmd *cobra.Command, args []string) {
			err := internal.Render(outputFile)
			if err != nil {
				log.Fatal(err.Error())
			}
		},
	}

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output filename for the rendered image")
	rootCmd.MarkFlagRequired("output")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
