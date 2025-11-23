package communicate

func MakeConsumerCfg(key string, handler MessageHandler) InHandlerSubCfg {
	return InHandlerSubCfg{
		StreamName:      key,
		Subjects:        []string{key},
		ConsumerDurable: key,
		ConsumerName:    key,
		Consume:         handler,
	}
}
