// Copyright 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	resourceclientset "cisco-app-networking.github.io/wcm-common/app-dns-publisher/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	regexIPv4 = `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`
	regexIPv6 = `((?:[0-9A-Fa-f]{1,4}))((?::[0-9A-Fa-f]{1,4}))*::((?:[0-9A-Fa-f]{1,4}))((?::[0-9A-Fa-f]{1,4}))*|((?:[0-9A-Fa-f]{1,4}))((?::[0-9A-Fa-f]{1,4})){7}`
)

var (
	targetsSpec = make(map[string][]string)
	addValue    = make([]string, 0)
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", "/home/ubuntu/kubeconfigs/central/kind-1.kubeconfig")
	resourceclient, err := resourceclientset.NewForConfig(config)
	lop := "helloworld.cluster-0.wcm-cisco.com"
	dnsCrd, err := resourceclient.ExternaldnsV1alpha1().DNSEndpoints("wcm-system").Get(context.TODO(), lop, metav1.GetOptions{})

	if err != nil {
		log.Fatalf("%v", err)
	}
	dnsValue, _ := json.Marshal(dnsCrd)
	rexp := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	submatchall := rexp.FindAllString(string(dnsValue), -1)
	for _, element := range submatchall {
		addValue = targetsSpec[lop]
		addValue = append(addValue, element)
		targetsSpec[lop] = updateTarget

	}
	fmt.Println(targetsSpec)
}
