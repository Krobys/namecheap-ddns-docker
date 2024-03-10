package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Namecheap Dynamic DNS client Version", version)
	fmt.Println("Git Repo:", gitrepo)

	domain := flag.String("domain", "", "Domain name e.g. example.com")
	host := flag.String("host", "", "Subdomain or hostname e.g. www")
	password := flag.String("password", "", "Dynamic DNS Password from Namecheap")

	flag.Parse()
	if *domain == "" || *host == "" || *password == "" {
		fmt.Println("ERROR domain, host and Dynamic DDNS password are mandatory")
		fmt.Printf("\nUsage of %s:\n", os.Args[1])
		flag.PrintDefaults()
		os.Exit(1)
	}

	pubIp, err := getPubIP()
	if err != nil {
		DDNSLogger(ErrorLog, *host, *domain, err.Error())
	} else {
		if err = setDNSRecord(*host, *domain, *password, pubIp); err != nil {
			DDNSLogger(ErrorLog, *host, *domain, err.Error())
			DDNSLogger(WarningLog, *host, *domain, "Ignoring above error. If this is not right, Re-run the process after fixing the error")
		} else {
			DDNSLogger(InformationLog, *host, *domain, "Record updated. "+pubIp)
		}
	}

	updateRecord(*domain, *host, *password)
}
