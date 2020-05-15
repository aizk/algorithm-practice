package main

import "strings"

func main() {
	numUniqueEmails(nil)
}

func numUniqueEmails(emails []string) int {
	rsp := map[string]bool{}

	for _, email := range emails {
		eml := strings.Split(email, "@")
		local := eml[0]
		domain := eml[1]
		// merge .
		local = strings.Replace(local, ".", "", -1)
		// remove +...
		idx := strings.Index(local, "+")
		if idx != -1 {
			local = local[:idx]
		}
		rsp[local+domain] = true
	}
	return len(rsp)
}
