package main

import (
  "fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit"
)

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

// Rand SSHDErrorLog
// Invalid user ts1 from 46.249.99.166 port 41064
// pam_unix(sshd:auth): check pass; user unknown
// pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=46.249.99.166
// Failed password for invalid user ts1 from 46.249.99.166 port 41064 ssh2
// Received disconnect from 46.249.99.166 port 41064:11: Bye Bye [preauth]
// Disconnected from invalid user ts1 46.249.99.166 port 41064 [preauth]
func RandSSHDErrorLog() string {
	versions := []string{
    fmt.Sprintf("Invalid user %s from %s port %d", gofakeit.Username(), gofakeit.IPv4Address(), gofakeit.Number(1, 10000)),
    "pam_unix(sshd:auth): check pass; user unknown",
    fmt.Sprintf("pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=%s", gofakeit.IPv4Address()),
    fmt.Sprintf("Failed password for invalid user %s from %s port %d ssh2", gofakeit.Username(), gofakeit.IPv4Address(), gofakeit.Number(1, 10000)),
    fmt.Sprintf("Received disconnect from %s port %d:11: Bye Bye [preauth]", gofakeit.IPv4Address(), gofakeit.Number(1, 10000)),
    fmt.Sprintf("Disconnected from invalid user %s %s port %d [preauth]", gofakeit.Username(), gofakeit.IPv4Address(), gofakeit.Number(1, 10000)),
  }
	return versions[rand.Intn(6)]
}

// Rand SSHDSuccessLog
// Accepted publickey for pockost from 51.15.25.173 port 53170 ssh2: RSA SHA256:FlVVLnPYxczG2CQCvnT/ewkzR5371Ab0B2AKoJIQq9E
// pam_unix(sshd:session): session opened for user pockost(uid=1000) by (uid=0)
// New session 12 of user pockost.
// Started session-12.scope - Session 12 of User pockost.
// pam_env(sshd:session): deprecated reading of user environment enabled
// Received disconnect from 51.15.25.173 port 53170:11: disconnected by user
// Disconnected from user pockost 51.15.25.173 port 53170
// session-12.scope: Deactivated successfully.
// pam_unix(sshd:session): session closed for user pockost
// Session 12 logged out. Waiting for processes to exit.
// Removed session 12.

func RandSSHDSuccessLog() string {
	versions := []string{
    fmt.Sprintf("pam_unix(sshd:session): session opened for user %s(uid=%d) by (uid=0)", gofakeit.Username(), gofakeit.Number(1000, 2000)),
  }
	return versions[rand.Intn(1)]
}
