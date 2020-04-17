
* CAP
    * ZooKeeper: CP
    * Eureka: AP

* Availability
    * zookeeper
        * not serves when electing
    * eureka:
        * always serves
        * consistency: final, not real-time
        * in protection mode
            * serve, but not sync to other eureka nodes

* relationship between nodes
    * zookeeper
        * leader
        * flower
    * eureka
        * pair

* Partition tolerance
    * zookeeper: half living principle
    * eureka: pretection mode



---
