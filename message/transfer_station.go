package message

import (
	"github.com/Unknwon/goconfig"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func getNatsConnect() (*nats.Conn, error) {
	cfg, err := goconfig.LoadConfigFile("etc/nats.ini")
	if err != nil {
		logrus.Error("goconfig.LoadConfigFile 「nats.ini」 is error:", err)
		return nil, err
	}
	conn, err := cfg.GetValue("nats", "host")
	if err != nil {
		logrus.Error("cfg.GetValue 「nats.ini」 is error:", err)
		return nil, err
	}
	nc, err := nats.Connect(conn)
	if err != nil {
		return nil, err
	}
	return nc, err
}

//发送消息
func SendMessage(subj string, message []byte) error {
	nc, err := getNatsConnect()
	defer nc.Close()
	if err != nil {
		return err
	}
	if err := nc.Publish(subj, message); err != nil {
		return err
	}
	return nil
}

//TODO 接受消息，应该需要载体
