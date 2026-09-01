/*
Copyright 2026 The Kubernetes Authors.

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

package plugin

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/pflag"
)

//nolint:gochecknoinits // Plugins are loaded during package initialization.
func init() {
	loadPlugin(&driverNamePlugin{})
}

type driverNamePlugin struct {
	ebsCsiPluginBase

	driverName string
}

func (p *driverNamePlugin) InitFlags(fs *pflag.FlagSet) {
	fs.StringVar(&p.driverName, "driver-name", "ebs.csi.aws.com", "Driver name to use.")
}

func (p *driverNamePlugin) GetDriverName() string {
	return p.driverName
}

func (p *driverNamePlugin) Init(_ string, _ *prometheus.Registry) error {
	return nil
}
