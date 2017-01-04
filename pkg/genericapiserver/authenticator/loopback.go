/*
Copyright 2016 The Kubernetes Authors.

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

package authenticator

import (
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/kubernetes/pkg/auth/authenticator"
	"k8s.io/kubernetes/pkg/auth/authenticator/bearertoken"
	"k8s.io/kubernetes/plugin/pkg/auth/authenticator/token/tokenfile"
)

// newAuthenticatorFromToken returns an authenticator.Request or an error
func NewAuthenticatorFromTokens(tokens map[string]*user.DefaultInfo) authenticator.Request {
	return bearertoken.New(tokenfile.New(tokens))
}
