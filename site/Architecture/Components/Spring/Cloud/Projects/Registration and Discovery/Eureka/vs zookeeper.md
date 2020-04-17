[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring) /
[Cloud](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Cloud) /
[Projects](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Cloud/Projects) /
[Registration and Discovery](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Cloud/Projects/Registration%20and%20Discovery) /
[Eureka](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Cloud/Projects/Registration%20and%20Discovery/Eureka) /
[vs zookeeper](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Cloud/Projects/Registration%20and%20Discovery/Eureka/vs%20zookeeper)


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
