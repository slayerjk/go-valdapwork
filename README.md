# go-ldap
Go - va helper functions for "go-ldap/ldap" package

<h2>Functions</h2>

* MakeLdapConnection - create *ldap.Conn
* StartTLSConnWoVerification - create *ldap.Conn with TLS
* LdapBind - make LDAP bind
* MakeSearchReq - get *ldap.Entry result using filter and attributes
* GetAttr - get single string result of *ldap.Entry for single attribute
* SearchEnabledSAMAByDisplayname - Search ONLY Enabled user's 'samaccountname' by it's 'displayname'
* ConvertPwdLastSetAttr - Convert LDAP attribute time 'pwdLastSet'(100-nanosecond steps since 12:00 AM, January 1, 1601, UTC) to time.Time
