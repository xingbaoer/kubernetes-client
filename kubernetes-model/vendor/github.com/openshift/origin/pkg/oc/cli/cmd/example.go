/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"

	"github.com/openshift/origin/pkg/oc/cli/util/clientcmd"
)

var (
	internalTYPELong = templates.LongDesc(`
		Single line title

		Description body`)

	internalTYPEExample = templates.Examples(`%s`)
)

type TYPEOptions struct {
	In          io.Reader
	Out, ErrOut io.Writer
}

// NewCmdTYPE implements a TYPE command
// This is an example type for templating.
func NewCmdTYPE(fullName string, f *clientcmd.Factory, in io.Reader, out, errout io.Writer) *cobra.Command {
	options := &TYPEOptions{
		In:     in,
		Out:    out,
		ErrOut: errout,
	}
	cmd := &cobra.Command{
		Use:     "NAME [...]",
		Short:   "A short description",
		Long:    internalTYPELong,
		Example: fmt.Sprintf(internalTYPEExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			kcmdutil.CheckErr(options.Complete(f, cmd, args))
			kcmdutil.CheckErr(options.Validate())
			if err := options.Run(); err != nil {
				// TODO: move me to kcmdutil
				if err == kcmdutil.ErrExit {
					os.Exit(1)
				}
				kcmdutil.CheckErr(err)
			}
		},
	}
	return cmd
}

func (o *TYPEOptions) Complete(f *clientcmd.Factory, c *cobra.Command, args []string) error {
	return nil
}

func (o *TYPEOptions) Validate() error { return nil }
func (o *TYPEOptions) Run() error      { return nil }
