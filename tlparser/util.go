package tlparser

import "strings"

func replaceKeyWords(input string) string {
	input = strings.Replace(input, "Api", "API", -1)
	input = strings.Replace(input, "Url", "URL", -1)
	input = strings.Replace(input, "Id", "ID", -1)
	input = strings.Replace(input, "Ttl", "TTL", -1)
	input = strings.Replace(input, "Html", "HTML", -1)
	input = strings.Replace(input, "Uri", "URI", -1)
	input = strings.Replace(input, "Ip", "IP", -1)
	input = strings.Replace(input, "Udp", "UDP", -1)

	return input
}
