[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[TCP](https://mengxianbin.github.io/cs-notes/site/Network/TCP) /
[FLAGS](https://mengxianbin.github.io/cs-notes/site/Network/TCP/FLAGS)

* SYN
    * Synchronization
        * Sequence Number
* ACK
    * Acknowledge
        * SN + 1
* FIN
    * Finish
        * terminate connection gracefully
        * only one side of the conversation is stopped
        * no data loss
        * receiver of FIN keeps communicating till it wants to

* RST
    * Reset
        * condition
            * sender feels something is wrong with the TCP connection
                * eg. receive unexpected `FLAG` or sequence number
            * the conversation should not exist
        * tell the other side to stop communication
        * the whole conversation is stopped
        * data is discarded
        * receiver has to stop communication
* PSH
    * Push
        * All data in buffer to be pushed to NL(sender) / AL (receiver)
        * Data is delivered in sequence

* URG
    * Urgent
        * Only the urgent data to be given to AL immediately
        * Data is delivered out of sequence

---
