// Copyright (c) 2019 Oasis Labs Inc.  All rights reserved.
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

package curve25519

import "testing"

// Upstream: `test-internals.c`

func TestAdds(t *testing.T) {
	var (
		result [32]byte
		a, b   Bignum25519
	)

	// a = (max_bignum + max_bignum)
	Add(&a, &maxBignum, &maxBignum)

	// b = ((max_bignum + max_bignum) * (max_bignum + max_bignum))
	Mul(&b, &a, &a)
	Contract(result[:], &b)
	if result != maxBignum2SquaredRaw {
		t.Fatalf("maxBignum2SquaredRaw != b")
	}
	Square(&b, &a)
	Contract(result[:], &b)
	if result != maxBignum2SquaredRaw {
		t.Fatalf("maxBignum2SquaredRaw != b (#2)")
	}

	// b = (max_bignum + max_bignum + max_bignum)
	AddAfterBasic(&b, &a, &maxBignum)

	// a = ((max_bignum + max_bignum + max_bignum) * (max_bignum + max_bignum + max_bignum))
	Mul(&a, &b, &b)
	Contract(result[:], &a)
	if result != maxBignum3SquaredRaw {
		t.Fatalf("maxBignum3SquaredRaw != a")
	}
	Square(&a, &b)
	Contract(result[:], &a)
	if result != maxBignum3SquaredRaw {
		t.Fatalf("maxBignum3SquaredRaw != a")
	}
}

func TestSubs(t *testing.T) {
	var (
		result     [32]byte
		zero, a, b Bignum25519
	)

	// a = max_bignum - 0, which expands to 2p + max_bignum - 0
	Sub(&a, &maxBignum, &zero)
	Contract(result[:], &a)
	if result != maxBignumRaw {
		t.Fatalf("maxBignumRaw != a")
	}

	// b = (max_bignum * max_bignum)
	Mul(&b, &a, &a)
	Contract(result[:], &b)
	if result != maxBignumSquaredRaw {
		t.Fatalf("maxBignumSquaredRaw != b")
	}
	Square(&b, &a)
	Contract(result[:], &b)
	if result != maxBignumSquaredRaw {
		t.Fatalf("maxBignumSquaredRaw != b (#2)")
	}

	// b = ((a - 0) - 0)
	SubAfterBasic(&b, &a, &zero)
	Contract(result[:], &b)
	if result != maxBignumRaw {
		t.Fatalf("maxBignumRaw != b")
	}

	// a = (max_bignum * max_bignum)
	Mul(&a, &b, &b)
	Contract(result[:], &a)
	if result != maxBignumSquaredRaw {
		t.Fatalf("maxBignumSquaredRaw != a")
	}
	Square(&a, &b)
	Contract(result[:], &a)
	if result != maxBignumSquaredRaw {
		t.Fatalf("maxBignumSquaredRaw != a (#2)")
	}
}
