// Package logging provides configuration and utility functions for handling
// logging to files, stderr, and stdout. Honey segregates logs using the
// standard syslog severity levels as defined in RFC 5424
// (https://tools.ietf.org/html/rfc5424). Log aggregators for a certain severity
// level also include all lower levels.
//
// The severity levels are:
//
//   * Emerg
//   * Alert
//   * Crit
//   * Err
//   * Warn
//   * Notice
//   * Info
//   * Debug
//
// More information on the severity levels (and related constants) can be found
// in the log/syslog package.
package logging
