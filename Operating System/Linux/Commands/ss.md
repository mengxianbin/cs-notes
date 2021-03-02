# ss

```man
NAME
       ss - another utility to investigate sockets

SYNOPSIS
       ss [options] [ FILTER ]

DESCRIPTION
       ss is used to dump socket statistics. It allows showing information similar to netstat.  It can display more TCP and state information than other tools.       

OPTIONS
       When no option is used ss displays a list of open non-listening sockets (e.g. TCP/UNIX/UDP) that have established connection.

       -h, --help
              Show summary of options.

       -V, --version
              Output version information.

       -n, --numeric
              Do not try to resolve service names. Show exact bandwidth values, instead of human-readable.              

       -a, --all
              Display both listening and non-listening (for TCP this means established connections) sockets.

       -l, --listening
              Display only listening sockets (these are omitted by default).

       -p, --processes
              Show process using socket.

       -i, --info
              Show internal TCP information. Below fields may appear:
              ...
              
       -K, --kill
              Attempts  to  forcibly close sockets. This option displays sockets that are successfully closed and silently skips sockets that the kernel does not support closing.
              It supports IPv4 and IPv6 sockets only.

       -s, --summary
              Print summary statistics. This option does not parse socket lists obtaining summary from various sources. It is useful when amount of sockets is so huge that  pars‐
              ing /proc/net/tcp is painful.                            

       -4, --ipv4
              Display only IP version 4 sockets (alias for -f inet).

       -6, --ipv6
              Display only IP version 6 sockets (alias for -f inet6).

       -t, --tcp
              Display TCP sockets.

       -u, --udp
              Display UDP sockets.              
```