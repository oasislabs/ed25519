// Copyright (c) 2020 Oasis Labs Inc.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Oasis Labs Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package ed25519

// Validate the implementation behavior against the tests presented in
// the paper "Taming the many EdDSAs" by Chalkias, Garillot, and
// Nikolaenko.
//
// Test data taken at commit 336651ba7f1c1ae90b7deac7d175290863a00b66 from
// https://github.com/novifinancial/ed25519-speccheck/blob/master/scripts/cases.json

import (
	"compress/gzip"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

var speccheckExpectedResults = []bool{
	false, // 0: small order A, small order R
	false, // 1: small order A, mixed order R
	true,  // 2: mixed order A, small order R
	true,  // 3: mixed order A, mixed order R
	true,  // 4: cofactored verify
	true,  // 5: cofactored verify computes 8(hA) instead of (8h mod L)A
	false, // 6: non-canonical S (S > L)
	false, // 7: non-canonical S (S >> L)
	false, // 8: mixed order A, non-canonical small order R (accepted if R reduced before hashing)
	true,  // 9: mixed order A, non-canonical small order R (accepted if R not reduced before hashing)
	false, // 10: non-canonical small order A, mixed order R (accepted if A reduced before hashing)
	false, // 11: non-canonical small order A, mixed order R (accepted if A not reduced before hashing)
}

var speccheckExpectedResultsZIP215 = []bool{
	true,  // 0: small order A, small order R
	true,  // 1: small order A, mixed order R
	true,  // 2: mixed order A, small order R
	true,  // 3: mixed order A, mixed order R
	true,  // 4: cofactored verify
	true,  // 5: cofactored verify computes 8(hA) instead of (8h mod L)A
	false, // 6: non-canonical S (S > L)
	false, // 7: non-canonical S (S >> L)
	false, // 8: mixed order A, non-canonical small order R (accepted if R reduced before hashing)
	true,  // 9: mixed order A, non-canonical small order R (accepted if R not reduced before hashing)
	true,  // 10: non-canonical small order A, mixed order R (accepted if A reduced before hashing)
	true,  // 11: non-canonical small order A, mixed order R (accepted if A not reduced before hashing)
}

type speccheckTestVector struct {
	Message   string `json:"message"`
	PublicKey string `json:"pub_key"`
	Signature string `json:"signature"`
}

func (v *speccheckTestVector) toComponents() ([]byte, PublicKey, []byte, error) {
	var pk PublicKey

	msg, err := hex.DecodeString(v.Message)
	if err != nil {
		return nil, pk, nil, fmt.Errorf("failed to decode message: %w", err)
	}
	rawPk, err := hex.DecodeString(v.PublicKey)
	if err != nil {
		return nil, pk, nil, fmt.Errorf("failed to decode public key: %w", err)
	}
	sig, err := hex.DecodeString(v.Signature)
	if err != nil {
		return nil, pk, nil, fmt.Errorf("failed to decode signature: %w", err)
	}
	if len(rawPk) != PublicKeySize {
		return nil, pk, nil, fmt.Errorf("invalid public key size")
	}
	if len(sig) != SignatureSize {
		return nil, pk, nil, fmt.Errorf("invalid signature size")
	}

	pk = PublicKey(rawPk)

	return msg, pk, sig, nil
}

func (v *speccheckTestVector) Run(t *testing.T, isBatch, isZIP215 bool) bool {
	msg, pk, sig, err := v.toComponents()
	if err != nil {
		t.Fatal(err)
	}

	opts := &Options{
		ZIP215Verify: isZIP215,
	}

	var sigOk bool
	switch isBatch {
	case false:
		sigOk = VerifyWithOptions(pk, msg, sig, opts)
	case true:
		var pks []PublicKey
		var sigs, msgs [][]byte
		for i := 0; i < minBatchSize*2; i++ {
			pks = append(pks, pk)
			msgs = append(msgs, msg)
			sigs = append(sigs, sig)
		}

		var valid []bool
		sigOk, valid, err = VerifyBatch(rand.Reader, pks, msgs, sigs, opts)
		if err != nil {
			t.Fatal(err)
		}
		for i, v := range valid {
			if v != sigOk {
				t.Fatalf("sigOk != valid[%d]", i)
			}
		}
	}

	return sigOk
}

func TestSpeccheck(t *testing.T) {
	f, err := os.Open("testdata/speccheck_cases.json.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	rd, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	defer rd.Close()

	var testVectors []speccheckTestVector

	dec := json.NewDecoder(rd)
	if err = dec.Decode(&testVectors); err != nil {
		t.Fatal(err)
	}

	for idx, tc := range testVectors {
		n := fmt.Sprintf("TestCase_%d", idx)
		expected := speccheckExpectedResults[idx]
		t.Run(n, func(t *testing.T) {
			if sigOk := tc.Run(t, false, false); sigOk != expected {
				t.Fatalf("behavior mismatch: %v (expected %v)", sigOk, expected)
			}
		})
		t.Run(n+"_Batch", func(t *testing.T) {
			if sigOk := tc.Run(t, true, false); sigOk != expected {
				t.Fatalf("behavior mismatch: %v (expected %v)", sigOk, expected)
			}
		})

		expected = speccheckExpectedResultsZIP215[idx]
		n = n + "_ZIP215"
		t.Run(n, func(t *testing.T) {
			if sigOk := tc.Run(t, false, true); sigOk != expected {
				t.Fatalf("behavior mismatch: %v (expected %v)", sigOk, expected)
			}
		})
		t.Run(n+"_Batch", func(t *testing.T) {
			if sigOk := tc.Run(t, true, true); sigOk != expected {
				t.Fatalf("behavior mismatch: %v (expected %v)", sigOk, expected)
			}
		})
	}
}
