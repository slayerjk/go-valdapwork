package main

import (
	"fmt"
	"os"

	ldap "github.com/slayerjk/go-valdapwork"
)

func main() {
	fqdn := "dom.example.com"
	baseDN := "dc=dom,dc=example,dc=com"
	acc := "domain-account"
	// bind user format must be "acc@dom.example.com"
	bindUser := acc + "@" + fqdn
	bindPassword := "domain-account-password"

	// start LDAP connection over TLS
	conn, err := ldap.StartTLSConnWoVerification(fqdn)
	if err != nil {
		fmt.Printf("failed to make TLS connection to LDAP:\n\t%v\n", err)
	}

	// make LDAP bind
	err = ldap.LdapBind(conn, bindUser, bindPassword)
	if err != nil {
		fmt.Printf("failed to make LDAP bind:\n\t%v\n", err)
		os.Exit(1)
	}

	// make LDAP search request to get several attributes
	filter := fmt.Sprintf("(&(objectClass=user)(samaccountname=%s))", acc)
	searchResult, err := ldap.MakeSearchReq(conn, baseDN, filter, "displayname", "employeeID")
	if err != nil {
		fmt.Printf("failed to make LDAP search request:\n\t%v\n", err)
		os.Exit(1)
	}
	// print result(displayName of user account)
	fmt.Println(searchResult[0].GetAttributeValue("displayName"))
	fmt.Println(searchResult[0].GetAttributeValue("employeeID"))

	// get single result of MakeSearchReq base on LDAP attribute name
	cn, err := ldap.GetAttr(conn, filter, acc, baseDN, "cn")
	if err != nil {
		fmt.Printf("failed to get CN LDAP attribute for %s", acc)
	}
	fmt.Println(cn)

	// close LDAP connection
	conn.Close()
}
