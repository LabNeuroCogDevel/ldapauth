package ldapauth

import (
  "testing"
  "ldapauth"
  "gopkg.in/ini.v1"
)

func TestLDAPAuth(t *testing.T){
  cfg, err := ini.Load("ldapuser.ini")
  if err != nil{
   t.Skip("Need ldapuser.ini file with user= and password=")
  }

  user,uerr := cfg.Section("").GetKey("user")
  pass,perr := cfg.Section("").GetKey("password")
  if perr !=nil || uerr != nil { t.Skip("Need ldapuser.ini file with user= and password=") }

  isauth := ldapauth.IsAuth(user.String(),pass.String())
  if ! isauth {
    t.Error("Expeciting auth for foranw, not auth'ed!", isauth)
  }
}

func TestLDAPAuthFail(t *testing.T){

  isauth := ldapauth.IsAuth("doesnotexist","willnotautheticate")
  if isauth {
    t.Error("Expeciting auth for nonsense user to fail", isauth)
  }
}
