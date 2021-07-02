/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-digitalocean-api/client/clientset/versioned/typed/database/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDatabaseV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDatabaseV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeDatabaseV1alpha1) ConnectionPools(namespace string) v1alpha1.ConnectionPoolInterface {
	return &FakeConnectionPools{c, namespace}
}

func (c *FakeDatabaseV1alpha1) Dbs(namespace string) v1alpha1.DbInterface {
	return &FakeDbs{c, namespace}
}

func (c *FakeDatabaseV1alpha1) Firewalls(namespace string) v1alpha1.FirewallInterface {
	return &FakeFirewalls{c, namespace}
}

func (c *FakeDatabaseV1alpha1) Replicas(namespace string) v1alpha1.ReplicaInterface {
	return &FakeReplicas{c, namespace}
}

func (c *FakeDatabaseV1alpha1) Users(namespace string) v1alpha1.UserInterface {
	return &FakeUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDatabaseV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}