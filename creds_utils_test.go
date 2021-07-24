package nkeys

import (
	"bytes"
	"testing"
)

func Test_ParseDecoratedJWTBad(t *testing.T) {
	v, err := ParseDecoratedJWT([]byte("foo"))
	if err != nil {
		t.Fatal(err)
	}
	if v != "foo" {
		t.Fatal("unexpected input was not returned")
	}
}

func Test_ParseDecoratedSeedBad(t *testing.T) {
	if _, err := ParseDecoratedNKey([]byte("foo")); err == nil {
		t.Fatal("Expected error")
	} else if err.Error() != "no nkey seed found" {
		t.Fatal(err)
	}
}

const (
	credsSeed      = `5355d473a21046e10e427bb3ac4aecba07b837fc8443e3b40b3f616e612bf4da978d4b7643f170ffecc21a9d`
	credsJwt       = `eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiJYVUlMUVI3T1JMREc2RFNBQzU3RzRHR05IQUhDUlZXQk1RWExQQ0kyQ0ZISUU3Uk1ZVzNRIiwiaWF0IjoxNjI3MTEwNTE2LCJpc3MiOiI0MTAzMmI2MWFkNWNiNzdmNzEyMjdhYzQzYmYwNjVjMTQyOWYwMTI0OTA5YmUwMTNhZjVlOGJiZTlkNWJmZGU0ODQ1NDE0ZjYiLCJzdWIiOiI1NTAyZmVhNjdiMzU4ZjM2MmU0OTIxYjcwOWQzMDcxYjBjMDZjMmVhZjhlODIxZGFkYzEzMzE4OTU5ZGYwYzIwZWFmYjg4OWMiLCJ0eXBlIjoidXNlciIsIm5hdHMiOnsicHViIjp7fSwic3ViIjp7fX19.19y0_YzJJyJQB90lJWm0_Wu9zpuf8zuSxR__19Khv1t_TwG9wA2euXSfNISi7OxNmZiuZPN3CWzTkyeUPNCUyQE`
	decoratedCreds = `-----BEGIN NATS USER JWT-----
eyJ0eXAiOiJqd3QiLCJhbGciOiJlZDI1NTE5In0.eyJqdGkiOiJYVUlMUVI3T1JMREc2RFNBQzU3RzRHR05IQUhDUlZXQk1RWExQQ0kyQ0ZISUU3Uk1ZVzNRIiwiaWF0IjoxNjI3MTEwNTE2LCJpc3MiOiI0MTAzMmI2MWFkNWNiNzdmNzEyMjdhYzQzYmYwNjVjMTQyOWYwMTI0OTA5YmUwMTNhZjVlOGJiZTlkNWJmZGU0ODQ1NDE0ZjYiLCJzdWIiOiI1NTAyZmVhNjdiMzU4ZjM2MmU0OTIxYjcwOWQzMDcxYjBjMDZjMmVhZjhlODIxZGFkYzEzMzE4OTU5ZGYwYzIwZWFmYjg4OWMiLCJ0eXBlIjoidXNlciIsIm5hdHMiOnsicHViIjp7fSwic3ViIjp7fX19.19y0_YzJJyJQB90lJWm0_Wu9zpuf8zuSxR__19Khv1t_TwG9wA2euXSfNISi7OxNmZiuZPN3CWzTkyeUPNCUyQE
------END NATS USER JWT------

************************* IMPORTANT *************************
NKEY Seed printed below can be used to sign and prove identity.
NKEYs are sensitive and should be treated as secrets.

-----BEGIN USER NKEY SEED-----
5355d473a21046e10e427bb3ac4aecba07b837fc8443e3b40b3f616e612bf4da978d4b7643f170ffecc21a9d
------END USER NKEY SEED------

*************************************************************
`
)

func Test_ParseDecoratedSeedAndJWT(t *testing.T) {
	// test with and without \r\n
	for _, creds := range [][]byte{[]byte(decoratedCreds),
		bytes.ReplaceAll([]byte(decoratedCreds), []byte{'\n'}, []byte{'\r', '\n'})} {
		kp, err := ParseDecoratedUserNKey(creds)
		if err != nil {
			t.Fatal(err)
		}
		pu, err := kp.Seed()
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(pu, []byte(credsSeed)) {
			t.Fatal("seeds don't match")
		}

		kp, err = ParseDecoratedNKey(creds)
		if err != nil {
			t.Fatal(err)
		}
		pu, err = kp.Seed()
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(pu, []byte(credsSeed)) {
			t.Fatal("seeds don't match")
		}

		jwt, err := ParseDecoratedJWT(creds)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal([]byte(jwt), []byte(credsJwt)) {
			t.Fatal("jwt don't match")
		}
	}
}
