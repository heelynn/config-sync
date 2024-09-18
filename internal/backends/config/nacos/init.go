package nacos

import _ "config-sync/internal/properties"

func init() {
	initClients()
	initConfigListener()
}
