/*
Copyright 2021 The Fluid Authors.

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

package kubeclient

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetPersistentVolumeClaim(client client.Reader, name, namespace string) (pvc *v1.PersistentVolumeClaim, err error) {
	pvc = &v1.PersistentVolumeClaim{}
	err = client.Get(context.TODO(),
		types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
		pvc)
	return
}
