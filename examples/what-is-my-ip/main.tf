data "http" "my_ip" {
  url = "http://ipv4.icanhazip.com"
}

data "http" "bad" {
  url = "../dir/"
}

output "my_ip" {
  value = data.http.my_ip.body
}