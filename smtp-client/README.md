# SMTP email client 

## Test against local SMTP server

### Variant 1: (no TLS) SMTP server

Install MailHog:

```sh
go install github.com/mailhog/MailHog@latest
MailHog # listens on port 1025
```

Run:

```sh
go run . -addr=localhost:1025 -tls=false
Email sent successfully through localhost:1025, using tls: false
```

### Variant 2: (TLS) SMTP server

Install Postfix:

```sh
postconf -d # check if postfix is installed

brew install postfix
sudo vim /etc/postfix/main.cf

#### add this
# TLS Settings
smtpd_use_tls=yes
smtpd_tls_cert_file=/etc/postfix/cert/postfix_cert.pem
smtpd_tls_key_file=/etc/postfix/cert/postfix_key.pem
smtpd_tls_security_level=may

# SMTP Authentication (optional)
smtpd_sasl_auth_enable = yes
smtpd_tls_auth_only = yes
smtpd_recipient_restrictions = permit_sasl_authenticated,permit_mynetworks,reject_unauth_destination

# Additional TLS settings (if needed)
smtp_tls_security_level = may
smtp_tls_note_starttls_offer = yes
smtp_tls_session_cache_database = btree:/var/lib/postfix/smtp_scache
smtpd_tls_session_cache_database = btree:/var/lib/postfix/smtpd_scache
#####
```

Create a certificate and key for TLS

```sh
sudo mkdir /etc/postfix/cert
sudo openssl req -new -x509 -days 365 -nodes -out /etc/postfix/cert/postfix_cert.pem -keyout /etc/postfix/cert/postfix_key.pem
```

Configure Postfix to Use Port 587 for TLS:

```sh
sudo vim /etc/postfix/master.cf

# uncomment this 
submission inet n - n - - smtpd
```


Restart Postfix

```sh
sudo postfix stop
sudo postfix start
```

Run without TLS:

```sh
go run . -addr=localhost:25 -tls=false
Email sent successfully through localhost:25, using tls: false
```

Run with TLS:

```sh
go run . -addr=localhost:587 -tls=true
Email sent successfully through localhost:587, using tls: true
```