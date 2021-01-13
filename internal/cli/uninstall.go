package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/posener/complete"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/hashicorp/waypoint/internal/clierrors"
	"github.com/hashicorp/waypoint/internal/clisnapshot"
	"github.com/hashicorp/waypoint/internal/pkg/flag"
	"github.com/hashicorp/waypoint/internal/serverinstall"
)

type UninstallCommand struct {
	*baseCommand

	platform         string
	snapshotFilename string
	skipSnapshot     bool
	autoApprove      bool
	deleteContext    bool
}

func (c *UninstallCommand) Run(args []string) int {
	ctx := c.Ctx
	log := c.Log.Named("uninstall")
	defer c.Close()

	// Initialize. If we fail, we just exit since Init handles the UI.
	if err := c.Init(
		WithArgs(args),
		WithFlags(c.Flags()),
		WithNoConfig(),
	); err != nil {
		return 1
	}

	if !c.autoApprove {
		c.ui.Output(strings.TrimSpace(autoApproveMsg), terminal.WithErrorStyle())
		return 1
	}

	sg := c.ui.StepGroup()
	defer sg.Wait()

	// Pre-uninstall work
	// - name the context we'll be uninstalling
	// - generate a snapshot of the current install
	s := sg.Add("")
	defer func() { s.Abort() }()

	contextDefault, err := c.contextStorage.Default()
	if err != nil {
		c.ui.Output(clierrors.Humanize(err), terminal.WithErrorStyle())
		return 1
	}
	s.Update("Uninstalling Waypoint server with context %q", contextDefault)
	s.Done()

	// Generate a snapshot
	s = sg.Add("")
	if !c.skipSnapshot {
		s.Update("Generating server snapshot...")
		w, err := os.Create(c.snapshotFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating snapshot file: %s", err)
			return 1
		}
		if err = clisnapshot.WriteSnapshot(ctx, c.project.Client(), w); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating snapshot: %s", err)
			return 1
		}
		s.Update("Snapshot %q generated", c.snapshotFilename)
	} else {
		s.Update("skip-snapshot set; not generating server snapshot")
		s.Status(terminal.StatusWarn)
	}
	s.Done()

	// Uninstall
	p, ok := serverinstall.Platforms[strings.ToLower(c.platform)]
	if !ok {
		c.ui.Output(
			"Error uninstalling server from %s: invalid platform",
			c.platform,
			terminal.WithErrorStyle(),
		)

		return 1
	}

	err = p.Uninstall(ctx, &serverinstall.InstallOpts{
		Log: log,
		UI:  c.ui,
	})
	if err != nil {
		// point to current docs on manual server cleanup
		c.ui.Output(
			"Error uninstalling server from %s: %s", c.platform, clierrors.Humanize(err),
			terminal.WithErrorStyle(),
		)

		return 1
	}

	// Post-uninstall cleanup of context
	if c.deleteContext {
		if err := c.contextStorage.Delete(contextDefault); err != nil {
			c.ui.Output(clierrors.Humanize(err), terminal.WithErrorStyle())
			return 1
		}
	}

	c.ui.Output("Waypoint server successfully uninstalled for %s platform", c.platform, terminal.WithSuccessStyle())

	return 0
}

func (c *UninstallCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *UninstallCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *UninstallCommand) Synopsis() string {
	return "Uninstall the Waypoint server"
}

func (c *UninstallCommand) Help() string {
	return formatHelp(`
Usage: waypoint server uninstall [options]
	Uninstall the Waypoint server. The platform should be
	specified as kubernetes, nomad, or docker. '-auto-approve'
	is required.

  By default, this command deletes the default server's
  context.

  This command does not destroy Waypoint resources,
  such as deployments and releases. Clear all workspaces
  prior to uninstall to prevent hanging resources.

` + c.Flags().Help())
}

func (c *UninstallCommand) Flags() *flag.Sets {
	return c.flagSet(0, func(set *flag.Sets) {
		f := set.NewSet("Command Options")
		f.BoolVar(&flag.BoolVar{
			Name:    "auto-approve",
			Target:  &c.autoApprove,
			Default: false,
			Usage:   "Auto-approve server uninstallation. Default is false.",
		})

		f.BoolVar(&flag.BoolVar{
			Name:    "delete-context",
			Target:  &c.deleteContext,
			Default: true,
			Usage: "Delete the context for the server once it's uninstalled. " +
				"Default is true.",
		})

		f.StringVar(&flag.StringVar{
			Name:    "platform",
			Target:  &c.platform,
			Default: "",
			Usage:   "Platform to uninstall the Waypoint server from.",
		})

		f.StringVar(&flag.StringVar{
			Name:    "snapshot-filename",
			Target:  &c.snapshotFilename,
			Default: fmt.Sprintf("waypoint-sever-snapshot-%d", time.Now().Unix()),
			Usage:   "Filename to write the snapshot to.",
		})

		f.BoolVar(&flag.BoolVar{
			Name:    "skip-snapshot",
			Target:  &c.skipSnapshot,
			Default: false,
			Usage:   "Skip creating a snapshot of the Waypoint server. Default is false.",
		})

		for name, platform := range serverinstall.Platforms {
			platformSet := set.NewSet(name + " Options")
			platform.UninstallFlags(platformSet)
		}
	})
}

var (
	autoApproveMsg = strings.TrimSpace(`
Uninstalling Waypoint server requires approval. 
Rerun the command with -auto-approve to continue with the uninstall.
`)
)
