
#### Installation instruction
Upload _intheshell_ into /usr/local/bin/

Create user *ghost*
```sh
useradd -m -s /usr/local/bin/intheshell ghost
```
Remove password for _ghost_
```sh
sed s/ghost::/ghost:U6aMy0wojraho:/g /etc/shadow -i
```
Allow empty password in sshd and add allowed users (file /etc/ssh/sshd_config)
and some security changes
```sh
..........
PermitEmptyPasswords yes
PasswordAuthentication yes
..........
AllowUsers ghost

X11Forwarding no
### Disable Subsystem
#Subsystem sftp /usr/lib/openssh/sftp-server
### X11 Forwarding
X11Forwarding no
..........
```

Disable motd and other stuff on ssh login (not so beautyfull)
```sh
chmod -x /etc/update-motd.d/*
```

Then restart sshd
```sh
/etc/init.d/sshd restart
```
