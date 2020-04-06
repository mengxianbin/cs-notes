[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Distributions](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Distributions) /
[Ubuntu](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Distributions/Ubuntu) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Distributions/Ubuntu/Problems) /
[aer error](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Distributions/Ubuntu/Problems/aer%20error)

* phenomena
    * ubuntu bloked in the start window without login
    * recovery mode to login
    * /var/log/syslog is bigger than 1GB
    * edit /etc/default/grub to add "noaer"

* cause
    * hardware compatibility

* solution https://itsfoss.com/pcie-bus-error-severity-corrected/

---

* uninstall ubuntu
    * login windows
    * delete ubuntu partitions
        * command: diskpart
            * sel
            * del
            * disk
            * part
        * soft: partition assist
* reinstall ubuntu
    * connect network
    * install 3rd-party driven

---
