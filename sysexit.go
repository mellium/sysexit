// Copyright 2017 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

// Package sysexit lists standard return codes that are often used by
// applications in Unix like environments.
//
// For more info on these errors see the man page for SYSEXITS(3) on a BSD.
package sysexit

// Code is an error that corresponds to an exit code.
type Code int

// Error satisfies the error interface for Code
func (Code) Error() string { return "" }

// Exit codes defined by SYSEXITS(3).
const (
	Ok             int  = 0  // successful termination
	ErrBase        Code = 64 // base value for error messages
	ErrUsage       Code = 64 // command line usage error
	ErrData        Code = 65 // data format error
	ErrNoInput     Code = 66 // cannot open input
	ErrNoUser      Code = 67 // addressee unknown
	ErrNoHost      Code = 68 // host name unknown
	ErrUnavailable Code = 69 // service unavailable
	ErrSoftware    Code = 70 // internal software error
	ErrOS          Code = 71 // system error (e.g., can't fork)
	ErrOSFile      Code = 72 // critical OS file missing
	ErrCantCreat   Code = 73 // can't create (user) output file
	ErrIO          Code = 74 // input/output error
	ErrTempFail    Code = 75 // temp failure; user is invited to retry
	ErrProtocol    Code = 76 // remote error in protocol
	ErrNoPerm      Code = 77 // permission denied
	ErrConfig      Code = 78 // configuration error
)
