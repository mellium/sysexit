// Copyright 2017 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

// Package sysexit lists standard return codes that are often used by
// applications in Unix like environments.
//
// For more info on these errors see the man page for SYSEXITS(3) on a BSD.
package sysexit // import "mellium.im/sysexit"

// Exit codes defined by SYSEXITS(3).
const (
	Ok             int = 0  // successful termination
	ErrBase        int = 64 // base value for error messages
	ErrUsage       int = 64 // command line usage error
	ErrData        int = 65 // data format error
	ErrNoInput     int = 66 // cannot open input
	ErrNoUser      int = 67 // addressee unknown
	ErrNoHost      int = 68 // host name unknown
	ErrUnavailable int = 69 // service unavailable
	ErrSoftware    int = 70 // internal software error
	ErrOS          int = 71 // system error (e.g., can't fork)
	ErrOSFile      int = 72 // critical OS file missing
	ErrCantCreat   int = 73 // can't create (user) output file
	ErrIO          int = 74 // input/output error
	ErrTempFail    int = 75 // temp failure; user is invited to retry
	ErrProtocol    int = 76 // remote error in protocol
	ErrNoPerm      int = 77 // permission denied
	ErrConfig      int = 78 // configuration error
)
