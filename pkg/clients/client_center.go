package clients

func Client(name, typ string, param map[interface{}]interface{}) BaseClientInterface {
	var res BaseClientInterface
	if res = checkoutClient(name); res != nil {
		return res
	}

	switch typ {
	case "StdoutClient":
		res = (&StdoutClient{}).InitP(param)
	case InfluxDBV1ClientType:
		res = (&InfluxdbV1Client{}).InitP(param)
	}

	if res == nil {
		return nil
	}

	setClient(name, res)
	return res
}

func checkoutClient(name string) BaseClientInterface {
	return clientCache[name]
}
func setClient(name string, client BaseClientInterface) {
	clientCache[name] = client
}

var clientCache = map[string]BaseClientInterface{}
