package list

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aziontech/azion-cli/utils"
	sdk "github.com/aziontech/edgeservices-go-sdk"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	// listCmd represents the list command
	listCmd := &cobra.Command{
		Use:           "list",
		Short:         "Lists resources in a given service",
		Long:          `Lists all resources found in a service by providing a service_id. Service_id can be found by listing your services`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return utils.ErrorMissingServiceIdArgument
			}

			service_id_arg := args[0]
			service_id, err := strconv.Atoi(service_id_arg)
			if err != nil {
				return utils.ErrorConvertingIdArgumentToInt
			}

			client, err := utils.CreateClient()
			if err != nil {
				return err
			}
			if err := listAllResources(client, service_id); err != nil {
				return err
			}
			return nil
		},
	}
	return listCmd
}

func listAllResources(client *sdk.APIClient, service_id int) error {
	c := context.Background()
	api := client.DefaultApi

	resp, httpResp, err := api.GetResources(c, int64(service_id)).Execute()
	if err != nil {
		if httpResp.StatusCode >= 500 {
			return utils.ErrorInternalServerError
		}
		return err
	}

	resources := resp.Resources

	for _, resource := range resources {
		fmt.Printf("ID: %d     Name: %s \n", resource.Id, resource.Name)
	}
	return nil
}
