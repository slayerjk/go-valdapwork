package main

import (
	"fmt"
	"os"

	ldap "github.com/slayerjk/go-valdapwork"
)

func main() {
	fqdn := "bank.corp.centercredit.kz"
	baseDN := "dc=bank,dc=corp,dc=centercredit,dc=kz"
	acc := "marchenm"
	bindUser := "marchenm@" + fqdn
	bindPassword := "sU9Xak:tgN"

	// make LDAP connection
	conn, err := ldap.MakeLdapConnection(fqdn)
	if err != nil {
		fmt.Printf("failed to make LDAP connection:\n\t%v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// start TLS
	err = ldap.StartTLSConnWoVerification(conn)
	if err != nil {
		fmt.Printf("failed to make TLS connection to LDAP:\n\t%v\n", err)
	}

	// make LDAP bind
	err = ldap.LdapBind(conn, bindUser, bindPassword)
	if err != nil {
		fmt.Printf("failed to make LDAP bind:\n\t%v\n", err)
		os.Exit(1)
	}

	// make LDAP search request
	filter := fmt.Sprintf("(&(objectClass=user)(samaccountname=%s))", acc)
	searchResult, err := ldap.MakeSearchReq(conn, baseDN, filter, "displayname", "employeeID")
	if err != nil {
		fmt.Printf("failed to make LDAP search request:\n\t%v\n", err)
		os.Exit(1)
	}

	// (OPTIONAL) check if single result
	if len(searchResult) > 1 {
		fmt.Printf("multiple result of LDAP search is mistake in request")
	}

	// print result(displayName of user account)
	fmt.Println(searchResult[0].GetAttributeValue("displayName"))
	fmt.Println(searchResult[0].GetAttributeValue("employeeID"))

}
