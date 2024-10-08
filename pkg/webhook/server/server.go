/*
Copyright 2023 The KusionStack Authors.

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

package server

import (
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"kusionstack.io/kuperator/pkg/webhook/server/generic"
)

// Add adds itself to the manager
func Add(mgr manager.Manager) error {
	server := mgr.GetWebhookServer()
	logger := mgr.GetLogger().WithName("webhook")

	// register admission handlers
	for name, handler := range generic.HandlerMap {
		if len(name) == 0 {
			logger.Info("Skip registering handlers without a name")
			continue
		}

		path := name
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		server.Register(path, &webhook.Admission{Handler: handler})
		logger.V(3).Info("Registered webhook handler", "path", path)
	}

	return nil
}
