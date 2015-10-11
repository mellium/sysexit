package xmpp

import (
	"bitbucket.org/SamWhited/koine"
	"encoding/xml"
)

type Stanza struct {
	Id      string `xml:"id,attr"`
	Inner   string `xml:",innerxml"`
	Sto     string `xml:"to,attr"`
	Sfrom   string `xml:"from,attr"`
	Body    string `xml:,chardata"`
	XMLName xml.Name
}

func NewStanza(raw string) (*Stanza, error) {
	stanza := new(Stanza)

	if err := xml.Unmarshal([]byte(raw), &stanza); err != nil {
		return stanza, err
	}

	return stanza, nil
}

func (s *Stanza) From() (jid.JID, error) {
	return jid.NewJID(s.Sfrom)
}

func (s *Stanza) To() (jid.JID, error) {
	return jid.NewJID(s.Sto)
}
