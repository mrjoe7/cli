// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package external_locations

import (
	"fmt"

	"github.com/databricks/cli/cmd/root"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/databricks/cli/libs/flags"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/spf13/cobra"
)

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var cmdOverrides []func(*cobra.Command)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "external-locations",
		Short: `An external location is an object that combines a cloud storage path with a storage credential that authorizes access to the cloud storage path.`,
		Long: `An external location is an object that combines a cloud storage path with a
  storage credential that authorizes access to the cloud storage path. Each
  external location is subject to Unity Catalog access-control policies that
  control which users and groups can access the credential. If a user does not
  have access to an external location in Unity Catalog, the request fails and
  Unity Catalog does not attempt to authenticate to your cloud tenant on the
  user’s behalf.
  
  Databricks recommends using external locations rather than using storage
  credentials directly.
  
  To create external locations, you must be a metastore admin or a user with the
  **CREATE_EXTERNAL_LOCATION** privilege.`,
		GroupID: "catalog",
		Annotations: map[string]string{
			"package": "catalog",
		},
	}

	// Apply optional overrides to this command.
	for _, fn := range cmdOverrides {
		fn(cmd)
	}

	return cmd
}

// start create command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var createOverrides []func(
	*cobra.Command,
	*catalog.CreateExternalLocation,
)

func newCreate() *cobra.Command {
	cmd := &cobra.Command{}

	var createReq catalog.CreateExternalLocation
	var createJson flags.JsonFlag

	// TODO: short flags
	cmd.Flags().Var(&createJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	cmd.Flags().StringVar(&createReq.AccessPoint, "access-point", createReq.AccessPoint, `The AWS access point to use when accesing s3 for this external location.`)
	cmd.Flags().StringVar(&createReq.Comment, "comment", createReq.Comment, `User-provided free-form text description.`)
	// TODO: complex arg: encryption_details
	cmd.Flags().BoolVar(&createReq.ReadOnly, "read-only", createReq.ReadOnly, `Indicates whether the external location is read-only.`)
	cmd.Flags().BoolVar(&createReq.SkipValidation, "skip-validation", createReq.SkipValidation, `Skips validation of the storage credential associated with the external location.`)

	cmd.Use = "create NAME URL CREDENTIAL_NAME"
	cmd.Short = `Create an external location.`
	cmd.Long = `Create an external location.
  
  Creates a new external location entry in the metastore. The caller must be a
  metastore admin or have the **CREATE_EXTERNAL_LOCATION** privilege on both the
  metastore and the associated storage credential.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("json") {
			err := cobra.ExactArgs(0)(cmd, args)
			if err != nil {
				return fmt.Errorf("when --json flag is specified, no positional arguments are required. Provide 'name', 'url', 'credential_name' in your JSON input")
			}
			return nil
		}
		check := cobra.ExactArgs(3)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if cmd.Flags().Changed("json") {
			err = createJson.Unmarshal(&createReq)
			if err != nil {
				return err
			}
		}
		if !cmd.Flags().Changed("json") {
			createReq.Name = args[0]
		}
		if !cmd.Flags().Changed("json") {
			createReq.Url = args[1]
		}
		if !cmd.Flags().Changed("json") {
			createReq.CredentialName = args[2]
		}

		response, err := w.ExternalLocations.Create(ctx, createReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range createOverrides {
		fn(cmd, &createReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newCreate())
	})
}

// start delete command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var deleteOverrides []func(
	*cobra.Command,
	*catalog.DeleteExternalLocationRequest,
)

func newDelete() *cobra.Command {
	cmd := &cobra.Command{}

	var deleteReq catalog.DeleteExternalLocationRequest

	// TODO: short flags

	cmd.Flags().BoolVar(&deleteReq.Force, "force", deleteReq.Force, `Force deletion even if there are dependent external tables or mounts.`)

	cmd.Use = "delete NAME"
	cmd.Short = `Delete an external location.`
	cmd.Long = `Delete an external location.
  
  Deletes the specified external location from the metastore. The caller must be
  the owner of the external location.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		deleteReq.Name = args[0]

		err = w.ExternalLocations.Delete(ctx, deleteReq)
		if err != nil {
			return err
		}
		return nil
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range deleteOverrides {
		fn(cmd, &deleteReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newDelete())
	})
}

// start get command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var getOverrides []func(
	*cobra.Command,
	*catalog.GetExternalLocationRequest,
)

func newGet() *cobra.Command {
	cmd := &cobra.Command{}

	var getReq catalog.GetExternalLocationRequest

	// TODO: short flags

	cmd.Use = "get NAME"
	cmd.Short = `Get an external location.`
	cmd.Long = `Get an external location.
  
  Gets an external location from the metastore. The caller must be either a
  metastore admin, the owner of the external location, or a user that has some
  privilege on the external location.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		getReq.Name = args[0]

		response, err := w.ExternalLocations.Get(ctx, getReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range getOverrides {
		fn(cmd, &getReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newGet())
	})
}

// start list command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var listOverrides []func(
	*cobra.Command,
)

func newList() *cobra.Command {
	cmd := &cobra.Command{}

	cmd.Use = "list"
	cmd.Short = `List external locations.`
	cmd.Long = `List external locations.
  
  Gets an array of external locations (__ExternalLocationInfo__ objects) from
  the metastore. The caller must be a metastore admin, the owner of the external
  location, or a user that has some privilege on the external location. There is
  no guarantee of a specific ordering of the elements in the array.`

	cmd.Annotations = make(map[string]string)

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)
		response, err := w.ExternalLocations.ListAll(ctx)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range listOverrides {
		fn(cmd)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newList())
	})
}

// start update command

// Slice with functions to override default command behavior.
// Functions can be added from the `init()` function in manually curated files in this directory.
var updateOverrides []func(
	*cobra.Command,
	*catalog.UpdateExternalLocation,
)

func newUpdate() *cobra.Command {
	cmd := &cobra.Command{}

	var updateReq catalog.UpdateExternalLocation
	var updateJson flags.JsonFlag

	// TODO: short flags
	cmd.Flags().Var(&updateJson, "json", `either inline JSON string or @path/to/file.json with request body`)

	cmd.Flags().StringVar(&updateReq.AccessPoint, "access-point", updateReq.AccessPoint, `The AWS access point to use when accesing s3 for this external location.`)
	cmd.Flags().StringVar(&updateReq.Comment, "comment", updateReq.Comment, `User-provided free-form text description.`)
	cmd.Flags().StringVar(&updateReq.CredentialName, "credential-name", updateReq.CredentialName, `Name of the storage credential used with this location.`)
	// TODO: complex arg: encryption_details
	cmd.Flags().BoolVar(&updateReq.Force, "force", updateReq.Force, `Force update even if changing url invalidates dependent external tables or mounts.`)
	cmd.Flags().StringVar(&updateReq.Name, "name", updateReq.Name, `Name of the external location.`)
	cmd.Flags().StringVar(&updateReq.Owner, "owner", updateReq.Owner, `The owner of the external location.`)
	cmd.Flags().BoolVar(&updateReq.ReadOnly, "read-only", updateReq.ReadOnly, `Indicates whether the external location is read-only.`)
	cmd.Flags().BoolVar(&updateReq.SkipValidation, "skip-validation", updateReq.SkipValidation, `Skips validation of the storage credential associated with the external location.`)
	cmd.Flags().StringVar(&updateReq.Url, "url", updateReq.Url, `Path URL of the external location.`)

	cmd.Use = "update NAME"
	cmd.Short = `Update an external location.`
	cmd.Long = `Update an external location.
  
  Updates an external location in the metastore. The caller must be the owner of
  the external location, or be a metastore admin. In the second case, the admin
  can only update the name of the external location.`

	cmd.Annotations = make(map[string]string)

	cmd.Args = func(cmd *cobra.Command, args []string) error {
		check := cobra.ExactArgs(1)
		return check(cmd, args)
	}

	cmd.PreRunE = root.MustWorkspaceClient
	cmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()
		w := root.WorkspaceClient(ctx)

		if cmd.Flags().Changed("json") {
			err = updateJson.Unmarshal(&updateReq)
			if err != nil {
				return err
			}
		}
		updateReq.Name = args[0]

		response, err := w.ExternalLocations.Update(ctx, updateReq)
		if err != nil {
			return err
		}
		return cmdio.Render(ctx, response)
	}

	// Disable completions since they are not applicable.
	// Can be overridden by manual implementation in `override.go`.
	cmd.ValidArgsFunction = cobra.NoFileCompletions

	// Apply optional overrides to this command.
	for _, fn := range updateOverrides {
		fn(cmd, &updateReq)
	}

	return cmd
}

func init() {
	cmdOverrides = append(cmdOverrides, func(cmd *cobra.Command) {
		cmd.AddCommand(newUpdate())
	})
}

// end service ExternalLocations
