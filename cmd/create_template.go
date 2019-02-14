package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var createTemplateOpts options.TemplateCreate

// createTemplateCmd represents the createTemplate command
var createTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a cluster template for provisioning",
	Long: `Create a template for provisioning

If creating a template from a file, do not specify a provider to use and instead simply run:

csctl create template -f <filename>

This file must be json (TODO support yaml). All other template-specific flags will be ignored if -f is used.

Otherwise, please use a provider subcommand to create a template with more advanced, provider-specific flags.`,

	// TODO using PersistentPreRunE here to org-scope all create template
	// subcommands would be nice, but it won't work currently because the
	// root command is already using it and this cobra issue is open:
	// https://github.com/spf13/cobra/issues/252
	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open(filename)
		if err != nil {
			return errors.Wrap(err, "opening file")
		}
		defer f.Close()

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			return errors.Wrap(err, "reading file")
		}

		var req types.CreateTemplateRequest

		err = json.Unmarshal(bytes, &req)
		if err != nil {
			return errors.Wrap(err, "unmarshalling file into request type")
		}

		template, err := clientset.Provision().Templates(organizationID).Create(&req)
		if err != nil {
			return err
		}

		t := resource.NewTemplates([]types.Template{*template})
		return t.Table(os.Stdout)
	},
}

func init() {
	createCmd.AddCommand(createTemplateCmd)

	createTemplateCmd.Flags().StringVarP(&filename, "filename", "f", "", "create a template from the given file (TODO must be json for now)")

	// No defaulting is performed here because the logic in many cases is nontrivial,
	// and we'd like to be consistent with where and how we default.
	createTemplateCmd.PersistentFlags().Int32VarP(&createTemplateOpts.MasterCount, "master-count", "m", 0, "number of nodes in master node pool")
	createTemplateCmd.PersistentFlags().Int32VarP(&createTemplateOpts.WorkerCount, "worker-count", "w", 0, "number of nodes in worker node pool")

	createTemplateCmd.PersistentFlags().StringVar(&createTemplateOpts.MasterKubernetesVersion, "master-kubernetes-version", "", "Kubernetes version for master node pool")
	createTemplateCmd.PersistentFlags().StringVar(&createTemplateOpts.WorkerKubernetesVersion, "worker-kubernetes-version", "", "Kubernetes version for worker node pool")

	createTemplateCmd.PersistentFlags().StringVar(&createTemplateOpts.DockerVersion, "docker-version", "", "Docker version for all node pools")

	createTemplateCmd.PersistentFlags().StringVar(&createTemplateOpts.Description, "description", "", "template description")
}
