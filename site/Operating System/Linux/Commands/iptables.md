[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[iptables](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/iptables)

# iptables

## targets

* ACCEPT
* DROP
* RETURN

# tables

* -t / --tables
    * filter
    * nat

# OPTIONS

* -A, --append chain
* -D, --delete chain
* -I, --insert chain
* -L, --list [chain]
* -P, --policy chain target
* -F, --flush [chain]

# PARAMETERS

* -p, --protocol protocol
* -s, --source address[/mask][,...]
* -d, --destination address[/mask][,...]
* -j, --jump target
* -m, --match match

## SEE ALSO

* iptables-apply
* iptables-save
* iptables-restore
* iptables-extensions

---
