package main

import "flag"

func scan(path string) {
	print(path)
	print("\nscan")
}

func stats(email string) {
	print(email)
	print("\nstats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
