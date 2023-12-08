// Copyright © 2021 The Knative Authors
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

package main

import (
	"fmt"
	"os"

	"knative.dev/kn-plugin-sink-kamelet/internal/root"
)

func main() {
	err := root.NewSinkKameletCommand().Execute()
	if err != nil {
		if err.Error() != "subcommand is required" {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
	return
}
