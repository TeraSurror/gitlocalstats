package main

import (
	"flag"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for git repositories")
	// Feel free to reach to me at this email!
	flag.StringVar(&email, "email", "harshshelar22@gmail.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		// Scans given folder, picks the ones which have a git repo and adds it to the .gitlocalstats in the home directory
		scan(folder)
		return
	}

	// Outputs the stats
	stats(email)
}
