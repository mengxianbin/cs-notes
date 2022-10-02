[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Markdown](https://mengxianbin.github.io/cs-notes/site/Language/Markdown) /
[mermaid](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/mermaid)

# Mermaid

- Arya - 在线 Markdown 编辑器
  - [https://markdown.lovejade.cn/?utm_source=github.com](https://markdown.lovejade.cn/?utm_source=github.com)

- 基于 vditor

#### Flow Chart

[https://github.com/knsv/mermaid#flowchart](https://github.com/knsv/mermaid#flowchart%E2%80%B8)

```mermaid
graph TD;
  A-->B;
  A-->C;
  B-->D;
  C-->D;
```

#### Sequence Diagram

[https://github.com/knsv/mermaid#sequence-diagram](https://github.com/knsv/mermaid#sequence-diagram%E2%80%B8)

```mermaid
sequenceDiagram
  participant Alice
  participant Bob
  Alice->John: Hello John, how are you?
  loop Healthcheck
      John->John: Fight against hypochondria
  end
  Note right of John: Rational thoughts <br/>prevail...
  John-->Alice: Great!
  John->Bob: How about you?
  Bob-->John: Jolly good!
```

#### Gantt Diagram

[https://github.com/knsv/mermaid#gantt-diagram](https://github.com/knsv/mermaid#gantt-diagram%E2%80%B8)

```mermaid
gantt
  title 项目开发流程
  section 项目确定
    需求分析       :a1, 2019-06-22, 3d
    可行性报告    :after a1, 5d
    概念验证       : 5d
  section 项目实施
    概要设计    :2019-07-05, 5d
    详细设计    :2019-07-08, 10d
    编码          :2019-07-15, 10d
    测试          :2019-07-22, 5d
  section 发布验收
    发布: 2d
    验收: 3d
```

---

