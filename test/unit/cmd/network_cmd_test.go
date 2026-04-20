// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd_test

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/agentbay/agentbay-cli/cmd"
)

func TestNetworkCmd(t *testing.T) {
	t.Run("network command has correct metadata", func(t *testing.T) {
		assert.Equal(t, "network", cmd.NetworkCmd.Use)
		assert.Equal(t, "Manage AgentBay network resources", cmd.NetworkCmd.Short)
		assert.Equal(t, "management", cmd.NetworkCmd.GroupID)
		assert.True(t, strings.Contains(cmd.NetworkCmd.Long, "network"))
	})

	t.Run("network has package subcommand", func(t *testing.T) {
		children := cmd.NetworkCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "package")
	})
}

func TestNetworkPackageCmd(t *testing.T) {
	t.Run("package command has correct metadata", func(t *testing.T) {
		assert.NotNil(t, cmd.NetworkPackageCmd)
		assert.Equal(t, "package", cmd.NetworkPackageCmd.Use)
		assert.Equal(t, "Manage network packages", cmd.NetworkPackageCmd.Short)
	})

	t.Run("package has list subcommand", func(t *testing.T) {
		children := cmd.NetworkPackageCmd.Commands()
		names := make([]string, len(children))
		for i, c := range children {
			names[i] = c.Name()
		}
		assert.Contains(t, names, "list")
	})
}

func TestNetworkPackageListCmd(t *testing.T) {
	t.Run("list command has correct metadata", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.NetworkPackageCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)
		assert.Equal(t, "list", listCmd.Use)
		assert.Equal(t, "List network packages", listCmd.Short)
		assert.True(t, strings.Contains(listCmd.Long, "network packages"))
	})

	t.Run("list command has required user-ali-uid flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.NetworkPackageCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)

		flag := listCmd.Flags().Lookup("user-ali-uid")
		assert.NotNil(t, flag)
		assert.Equal(t, "", flag.DefValue)
		assert.True(t, strings.Contains(flag.Usage, "required"))
	})

	t.Run("list command has required biz-region-id flag", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.NetworkPackageCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)

		flag := listCmd.Flags().Lookup("biz-region-id")
		assert.NotNil(t, flag)
		assert.Equal(t, "", flag.DefValue)
		assert.True(t, strings.Contains(flag.Usage, "required"))
	})

	t.Run("list command flags have no default values", func(t *testing.T) {
		var listCmd *cobra.Command
		for _, c := range cmd.NetworkPackageCmd.Commands() {
			if c.Name() == "list" {
				listCmd = c
				break
			}
		}

		assert.NotNil(t, listCmd)

		userAliUidFlag := listCmd.Flags().Lookup("user-ali-uid")
		assert.NotNil(t, userAliUidFlag)
		assert.Equal(t, "", userAliUidFlag.DefValue)

		bizRegionIdFlag := listCmd.Flags().Lookup("biz-region-id")
		assert.NotNil(t, bizRegionIdFlag)
		assert.Equal(t, "", bizRegionIdFlag.DefValue)
	})
}
