/*
Copyright © 2022 4strodev jumarurex@gmail.com

*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/4strodev/owl"
	owl_template "github.com/4strodev/owl/template"
	"github.com/spf13/cobra"
)

var (
	projectName     string = "default-project"
	projectTemplate string = "go-cli"
	moduleName      string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <PROJECT NAME>",
	Short: "Creates a new project",
	Long: `Creates a new command passing it a module, project name and a template.

It uses owl to generate projects. See https://github.com/4strodev/owl.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		projectName = args[0]
		createNewProject(projectName, projectTemplate)
	},
}

func createNewProject(name, template string) {
	// checking if name was given
	project := owl.NewProject(owl.ProjectConfig{
		Name:               projectName,
		TemplateName:       projectTemplate,
		LocalTemplatesDirs: []string{path.Join(os.Getenv("HOME"), ".raven/templates")},
		VerboseOutput:      Verbose,
		TempDir:            "/tmp/raven",
	}, owl_template.TemplateConfig{
		ConfigType: "toml",
		Context: map[string]any{
			"ModuleName": moduleName,
			"ProjectName": projectName,
		},
	})

	err := project.Create()
	if err != nil {
		switch err.Error() {
		case owl.DIR_EXISTS:
			fmt.Printf("Folder %s already exists\n", projectName)
			break
		case owl.TEMPLATE_NOT_FOUND:
			fmt.Printf("Template '%s' not found\n", project.Config.TemplateName)
			break
		default:
			fmt.Printf("Error creating project: %s\n", err)
			break
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringVarP(&projectTemplate, "template", "t", projectTemplate, "--template <TEMPLATE>")
	createCmd.Flags().StringVarP(&moduleName, "module", "m", "", "--module <MODULE NAME>")

	createCmd.MarkFlagRequired("module")
}
