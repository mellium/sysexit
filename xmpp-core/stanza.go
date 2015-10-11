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
	Body    string `xml:,chardata,ommitempty"`
	Lang    string `xml:"xml:lang,attr,ommitempty"`
	XMLName xml.Name
}

type Message struct {
	stanza
}

type Iq struct {
	stanza
}

type Presence struct {
	stanza
}

func NewStanza(raw string) (*stanza, error) {
	s := new(stanza)

	if err := xml.Unmarshal([]byte(raw), &s); err != nil {
		return s, err
	}

	return s, nil
}

func (s *stanza) From() (jid.JID, error) {
	return jid.NewJID(s.Sfrom)
}

func (s *stanza) To() (jid.JID, error) {
	return jid.NewJID(s.Sto)
}
