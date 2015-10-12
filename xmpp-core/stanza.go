package xmpp

import (
	"bitbucket.org/SamWhited/koine"
	"encoding/xml"
)

type stanza struct {
	Id      string `xml:"id,attr"`
	Inner   string `xml:",innerxml"`
	Sto     string `xml:"to,attr"`
	Sfrom   string `xml:"from,attr"`
	Lang    string `xml:"xml:lang,attr"`
	XMLName xml.Name
}

func (s *stanza) From() (jid.JID, error) {
	return jid.NewJID(s.Sfrom)
}

func (s *stanza) To() (jid.JID, error) {
	return jid.NewJID(s.Sto)
}
