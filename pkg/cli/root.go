// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"
	topo "github.com/onosproject/onos-topo/pkg/cli"
	"github.com/spf13/cobra"
	"os"
)

// Execute runs the root command and any sub-commands.
func Execute() {
	rootCmd := GetRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetRootCommand returns the root onos command
func GetRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                    "onos",
		Short:                  "ONOS command line client",
		BashCompletionFunction: getBashCompletions(),
	}
	cmd.AddCommand(topo.GetCommand())
	cmd.AddCommand(getCompletionCommand())
	return cmd
}
