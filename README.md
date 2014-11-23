# Honey

**Honey** is the start of an experimental XMPP server written in Go (read, "it doesn't work yet").

## Features

Honey will support connections via TLS 1, 1.1, and 1.2 only. If you need legacy SSLv3 support (or an unencrypted
socket), go away. For authentication, Honey will support `PLAIN`, `SCRAM-SHA-1`, and `SCRAM-SHA-1-PLUS` if you need
legacy `CRAM-MD5` or `DIGEST-MD5`: also go away. Honey supports `UTF-8` encoded streams. If you need anything else:
You're wrong. Go away.

## License

Honey may be used under the terms of the MIT License or the BSD 2-Clause
License at your discretion. For more information, see the [LICENSE][LICENSE]
file.

[LICENSE]: ./LICENSE.md
