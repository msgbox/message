package messages

import (
	"code.google.com/p/goprotobuf/proto"
	"regexp"
)

// Parses a Message from the queue

func Parse(data []byte) *Message {
	msg := &Message{}
	proto.Unmarshal(data, msg)
	return msg
}

// Need Error Handling here if no matches
func ParseReceiver(s string) (string, string) {
	reg, _ := regexp.Compile("^([A-Za-z0-9_-]+):([A-Za-z0-9_-]+)")
	re := reg.FindAllStringSubmatch(s, 1)

	return re[0][1], re[0][2]
}
