package wechat

import "github.com/elastic/beats/v7/libbeat/outputs/codec"

type Config struct {
	CorpId     string       `config:"corp_id"`
	CorpSecret string       `config:"corp_secret"`
	AgentId    int64        `config:"agent_id"`
	ToParty    string       `config:"to_party"`
	Codec      codec.Config `config:"codec"`
}
