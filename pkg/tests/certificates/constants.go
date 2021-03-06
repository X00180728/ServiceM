// Package certificates defines sample certificates used for unit testing.
package certificates

// SampleCertificatePEM is a sample PEM cert used for tests
// Source: https://golang.org/src/crypto/x509/example_test.go
const SampleCertificatePEM = `-----BEGIN CERTIFICATE-----
MIIDLjCCAhYCCQCBwIhgLWe6HjANBgkqhkiG9w0BAQUFADBZMQswCQYDVQQGEwJV
UzEQMA4GA1UECgwHQXogTWVzaDE4MDYGA1UEAwwvNjNkMDQ0YzktNzdjNy00MmFl
LWFmZGMtNjM2YTFiNmFiNGUyLmF6dXJlLm1lc2gwHhcNMjIwMTA1MjIxMjE4WhcN
MjIwMjA0MjIxMjE4WjBZMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHQXogTWVzaDE4
MDYGA1UEAwwvNjNkMDQ0YzktNzdjNy00MmFlLWFmZGMtNjM2YTFiNmFiNGUyLmF6
dXJlLm1lc2gwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDD3+gqR5tL
q3w2KZOVCJRaQ2+0bdDmqvWf4YZjsYlIWUMSxQNhX9fm6u/X/fUbwVMpDP3t2A7A
rgJPiakti8676Ws7utVbYi2PvjLfcVtsM0UBtAqXfHN2Rg+Ne7B9AanepUeJIfzs
+/jr6MAhuhTZA/RquhLbRGJKrmHsgnGuAyGn581TXiL52HUvbJ89BbexpcQtUnqF
Uj8JhnHWKTuoNPcLlDMRL5fRX08Zyhzxiyg66ALoZduHNu6HV/Z0YXHlxePKZCIR
rbx58a74q6zYBTWdWqkKhKF1wFYWBwi2ppIPW2U47TOV0IsnWs9o7DsWkFMpf97S
pE7vSyxpPefNAgMBAAEwDQYJKoZIhvcNAQEFBQADggEBABzMgrcHWw05eNvK+XH4
6DRzR3X37N13fKt2kYCPsuosnuti5MRVDVcSvOs4/Z47u9Ade4DOeNeCVTF9TZ4q
qeOdaR7qY/Pw5TUPlQqTpgr6GBQ/JF0hXQ3aAF1kfwVmLZvt4IirISUGluKU1Cwi
BYjefckldL3R1ifckpOC49OBZru2v6YlySQ6+0b7QXnzCMwrbbk/mlaC/Hls7ZN8
Buxx3GKtD73M3CCxeO8zpw9FeX6NSGh5UwsHcORTN7Cv7MMlTlJmUKKTovzzguGj
e+SS+bQ5uuXLb3bomn54Tq8nVFlZJheNRFfSSaqXTIb6SFngVdmMtp1QD/BJ7ke3
v+I=
-----END CERTIFICATE-----`

// SamplePrivateKeyPEM is a sample PEM private key used for tests
const SamplePrivateKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDD3+gqR5tLq3w2
KZOVCJRaQ2+0bdDmqvWf4YZjsYlIWUMSxQNhX9fm6u/X/fUbwVMpDP3t2A7ArgJP
iakti8676Ws7utVbYi2PvjLfcVtsM0UBtAqXfHN2Rg+Ne7B9AanepUeJIfzs+/jr
6MAhuhTZA/RquhLbRGJKrmHsgnGuAyGn581TXiL52HUvbJ89BbexpcQtUnqFUj8J
hnHWKTuoNPcLlDMRL5fRX08Zyhzxiyg66ALoZduHNu6HV/Z0YXHlxePKZCIRrbx5
8a74q6zYBTWdWqkKhKF1wFYWBwi2ppIPW2U47TOV0IsnWs9o7DsWkFMpf97SpE7v
SyxpPefNAgMBAAECggEATsKJp/aDCzo5B85P+W0pueHD2NkPVrEHcvJMB2oruVur
DLELWuwe9EsjhcYn+LETrz36HNjzlaZiZ3kC/b1ps0V4SNwnTkd76oCgFBiQmkFD
ThwG5kK0aqphNpK1tI4mr8/lo8521RO8U5+TIfygxWJBtWh8jI5Ct6TG20LYUw9a
QMhgmEFVXaBRyoIhccuWahJHSwZzlxlmLTj06Gf+Uv9Snhwy7LJe81i9CNWVn8E0
zW+77vUWQ1/AXIyh0fLmQhisHs6d/wbVr9E8GBAyyzN21uzoXNSyWxnwlGk/K1IQ
76KrRVw7zIQ7iqrEsycMtY8uoW8CkRHZOYvtAS5OQQKBgQD4IllwZRbiWFaRXN04
bUgiFjBQjkCMKyPk1b9MryaG4kIgxN9YQRiwwFWueaW4p+HyujT8pAl4xo5RbH37
xKPqgPCQ1XzH9mPo7Mx0OCyv9GaAXlq4FqiJU5T5xF6SoWSgJTKgVPfNtGLAzWaX
l/BRY+19ATAL1kSRXKq7cHpJjwKBgQDKFXZpq5QPXk37CE1hpN6cs8cKkvfU4oaq
V4lC+4TlAah8JjtzXNyAbKtGdV9Q9kgsgDBeaTBY4MZrtnhh6JVY3twGaRBq6pcv
0IleaVVhp7eOwMA4W5AYSnZ6LahFY0YFyzFeEgyzqwbQlFX+A9ovXX+DJlBoM6pn
gcowfqNy4wKBgAVs8tmzTCnM1q+9ARVPxmkAZTQNuDmYY+OIDPPHTKdcYSfIRj3u
xnRu8DCtdkMwYI9nJOt1RsO+S7RaE/MiXJcvFJOGJ4FT0OFx9BKCe++o/2jFJ2Sp
EixWiIZhldPM9Z9O0OmSkgyMajBfDWQ5LUcKUVIPaZaIq90l0pHgprvfAoGBALBc
eMIR3p5m8/FQNpAv3aOuddfxmV5t74675GvTrBBcGRl4GEw+z6U4sWVFS9ERjr1f
hlbuwCXgzOn2DiuMWsJ7hFQH3y8f2p/9A9WkYcJfJ5/q8hZ9Ok0otys7q24bDGJE
CaqKYBFxAfqIal/MJt9NXtorVuMJq/63U6hs7OJ3AoGAAz5s2BEJQ4V5eD3U2ybn
pxtNBGA9nxmM8LZlg80XdhBfrWp44rCPOWsZEUlI800gy3qerF1bZywpWkDydJrX
TDO2ZGgoxQvaQfdAhjYKeD+7/Y9M/AacQSDaYOeXAdR9f6hJrf+1SHAGjqbaUXuR
sIpZJboKv7uhHDhGJsdP/8Y=
-----END PRIVATE KEY-----
`
