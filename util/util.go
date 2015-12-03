package util

import (
	"crypto/rand"
	"encoding/hex"

	"../config"

	"bitbucket.org/SamWhited/koine"
)

func Id() (string, error) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	return hex.EncodeToString(b), err
}

func ServesHost(host jid.JID, C *config) (served bool) {
	// TODO: Write a more efficient way to do this.
	served = false
	for h := range C.Hosts {
		jid, err := jid.NewJID(h.Name)
		if err != nil {
			continue
		}
		if stream.To().Equals(jid) {
			served = true
			break
		}
	}
	return
}
