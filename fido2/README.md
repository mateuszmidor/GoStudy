

 go get  github.com/jsha/minica
 go install github.com/jsha/minica
 minica --domains localhost
 import generated "minica.pem" as firefox trusted CA 
 firefox https://localhost:8888