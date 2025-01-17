/*
Copyright 2021 The Knative Authors

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

package certificates

const (
	Organization        = "knative.dev"
	LegacyFakeDnsName   = "data-plane." + Organization
	FakeDnsName         = LegacyFakeDnsName
	DataPlaneNamePrefix = "knative-"

	//These keys are meant to line up with cert-manager, see
	//https://cert-manager.io/docs/usage/certificate/#additional-certificate-output-formats
	CaCertName     = "ca.crt"
	CertName       = "tls.crt"
	PrivateKeyName = "tls.key"

	//These should be able to be deprecated some time in the future when the new names are fully adopted
	SecretCaCertKey = "ca-cert.pem"
	SecretCertKey   = "public-cert.pem"
	SecretPKKey     = "private-key.pem"
)
