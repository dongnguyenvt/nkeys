// Copyright 2018 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package nkeys

import (
	"crypto/rand"
	"testing"
)

func TestVerify(t *testing.T) {
	var raw [40]byte
	_, _ = rand.Read(raw[:])
	pub, priv, err := GenerateKey(raw[:])
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("Hello World!")
	sig, err := Sign(priv, data)
	if err != nil {
		t.Fatal(err)
	}
	if !Verify(pub, data, sig) {
		t.Fatal("invalid data")
	}
}
