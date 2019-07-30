/*
 Copyright 2019 IXON B.V.

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

import (
	"github.com/go-acme/lego/providers/dns/cloudns"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
	log "github.com/sirupsen/logrus"
	restclient "k8s.io/client-go/rest"
	"os"
)

const ProviderName = "cloudns"

var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		log.Fatal("Please set the GROUP_NAME env variable.")
	}

	cmd.RunWebhookServer(GroupName,
		&clouDNSProviderSolver{},
	)
}

type clouDNSProviderSolver struct {
	provider *cloudns.DNSProvider
}

func (c clouDNSProviderSolver) Name() string {
	return ProviderName
}

func (c clouDNSProviderSolver) Present(ch *v1alpha1.ChallengeRequest) error {
	return c.provider.Present(ch.ResolvedFQDN, "", ch.Key)
}

func (c clouDNSProviderSolver) CleanUp(ch *v1alpha1.ChallengeRequest) error {
	return c.provider.CleanUp(ch.ResolvedFQDN, "", ch.Key)
}

func (c clouDNSProviderSolver) Initialize(kubeClientConfig *restclient.Config, stopCh <-chan struct{}) error {
	provider, err := cloudns.NewDNSProvider()

	if err != nil {
		return err
	}

	c.provider = provider

	return nil
}
