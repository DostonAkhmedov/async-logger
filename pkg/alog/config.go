package alog

type Config struct {
	MsgChanBufferLength int    `yaml:"msg_chan_buffer_length"`
	ErrChanBufferLength int    `yaml:"err_chan_buffer_length"`
	Out                 string `yaml:"out"`
}

func NewConfig(
	msgChanBufferLength int,
	errChanBufferLength int,
	out string,
) *Config {
	return &Config{
		MsgChanBufferLength: msgChanBufferLength,
		ErrChanBufferLength: errChanBufferLength,
		Out:                 out,
	}
}
