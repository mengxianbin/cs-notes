[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Guidelines](https://mengxianbin.github.io/cs-notes/site/Guidelines) /
[Commit Message Conventions](https://mengxianbin.github.io/cs-notes/site/Guidelines/Commit%20Message%20Conventions)

# Commit Message Conventions

## Reference

- [github/angular/CONTRIBUTING.md#type](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type)
- [github/conventional-commits/v1.0.0-beta.4.md#summary](https://github.com/conventional-commits/conventionalcommits.org/blob/master/content/v1.0.0-beta.4/index.md#summary%E2%80%B8)
- [github/joshbuchea/semantic_commit_messages](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716#example)
- [seesparkbox/semantic_commit_messages](https://seesparkbox.com/foundry/semantic_commit_messages)
- [karma-runner/git-commit-msg](http://karma-runner.github.io/1.0/dev/git-commit-msg.html)

---

## Why

- Automatically determining a semantic version.
  - Automatically generating CHANGELOGs.
  - Automatically triggering build and publish processes.
- Communicating the nature of changes to teammates.

---

## Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

---

## Type

- feat
- fix
- docs
- style
- test
- chore
- refactor
- perf
- build
- ci

---
