
#### Installation instruction
Upload _intheshell_ into /usr/local/bin/

Create user *ghost*
```sh
useradd -m -s /usr/local/bin/intheshell ghost
```
Remove password for _ghost_
```sh
sed -ri s/ghost:(!)?:/ghost:U6aMy0wojraho:/g /etc/shadow
```
Allow empty password in sshd and add allowed users (file /etc/ssh/sshd_config)
and some security changes
```sh
..........
PermitEmptyPasswords yes
PasswordAuthentication yes
..........
AllowUsers ghost

### Disable Subsystem
#Subsystem sftp /usr/lib/openssh/sftp-server
### X11 Forwarding
X11Forwarding no

# Adding chroot
Match User ghost
    ChrootDirectory /chroot/ghost
    AllowTcpForwarding no
..........
```

Adding chroot for user ghost
```sh
dir=/chroot/ghost
mkdir -p $dir
mkdir -p $dir/{dev,lib64,lib,bin,etc}
mknod -m 666 $dir/dev/null c 1 3
mknod -m 666 $dir/dev/tty c 5 0
mknod -m 666 $dir/dev/zero c 1 5
mknod -m 666 $dir/dev/random c 1 8
chown root:root $dir
chmod 0755 $dir
mkdir -p $dir/lib/x86_64-linux-gnu/
cp -v /lib/x86_64-linux-gnu/{libncurses.so.5,libtinfo.so.5,libdl.so.2,libc.so.6} $dir/lib/
cp -v /lib64/ld-linux-x86-64.so.2 $dir/lib64/
cat /etc/passwd | grep ghost > $dir/etc/passwd
touch $dir/etc/group
cp -av /bin/stty $dir/bin
cp -av /usr/local/bin/intheshell $dir/bin
cp -av /usr/local/bin/intheshell $dir/usr/local/bin
```

Disable motd and other stuff on ssh login (not so beautyfull)
```sh
chmod -x /etc/update-motd.d/*
```

Then restart sshd
```sh
/etc/init.d/ssh restart
```
