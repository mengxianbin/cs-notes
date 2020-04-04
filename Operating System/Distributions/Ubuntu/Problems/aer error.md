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
