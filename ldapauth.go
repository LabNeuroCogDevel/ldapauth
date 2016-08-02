package ldapauth

import (
    "fmt"
//    "log"
    "strings"
    "gopkg.in/ldap.v2"
)

func AuthLdap(user string,password string) error{
   var ldap_server string = "acct.upmchs.net"
   //var domain      string = "1UPMC-ACCT"
   var ldap_port   uint16 = 389

   var duser string = strings.Join([]string{"1UPMC-ACCT", user},"\\")
   //var uname string = "1UPMC-ACCT\\foranw"

   l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldap_server, ldap_port))
   //err = l.Bind( uname,password)
   err = l.Bind(duser,password)
   return(err)

}
func IsAuth(user string, password string) bool {
  err:=AuthLdap(user,password)
  if(err != nil ){
    return false
    //log.Fatal(err)
  } else {
   return true
  }
}
