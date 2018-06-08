package main

import (
	"github.com/spf13/cobra"
	"github.com/solo-io/qloo/pkg/api/types/v1"
)

var schemaGetCmd = &cobra.Command{
	Use:   "get [NAME]",
	Short: "return a schema by its name or list all schemas",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || args[0] == "" {
			list, err := listSchemas()
			if err != nil {
				return err
			}
			for _, msg := range list {
				if err := printAsYaml(msg); err != nil {
					return err
				}
			}
			return nil
		}
		msg, err := getSchema(args[0])
		if err != nil {
			return err
		}
		return printAsYaml(msg)
	},
}

func init() {
	schemaCmd.AddCommand(schemaGetCmd)
}

func getSchema(name string) (*v1.Schema, error) {
	cli, err := makeClient()
	if err != nil {
		return nil, err
	}
	return cli.V1().Schemas().Get(name)
}

func listSchemas() ([]*v1.Schema, error) {
	cli, err := makeClient()
	if err != nil {
		return nil, err
	}
	return cli.V1().Schemas().List()
}
