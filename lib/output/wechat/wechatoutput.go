package wechat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"github.com/elastic/beats/v7/libbeat/publisher"
	workwx "github.com/xen0n/go-workwx/v2"
)

func init() {
	outputs.RegisterType("wechat", NewWeChatOutput)
}

var (
	logger = logp.NewLogger("output.wechat")
)

type WeChatOutput struct {
	App     *workwx.WorkwxApp
	conf    *Config
	encoder codec.Codec
}

// NewWeChatOutput 创建 WeChatOutput 实例
func NewWeChatOutput(_ outputs.IndexManager, beat beat.Info, observer outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	config := &Config{}
	if err := cfg.Unpack(config); err != nil {
		return outputs.Fail(err)
	}

	// logger.Info("config", config)

	encoder, err := codec.CreateEncoder(beat, config.Codec)
	if err != nil {
		return outputs.Fail(err)
	}

	client := workwx.New(config.CorpId)

	app := client.WithApp(config.CorpSecret, config.AgentId)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	out := &WeChatOutput{
		App:     app,
		conf:    config,
		encoder: encoder,
	}

	return outputs.Success(1, 0, out)
}

// Publish 发送日志数据到企业微信
func (w *WeChatOutput) Publish(ctx context.Context, batch publisher.Batch) error {
	events := batch.Events()
	for _, event := range events {
		data, err := w.encoder.Encode("wechat", &event.Content)
		err = w.sendToWeChat(data)
		if err != nil {
			fmt.Println("WeChat send error:", err)
			logger.Error("WeChat send error: %s", err)
			continue
		}
	}
	batch.ACK()
	return nil
}

// sendToWeChat 发送消息到微信 Webhook
func (w *WeChatOutput) sendToWeChat(message []byte) error {

	jsonData, _ := json.Marshal(string(message))

	// send to party(parties)
	to := &workwx.Recipient{
		PartyIDs: []string{w.conf.ToParty},
	}

	err := w.App.SendTextMessage(to, string(jsonData), false)
	if err != nil {
		return err
	}

	return nil
}

// String 实现接口方法
func (w *WeChatOutput) String() string {
	return "WeChat-Output"
}

func (w *WeChatOutput) Close() error {
	return nil
}
