package customvalidator

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

// SMTPError ...
type SMTPError struct {
	Err error
}

// Error ...
func (e SMTPError) Error() string {
	return e.Err.Error()
}

// Code ...
func (e SMTPError) Code() string {
	return e.Err.Error()[0:3]
}

// NewSMTPError ...
func NewSMTPError(err error) SMTPError {
	return SMTPError{
		Err: err,
	}
}

var (
	errBadFormat        = errors.New("invalid format")
	errUnresolvableHost = errors.New("unresolvable host")

	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// ValidateFormatMail ...
func ValidateFormatMail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errBadFormat
	}
	return nil
}

// ValidateHostMail ...
func ValidateHostMail(email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return errUnresolvableHost
	}

	client, err := smtp.Dial(fmt.Sprintf("%s:%d", mx[0].Host, 25))
	if err != nil {
		return NewSMTPError(err)
	}
	defer func() {
		err = client.Close()
		if err != nil {
			panic(err)
		}
	}()
	err = client.Hello("checkmail.me")
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Mail("lansome-cowboy@gmail.com")
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Rcpt(email)
	if err != nil {
		return NewSMTPError(err)
	}
	return nil
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
