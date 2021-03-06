/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyCheckCluster(
	pxMeta metav1.ObjectMeta,
	dbName string, podIndex int,
	clusterStats map[string]string) GomegaAsyncAssertion {

	return Eventually(
		func() bool {
			tunnel, en, err := f.GetEngine(pxMeta, dbName, podIndex)
			if err != nil {
				return false
			}
			defer tunnel.Close()
			defer en.Close()

			//rows := make([]map[string]string, 0)
			rows, err := en.QueryString("show status like \"wsrep%%\"")
			if err != nil {
				return false
			}

			ch := true
			for _, m := range rows {
				if m["Variable_name"] == "wsrep_local_state" {
					ch = ch && m["Value"] == clusterStats["wsrep_local_state"]
				}
				if m["Variable_name"] == "wsrep_local_state_comment" {
					ch = ch && m["Value"] == clusterStats["wsrep_local_state_comment"]
				}
				if m["Variable_name"] == "wsrep_evs_state" {
					ch = ch && m["Value"] == clusterStats["wsrep_evs_state"]
				}
				if m["Variable_name"] == "wsrep_cluster_size" {
					ch = ch && m["Value"] == clusterStats["wsrep_cluster_size"]
				}
				if m["Variable_name"] == "wsrep_cluster_status" {
					ch = ch && m["Value"] == clusterStats["wsrep_cluster_status"]
				}
				if m["Variable_name"] == "wsrep_connected" {
					ch = ch && m["Value"] == clusterStats["wsrep_connected"]
				}
				if m["Variable_name"] == "wsrep_ready" {
					ch = ch && m["Value"] == clusterStats["wsrep_ready"]
				}
			}

			return ch
		},
		time.Minute*10,
		time.Second*20,
	)
}

func (f *Framework) RemoverPrimary(meta metav1.ObjectMeta, primaryPodIndex int) error {
	if _, err := f.kubeClient.CoreV1().Pods(meta.Namespace).Get(
		context.TODO(),
		fmt.Sprintf("%s-%d", meta.Name, primaryPodIndex),
		metav1.GetOptions{},
	); err != nil {
		return err
	}

	return f.kubeClient.CoreV1().Pods(meta.Namespace).Delete(
		context.TODO(),
		fmt.Sprintf("%s-%d", meta.Name, primaryPodIndex),
		metav1.DeleteOptions{})
}
