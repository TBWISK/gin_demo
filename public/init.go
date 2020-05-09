package public

func init() {
	ips := GetLocalIPs()
	if len(ips) > 0 {
		LocalIP = ips[0]
	}
}
