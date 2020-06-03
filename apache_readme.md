

### 1) Run `systemd` process in Linux.
    
   `sudo nano /lib/systemd/system/fortecom.service`
   
      ```.env                                                                                                                                          [Unit]
            Description=fortecom web server on Golang
            After=network.target
            
            [Service]
            Type=simple
            WorkingDirectory=/home/egorm/fortecom
            EnvironmentFile=/home/egorm/fortecom/.env
            ExecStart=/home/egorm/fortecom/rustamV
            
            [Install]
            WantedBy=multi-user.target
      ```
   1.1. Then reload system deamon service to imvoke your new service
   
   `systemctl daemon-reload`
   
   1.2. Start your new service
   
   `service fortecom start`
   
   1.3. Check status and logs in stdout of your service:
      
   `service fortecom status -l` -l for full logs
    
### 2) Configure VirtualHost for Apache and your server: 

`sudo nano /etc/httpd/conf.d/fotecom.conf`
``` 
<VirtualHost *:80>
        DocumentRoot /home/egorm/fortecom
        ErrorLog  /home/egorm/fortecom/logs/apache_logs/error.log
        CustomLog  /home/egorm/fortecom/logs/apache_logs/access.log combined
        ProxyPreserveHost On
        ProxyRequests Off
        ServerName www.fortecom123456.com
        ServerAlias example.com
        ProxyPass / http://127.0.0.1:8899/
        ProxyPassReverse / http://127.0.0.1:8899/
</VirtualHost>
```

Make possible for Apache to write logs in folder with project: 

`chcon -t httpd_sys_rw_content_t /home/egorm/fortecom/logs -R` 

2.1 sudo systemctl restart httpd 
