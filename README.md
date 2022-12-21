# whiterabb17/strongarm
 Simple tool to bruteforce (spray actually) different network protocols.
 whiterabb17/strongarm also supports restoration of interrupted tasks ("-restore").
 
 whiterabb17/strongarm currently supports: **rdp, ssh, ftp, Windows LDAP, http basic** and **digest authentication**


```
go run . -ul testUsernames.txt -pl testPasswords.txt -p ftp -t 192.168.56.102:21 -w 10
---------------+
Success: user:123
-------------------
```


```
go run . -ul testUsernames.txt -pl testPasswords.txt -p ftp -t 192.168.56.102:21 -w 10
--------

CTRL+C

go run . -restore
-------+
Success: user:123
-------------------
```

-ul   Path to file with **usernames**

-ul   Path to file with **passwords**

-p   Protocol to brute ( winldap, rdp, ssh, ftp, httpbasic, httpdigest )

-t   Target host. http://127.0.0.1:667/protected/folder/

-w   Number of workers (threads)

-restore use "progress.gob" to restore task
 

 
**Examples:**

```
strongarm.exe -ul testUsernames.txt -pl testPasswords.txt -p ssh -t 192.168.56.102 -w 10

strongarm.exe -ul testUsernames.txt -pl testPasswords.txt -p ftp -t 192.168.56.102:21 -w 10

strongarm.exe -ul testUsernames.txt -pl testPasswords.txt -p rdp -t 192.168.56.105 -w 10

strongarm.exe -ul testUsernames.txt -pl testPasswords.txt -p httpbasic -t http://192.168.56.102:80/2 -w 10 -ru -rp

strongarm.exe -ul testUsernames.txt -pl testPasswords.txt -p httpdigest -t http://192.168.56.102/1 -w 10

strongarm.exe -ul testUsernames.txt -pl testPasswords.txt -p winldap -t 192.168.56.106 -w 10

strongarm.exe -restore
```
