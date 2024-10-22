/*
Copyright 2024 The Kubeflow Authors.

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

package core

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	runtime "github.com/kubeflow/training-operator/pkg/runtime.v2"
)

type Registry map[string]func(ctx context.Context, client client.Client, indexer client.FieldIndexer) (runtime.Runtime, error)

func NewRuntimeRegistry() Registry {
	return Registry{
		TrainingRuntimeGroupKind:        NewTrainingRuntime,
		ClusterTrainingRuntimeGroupKind: NewClusterTrainingRuntime,
	}
}