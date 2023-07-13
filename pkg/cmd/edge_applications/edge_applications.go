package edge_applications

import (
	"github.com/MakeNowJust/heredoc"
	msg "github.com/aziontech/azion-cli/messages/edge_applications"
	buildCmd "github.com/aziontech/azion-cli/pkg/cmd/edge_applications/build"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/delete"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/describe"
	initCmd "github.com/aziontech/azion-cli/pkg/cmd/edge_applications/init"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/list"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/ls"
	publishCmd "github.com/aziontech/azion-cli/pkg/cmd/edge_applications/publish"
	"github.com/aziontech/azion-cli/pkg/cmd/edge_applications/update"
	"github.com/aziontech/azion-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	edge_applicationsCmd := &cobra.Command{
		Use:   msg.EdgeApplicationsUsage,
		Short: msg.EdgeApplicationsShortDescription,
		Long:  msg.EdgeApplicationsLongDescription,
		Example: heredoc.Doc(`
		$ azioncli edge_applications --help
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	edge_applicationsCmd.AddCommand(initCmd.NewCmd(f))
	edge_applicationsCmd.AddCommand(buildCmd.NewCmd(f))
	edge_applicationsCmd.AddCommand(publishCmd.NewCmd(f))
	edge_applicationsCmd.AddCommand(describe.NewCmd(f))
	edge_applicationsCmd.AddCommand(delete.NewCmd(f))
	edge_applicationsCmd.AddCommand(list.NewCmd(f))
	edge_applicationsCmd.AddCommand(update.NewCmd(f))
	edge_applicationsCmd.AddCommand(ls.NewCmd(f))

	edge_applicationsCmd.Flags().BoolP("help", "h", false, msg.EdgeApplicationsFlagHelp)

	return edge_applicationsCmd
}
