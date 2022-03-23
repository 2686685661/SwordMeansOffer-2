/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/18 12:47 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _1_功能选项模式_options_


type Message struct {
	id int
	name string
	address string
	phone int
}

type Option func( msg *Message)

var defaultMessage = Message{
	id:      -1,
	name:    "-1",
	address: "-1",
	phone:   -1,
}

func WithId(id int) Option {
	return func(msg *Message) {
		msg.id = id
	}
}

func WithName(name string) Option {
	return func(msg *Message) {
		msg.name = name
	}
}

func WithAddress(address string) Option {
	return func(msg *Message) {
		msg.address = address
	}
}

func WithPhone(phone int) Option {
	return func(msg *Message) {
		msg.phone = phone
	}
}


func NewByOption(opts ...Option) Message {
	msg := defaultMessage
	for _, o := range opts {
		o(&msg)
	}
	return msg
}

func NewRequireIdByOption(id int, opts ...Option) Message {
	msg := defaultMessage
	msg.id = id
	for _, o := range opts {
		o(&msg)
	}
	return msg
}