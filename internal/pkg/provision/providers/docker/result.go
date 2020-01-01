// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package docker

import (
	"github.com/talos-systems/talos/cmd/osctl/pkg/client/config"
	"github.com/talos-systems/talos/internal/pkg/provision"
)

type result struct {
	talosConfig *config.Config

	clusterInfo provision.ClusterInfo
}

func (res *result) TalosConfig() *config.Config {
	return res.talosConfig
}

func (res *result) Info() provision.ClusterInfo {
	return res.clusterInfo
}