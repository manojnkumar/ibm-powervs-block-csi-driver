/*
Copyright 2019 The Kubernetes Authors.

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

package main

//DONE

import (
	"flag"

	"github.com/ppc64le-cloud/powervs-csi-driver/pkg/driver"

	"k8s.io/klog/v2"
)

func main() {
	fs := flag.NewFlagSet("powervs-csi-driver", flag.ExitOnError)
	options := GetOptions(fs)

	drv, err := driver.NewDriver(
		driver.WithEndpoint(options.ServerOptions.Endpoint),
		//driver.WithExtraTags(options.ControllerOptions.ExtraTags),
		//river.WithExtraVolumeTags(options.ControllerOptions.ExtraVolumeTags),
		driver.WithMode(options.DriverMode),
		driver.WithDebug(options.ServerOptions.Debug),
		driver.WithVolumeAttachLimit(options.NodeOptions.VolumeAttachLimit),
		//driver.WithKubernetesClusterID(options.ControllerOptions.KubernetesClusterID),
	)
	if err != nil {
		klog.Fatalln(err)
	}
	if err := drv.Run(); err != nil {
		klog.Fatalln(err)
	}
}
