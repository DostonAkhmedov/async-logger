package alog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type log struct {
	msgChan   chan string
	errorChan chan error
	done      chan bool
	writer    io.Writer
}

func New(cfg *Config) *log {
	var (
		w   io.Writer
		err error
	)

	if cfg.Out == "" {
		w = os.Stdout
	} else {
		w, err = os.Create(cfg.Out)
		if err != nil {
			fmt.Printf("unable to open log file error: %s\n", err.Error())
		}
	}

	return &log{
		msgChan:   make(chan string, cfg.MsgChanBufferLength),
		errorChan: make(chan error, cfg.ErrChanBufferLength),
		done:      make(chan bool),
		writer:    w,
	}
}

func (l *log) Start() {
	go func() {
		for {
			select {
			case msg, ok := <-l.msgChan:
				if !ok {
					l.done <- true
					return
				}
				l.write(msg)
			case err, ok := <-l.errorChan:
				if !ok {
					l.done <- true
					return
				}
				l.write(err.Error())
			}
		}
	}()
}

func (l *log) Stop() {
	close(l.msgChan)
	close(l.errorChan)
	<-l.done
}

func (l *log) Info(msg string) {
	l.msgChan <- msg
}

func (l *log) Fatal(err error) {
	l.errorChan <- err
}

func (l *log) write(msg string) {
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	now := time.Now().Format("2006-02-01 15:04:05")
	_, err := fmt.Fprintf(l.writer, "[%s] - %s", now, msg)
	if err != nil {
		fmt.Printf("alog write msg error: %s\n", err.Error())
	}
}
