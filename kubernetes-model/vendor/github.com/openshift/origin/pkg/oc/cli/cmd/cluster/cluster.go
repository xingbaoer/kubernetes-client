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
package cluster

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"

	"github.com/openshift/origin/pkg/oc/bootstrap/docker"
	"github.com/openshift/origin/pkg/oc/cli/util/clientcmd"
)

const ClusterRecommendedName = "cluster"

var (
	clusterLong = templates.LongDesc(`
		Manage a local OpenShift cluster

		The OpenShift cluster will run as an all-in-one container on a Docker host. The Docker host
		may be a local VM (ie. using docker-machine on OS X and Windows clients), remote machine, or
		the local Unix host.

		Use the 'up' command to start a new cluster (master and node) on a single machine. Use the
		'join' command on another machine to connect to the first cluster.

		To use an existing Docker connection, ensure that Docker commands are working and that you
		can create new containers. For OS X and Windows clients, a docker-machine with the VirtualBox
		driver can be created for you using the --create-machine option.

		By default, etcd data will not be preserved between container restarts. If you wish to
		preserve your data, specify a value for --host-data-dir and the --use-existing-config flag.

		Default routes are setup using nip.io and the host ip of your cluster. To use a different
		routing suffix, use the --routing-suffix flag.`)
)

func NewCmdCluster(name, fullName string, f *clientcmd.Factory, in io.Reader, out, errout io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   fmt.Sprintf("%s ACTION", name),
		Short: "Start and stop OpenShift cluster",
		Long:  clusterLong,
		Run:   cmdutil.DefaultSubCommandRun(errout),
	}

	cmds.AddCommand(docker.NewCmdUp(docker.CmdUpRecommendedName, fullName+" "+docker.CmdUpRecommendedName, f, out, errout))
	cmds.AddCommand(docker.NewCmdJoin(docker.CmdJoinRecommendedName, fullName+" "+docker.CmdJoinRecommendedName, f, in, out))
	cmds.AddCommand(docker.NewCmdDown(docker.CmdDownRecommendedName, fullName+" "+docker.CmdDownRecommendedName, f, out))
	cmds.AddCommand(docker.NewCmdStatus(docker.CmdStatusRecommendedName, fullName+" "+docker.CmdStatusRecommendedName, f, out))
	return cmds
}
