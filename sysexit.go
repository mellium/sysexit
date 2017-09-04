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
	OK          Code = 0  // successful termination
	BASE        Code = 64 // base value for error messages
	USAGE       Code = 64 // command line usage error
	DATAERR     Code = 65 // data format error
	NOINPUT     Code = 66 // cannot open input
	NOUSER      Code = 67 // addressee unknown
	NOHOST      Code = 68 // host name unknown
	UNAVAILABLE Code = 69 // service unavailable
	SOFTWARE    Code = 70 // internal software error
	OSERR       Code = 71 // system error (e.g., can't fork)
	OSFILE      Code = 72 // critical OS file missing
	CANTCREAT   Code = 73 // can't create (user) output file
	IOERR       Code = 74 // input/output error
	TEMPFAIL    Code = 75 // temp failure; user is invited to retry
	PROTOCOL    Code = 76 // remote error in protocol
	NOPERM      Code = 77 // permission denied
	CONFIG      Code = 78 // configuration error
)
