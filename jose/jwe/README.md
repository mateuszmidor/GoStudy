# JSON Web Encryption

Secure the content with private key encryption and public key decryption.  
https://pkg.go.dev/github.com/go-jose/go-jose/v3@v3.0.0#example-package-JWE  

## Run

```bash
go run .
```

```json
input:
Lorem ipsum dolor sit amet

encrypted + serialized:
{"protected":"eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ","encrypted_key":"ovNWUw3nXSjaJ7tceSyx99GS5sAqs1O5AXyPdiLa0taci9JS4PeHyy8D3Y5Qh7tyDSftY_pY0mULYiatHEvxgo8QXD78ddpIhHE5WVtlQJ08UMXS2mmpswB1v49O3XV2vTGdrpasz6ERUA7I2Fq59SzxJUYYuK4Wmm_FjKPEPFhozg0xhL9hUDQHGkSTOv1AsV_RbYpJ4pZQTfqmEDkDo78m4-6vkPsM7qmaMkJC3HH9_JBXIfLRx9MOFnVu2NdTVa5rrqtEgSemR-ur96Aj_DifVCuvRbxH5sxScZM2DSuGbt9CTWAj_dUPqtJwHFZX_4vJaRfVm2B3gR2wVvv6_g","iv":"_JXBol8ZesGXzG2v","ciphertext":"RyGwyvB_kJ-2KoDLa0HVXtHLbJD2ddwJPW0","tag":"wbE9N6MiXtEp8aQ55A1ZCA"}

encrypted + compact serialized:
eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ.ovNWUw3nXSjaJ7tceSyx99GS5sAqs1O5AXyPdiLa0taci9JS4PeHyy8D3Y5Qh7tyDSftY_pY0mULYiatHEvxgo8QXD78ddpIhHE5WVtlQJ08UMXS2mmpswB1v49O3XV2vTGdrpasz6ERUA7I2Fq59SzxJUYYuK4Wmm_FjKPEPFhozg0xhL9hUDQHGkSTOv1AsV_RbYpJ4pZQTfqmEDkDo78m4-6vkPsM7qmaMkJC3HH9_JBXIfLRx9MOFnVu2NdTVa5rrqtEgSemR-ur96Aj_DifVCuvRbxH5sxScZM2DSuGbt9CTWAj_dUPqtJwHFZX_4vJaRfVm2B3gR2wVvv6_g._JXBol8ZesGXzG2v.RyGwyvB_kJ-2KoDLa0HVXtHLbJD2ddwJPW0.wbE9N6MiXtEp8aQ55A1ZCA

decrypted:
Lorem ipsum dolor sit amet
```