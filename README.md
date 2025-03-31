# go-ldap
Go - va helper functions for "go-ldap/ldap" package

<h2>Functions</h2>

* MakeLdapConnection - create *ldap.Conn
* StartTLSConnWoVerification - create *ldap.Conn with TLS
* LdapBind - make LDAP bind
* MakeSearchReq - get *ldap.Entry result using filter and attributes
* GetAttr - get single string result of *ldap.Entry for single attribute
* SearchEnabledSAMAByDisplayname - Search ONLY Enabled user's 'samaccountname' by it's 'displayname'

Need LDAP data, check below json file example.

<h3>LDAP Data json file example</h3>

```
{
    "ldap-bind-user": "<LDAP BIND USER>",
    "ldap-bind-pass": "<LDAP BIND USER'S PASS",
    "ldap-fqdn": "<DOMAIN FQDN>",
    "ldap-basedn": "DC=DOMAIN,DC=EXAMPLE,DC=COM",
}
```


