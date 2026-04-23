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
	Long: `List network packages for a specified region.

Examples:
  # List network packages (uses default region cn-hangzhou)
  agentbay network package list

  # List network packages for a specific region
  agentbay network package list --biz-region-id cn-shanghai

  # List with verbose output
  agentbay network package list --biz-region-id cn-hangzhou -v`,
	RunE: runNetworkPackageList,
}

var (
	networkPackageBizRegionId string
)

func init() {
	networkPackageListCmd.Flags().StringVar(&networkPackageBizRegionId, "biz-region-id", "cn-hangzhou", "Biz Region ID (default: cn-hangzhou)")

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
	fmt.Printf("%s %s %s\n",
		padString("NETWORK PACKAGE ID", 28),
		padString("OFFICE SITE ID", 34),
		"EIP ADDRESSES")
	fmt.Printf("%s %s %s\n",
		padString("------------------", 28),
		padString("--------------", 34),
		"-------------")

	// Print table rows
	for _, item := range items {
		if item == nil {
			continue
		}
		networkPackageId := item.GetNetworkPackageId()
		officeSiteId := item.GetOfficeSiteId()
		eipAddresses := item.GetEipAddresses()

		fmt.Printf("%s %s %s\n",
			padString(truncateString(networkPackageId, 28), 28),
			padString(truncateString(officeSiteId, 34), 34),
			eipAddresses)
	}

	return nil
}
