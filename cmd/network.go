// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/agentbay/agentbay-cli/internal/agentbay"
	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/agentbay/agentbay-cli/internal/config"
)

var NetworkCmd = &cobra.Command{
	Use:     "network",
	Short:   "Manage AgentBay network resources",
	Long:    "Query and manage network resources for AgentBay services.",
	GroupID: "management",
}

var NetworkPackageCmd = &cobra.Command{
	Use:   "package",
	Short: "Manage network packages",
	Long:  "Query and manage network packages for AgentBay services.",
}

var networkPackageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List network packages",
	Long: `List network packages for a specified user and region.

Both --user-ali-uid and --biz-region-id are required parameters.

Examples:
  # List network packages
  agentbay network package list --user-ali-uid 1234567890 --biz-region-id cn-hangzhou

  # List with verbose output
  agentbay network package list --user-ali-uid 1234567890 --biz-region-id cn-hangzhou -v`,
	RunE: runNetworkPackageList,
}

var (
	networkPackageUserAliUid  string
	networkPackageBizRegionId string
)

func init() {
	networkPackageListCmd.Flags().StringVar(&networkPackageUserAliUid, "user-ali-uid", "", "User Ali UID (required)")
	networkPackageListCmd.Flags().StringVar(&networkPackageBizRegionId, "biz-region-id", "", "Biz Region ID (required)")
	networkPackageListCmd.MarkFlagRequired("user-ali-uid")
	networkPackageListCmd.MarkFlagRequired("biz-region-id")

	NetworkPackageCmd.AddCommand(networkPackageListCmd)
	NetworkCmd.AddCommand(NetworkPackageCmd)
}

func runNetworkPackageList(cmd *cobra.Command, args []string) error {
	fmt.Printf("[LIST] Fetching network packages...\n")

	cfg, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to load configuration: %w", err)
	}

	apiClient := agentbay.NewClientFromConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &client.DescribeNetworkPackagesRequest{
		UserAliUid:  &networkPackageUserAliUid,
		BizRegionId: &networkPackageBizRegionId,
	}

	fmt.Printf("Requesting network packages...")
	resp, err := apiClient.DescribeNetworkPackages(ctx, req)
	if err != nil {
		printRequestIDFromErrIfVerbose(cmd, err)
		return fmt.Errorf("[ERROR] Failed to fetch network packages: %w", err)
	}
	fmt.Printf(" Done.")

	// Log Action and Request ID
	if resp.Body != nil && resp.Body.GetRequestId() != nil {
		fmt.Printf(" (Action: DescribeNetworkPackages, Request ID: %s)\n", *resp.Body.GetRequestId())
	} else {
		fmt.Printf(" (Action: DescribeNetworkPackages)\n")
	}

	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose && resp.Body != nil && resp.Body.GetRequestId() != nil && *resp.Body.GetRequestId() != "" {
		printRequestIDIfVerbose(cmd, *resp.Body.GetRequestId())
	}

	if resp.Body == nil {
		return fmt.Errorf("[ERROR] Invalid response: missing body")
	}

	if resp.Body.GetSuccess() != nil && !*resp.Body.GetSuccess() {
		errorMsg := "unknown error"
		if resp.Body.GetMessage() != nil {
			errorMsg = *resp.Body.GetMessage()
		}
		return fmt.Errorf("[ERROR] API request failed: %s", errorMsg)
	}

	data := resp.Body.GetData()
	if data == nil || len(data.Items) == 0 {
		fmt.Printf("\n[EMPTY] No network packages found.\n")
		return nil
	}

	items := data.Items
	fmt.Printf("\n[OK] Found %d network package(s)\n\n", len(items))

	// Print table header
	fmt.Printf("%s %s\n",
		padString("NETWORK PACKAGE ID", 40),
		"EIP ADDRESSES")
	fmt.Printf("%s %s\n",
		padString("------------------", 40),
		"-------------")

	// Print table rows
	for _, item := range items {
		if item == nil {
			continue
		}
		networkPackageId := item.GetNetworkPackageId()
		eipAddresses := item.GetEipAddresses()

		fmt.Printf("%s %s\n",
			padString(truncateString(networkPackageId, 40), 40),
			eipAddresses)
	}

	return nil
}
