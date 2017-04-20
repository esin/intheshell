
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
```sh
..........
PermitEmptyPasswords yes
PasswordAuthentication yes
..........
AllowUsers ghost
..........
```
Then restart sshd
```sh
/etc/init.d/sshd restart
```
