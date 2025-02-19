package cmd

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	pluginInstallShort   = "Install required plugin images from your configuration"
	pluginInstallExample = `# Install required plugins specified in directory
cloudquery plugin install ./directory
# Install required plugins specified in directory and config files
cloudquery plugin install ./directory ./aws.yml ./pg.yml
`
)

func newCmdPluginInstall(deprecated bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install [files or directories]",
		Short:   pluginInstallShort,
		Long:    pluginInstallShort,
		Example: pluginInstallExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    installPlugin,
	}
	if deprecated {
		cmd.Deprecated = "use `cloudquery plugin install` instead"
	}
	return cmd
}

func installPlugin(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	sources := specReader.Sources
	destinations := specReader.Destinations

	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, destinations)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := auth.GetTeamForToken(authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name: %w", err)
	}
	opts := []managedplugin.Option{
		managedplugin.WithNoExec(),
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithAuthToken(authToken.Value),
		managedplugin.WithTeamName(teamName),
	}
	if cqDir != "" {
		opts = append(opts, managedplugin.WithDirectory(cqDir))
	}
	if disableSentry {
		opts = append(opts, managedplugin.WithNoSentry())
	}

	sourcePluginConfigs := make([]managedplugin.Config, len(sources))
	sourceRegInferred := make([]bool, len(sources))
	for i, source := range sources {
		sourcePluginConfigs[i] = managedplugin.Config{
			Name:     source.Name,
			Version:  source.Version,
			Path:     source.Path,
			Registry: SpecRegistryToPlugin(source.Registry),
		}
		sourceRegInferred[i] = source.RegistryInferred()
	}
	destinationPluginConfigs := make([]managedplugin.Config, len(destinations))
	destinationRegInferred := make([]bool, len(destinations))
	for i, destination := range destinations {
		destinationPluginConfigs[i] = managedplugin.Config{
			Name:     destination.Name,
			Version:  destination.Version,
			Path:     destination.Path,
			Registry: SpecRegistryToPlugin(destination.Registry),
		}
		destinationRegInferred[i] = destination.RegistryInferred()
	}
	if clist, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...); err != nil {
		return enrichClientError(clist, sourceRegInferred, err)
	}
	if clist, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...); err != nil {
		return enrichClientError(clist, destinationRegInferred, err)
	}

	return nil
}

// enrichClientError gets the index of the failed client (which is one more than the last one on the list) and checks if the registry was inferred.
// If so, adds a hint to the error message.
func enrichClientError(clientsList managedplugin.Clients, inferredList []bool, err error) error {
	if err == nil {
		return nil
	}
	if !strings.Contains(err.Error(), "not found") {
		return err
	}
	l := len(clientsList)
	il := len(inferredList)
	if l > il {
		return err // shouldn't happen
	}
	if !inferredList[l] {
		return err
	}

	return fmt.Errorf("%w. Hint: make sure to use the latest plugin version from hub.cloudquery.io or to keep using an outdated version add `registry: github` to your configuration", err)
}
