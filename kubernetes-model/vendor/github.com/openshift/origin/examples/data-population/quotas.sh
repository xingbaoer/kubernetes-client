#!/bin/bash
#
# Copyright (C) 2015 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#         http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


# Quotas

# Populates each project with a quota

source $(dirname "${BASH_SOURCE}")/common.sh

echo "Populating quotas"

export KUBECONFIG=${OPENSHIFT_ADMIN_CONFIG}

QUOTA=$(dirname "${BASH_SOURCE}")/quota.yaml

for ((i=1; i <=$NUM_PROJECTS; i++))
do
  oc create -f $QUOTA --namespace=${PROJECT_NAME_PREFIX}${i}
done

echo "Done"