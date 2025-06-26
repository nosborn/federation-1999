# cmd/telnetd

## Historical Note

In 1999 we used a custom telnetd derived from 4.4BSD telnetd and modified to drop root privileges before running login.

Modern Linux makes it possible to run a stock telnetd as an unprivileged user.
