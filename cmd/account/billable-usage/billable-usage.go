// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billable_usage

import (
	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/spf13/cobra"
)

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var cmdOverrides []func(*cobra.Command)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billable-usage",
		Short: `This API allows you to download billable usage logs for the specified account and date range.`,
		Long: `This API allows you to download billable usage logs for the specified account
  and date range. This feature works with all account types.`,
		GroupID: "billing",
		Annotations: map[string]string{
			"package": "billing",
		},
	}

	// Apply optional overrides to this command.
	for _, fn := range cmdOverrides {
		fn(cmd)
	}

	return cmd
}

// start download command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var downloadOverrides []func(
	*cobra.Command,
	*billing.DownloadRequest,
)

func newDownload() *cobra.Command {
	cmd := &cobra.Command{}

	var downloadReq billing.DownloadRequest

	// TODO: short flags

	cmd.Flags().BoolVar(&downloadReq.PersonalData, "personal-data", downloadReq.PersonalData, `Specify whether to include personally identifiable information in the billable usage logs, for example the email addresses of cluster creators.`)

	cmd.Use = "download START_MONTH END_MONTH"
	cmd.Short = `Return billable usage logs.`
	cmd.Long = `Return billable usage logs.
  
  Returns billable usage logs in CSV format for the specified account and date
  range. For the data schema, see [CSV file schema]. Note that this method might
  take multiple minutes to complete.
  
  **Warning**: Depending on the queried date range, the number of workspaces in
  the account, the size of the response and the internet speed of the caller,
  this API may hit a timeout after a few minutes. If you experience this, try to
  mitigate by calling the API with narrower date ranges.
  
  [CSV file schema]: https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(2)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustAccountClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		a := root.AccountClient(ctx)

		downloadReq.StartMonth = args[0]
		downloadReq.EndMonth = args[1]

		response, err := a.BillableUsage.Download(ctx, downloadReq)
		if err != nil {
			return err
		}
		defer response.Contents.Close()
		return cmdio.RenderReader(ctx, response.Contents)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range downloadOverrides {
		fn(cmd, &downloadReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newDownload())
	})
}

// end service BillableUsage
