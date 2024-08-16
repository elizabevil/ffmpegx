package optionx

type Protocol struct {
	ProtocolWhitelist string `json:"protocol_whitelist" flag:"-protocol_whitelist"` // "List of protocols that are allowed to be used", OFFSET(protocol_whitelist), AV_OPT_TYPE_STRING, { .str = NULL },  0, 0, D },
}
