variable "do_token" {} # DigitalOcean API token
variable "ssh_fingerprint" {} # fingerprint of ssh key uploaded to DigitalOcean 
 
provider "digitalocean" {
 token = var.do_token
}
 
resource "digitalocean_droplet" "terraform-study" {
 image  = "ubuntu-18-04-x64"
 name   = "droplet"
 region = "fra1"
 size   = "s-1vcpu-1gb"
 ssh_keys = [
     var.ssh_fingerprint
 ]
}
