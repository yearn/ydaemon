# yDaemon

## Install on server
Once on your server, create the configuration to start the daemon.
- We are using Ubuntu 22.04 (LTS) x64 for the server.
- We are using Golang for the API with GoEthereum, Gorm and GinGonic

### server.service
Run `nano /etc/systemd/system/server.service` and put this default config in it
```
[Unit]
Description=Yearn data API

[Service]
ExecStart=/root/ydaemon/yDaemon
Restart=always
EnvironmentFile=/etc/systemd/system/server.conf

[Install]
WantedBy=multi-user.target
```

### server.service
Run `nano /etc/systemd/system/server.conf` and put the environment variables in it.
```
RPC_URI_FOR_1=
RPC_URI_FOR_10=
RPC_URI_FOR_250=
RPC_URI_FOR_8453=
RPC_URI_FOR_42161=
```

### Needed packages
Install the packages needed to run the app
```
sudo apt-get install nginx #for the reverse proxy
sudo apt-get install certbot python3-certbot-nginx #for the SSL certificate
sudo apt-get update #some general update
sudo apt install golang #go
```

### Nginx configuration
First, be sure to point your domain name to the API on your DNS configuration.  
Then, we will setup Nginx. First, navigate to the enable websites with `cd /etc/nginx/sites-enabled`, remove the default configuration with `rm -rf default` and create the one for our API with `nano yearn.server.conf`.  
Add this in your nginx configuration file (`yearn.server.conf`). Be sure to replace `ydaemon.yearn.finance` by your own domain.
```perl
upstream golang {
    server 127.0.0.1:8080;
    keepalive 4;
}

server {
    location / {
        proxy_pass http://golang;
    }

    listen 80;
    listen [::]:80;
    server_name ydaemon.yearn.finance;
}
```
Run the `sudo certbot --nginx` command to ask for a SSL certificate. Follow the instruction and wait to get a `Successfully received certificate.`.

### Build the executable
Now we will need to build the executable. First, clone this repo in `~/`.
You can now build the project the same way you could do it on your machine: `go build -o yDaemon ./cmd && ./yDaemon`

### Start the daemon   
Run the following commands:
```sh
systemctl daemon-reload #reload the services
sudo systemctl start server.service #start your service
sudo systemctl status server.service #verify that your service is ready
```



