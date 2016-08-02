#
## need
`ldapuser.ini` is needed for testing. it looks like
> user=ldapuser
> password=ldappass

## local testing
GOPATH="$GOPATH:$HOME/src/dbexperements/go/" go test -v ldapauth_test.go
