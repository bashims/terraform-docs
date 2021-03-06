/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package hcl

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'tfvars hcl' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "hcl [PATH]",
		Short:       "Generate HCL format of terraform.tfvars of inputs",
		Annotations: cli.Annotations("tfvars hcl"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}
	cmd.PersistentFlags().BoolVar(&config.Settings.Description, "description", false, "show Descriptions on variables")
	return cmd
}
