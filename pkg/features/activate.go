/*
Copyright 2019 The KubeOne Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package features

import (
	"github.com/pkg/errors"

	kubeoneapi "github.com/kubermatic/kubeone/pkg/apis/kubeone"
	"github.com/kubermatic/kubeone/pkg/state"
	"github.com/kubermatic/kubeone/pkg/templates/kubeadm/kubeadmargs"
)

// Activate configured features.
// Installing CRDs, creating policies and so on
func Activate(s *state.State) error {
	if err := installKubeSystemPSP(s.Cluster.Features.PodSecurityPolicy, s); err != nil {
		return errors.Wrap(err, "failed to install PodSecurityPolicy")
	}

	if err := installMetricsServer(s.Cluster.Features.MetricsServer, s); err != nil {
		return errors.Wrap(err, "failed to install metrics-server")
	}

	if err := installNodeLocalDNSCache(s.Cluster.Features.NodeLocalDNSCache, s); err != nil {
		return errors.Wrap(err, "failed to install nodeLocal DNSCache")
	}

	return nil
}

// UpdateKubeadmClusterConfiguration update additional config options in the kubeadm's
// v1beta1.ClusterConfiguration according to enabled features
func UpdateKubeadmClusterConfiguration(featuresCfg kubeoneapi.Features, args *kubeadmargs.Args) {
	updatePSPKubeadmConfig(featuresCfg.PodSecurityPolicy, args)
	updateDynamicAuditLogsKubeadmConfig(featuresCfg.DynamicAuditLog, args)
	updateOIDCKubeadmConfig(featuresCfg.OpenIDConnect, args)
}
