// Copyright 2020 The jackal Authors
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

package coremodel

import (
	"fmt"

	"github.com/jackal-xmpp/stravaganza"
	"github.com/jackal-xmpp/stravaganza/jid"
	"github.com/ortuman/jackal/version"
)

// ClusterMember represents a cluster instance address and port.
type ClusterMember struct {
	InstanceID string
	Host       string
	Port       int
	APIVer     *version.SemanticVersion
}

// String returns Member string representation.
func (m *ClusterMember) String() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}

// User represents a user entity.
type User struct {
	Username string
	Scram    struct {
		SHA1           string
		SHA256         string
		SHA512         string
		SHA3512        string
		Salt           string
		IterationCount int
		PepperID       string
	}
}

// Resource represents a resource entity.
type Resource struct {
	InstanceID string
	JID        *jid.JID
	Presence   *stravaganza.Presence
	Context    map[string]string
}

// IsAvailable returns presence available value.
func (r *Resource) IsAvailable() bool {
	if r.Presence != nil {
		return r.Presence.IsAvailable()
	}
	return false
}

// Priority returns resource presence priority.
func (r *Resource) Priority() int8 {
	if r.Presence != nil {
		return r.Presence.Priority()
	}
	return 0
}

// Value returns resource context value associated to cKey.
func (r *Resource) Value(cKey string) string {
	return r.Context[cKey]
}