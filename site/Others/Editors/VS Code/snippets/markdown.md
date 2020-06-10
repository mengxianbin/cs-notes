[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Others](https://mengxianbin.github.io/cs-notes/site/Others) /
[Editors](https://mengxianbin.github.io/cs-notes/site/Others/Editors) /
[VS Code](https://mengxianbin.github.io/cs-notes/site/Others/Editors/VS%20Code) /
[snippets](https://mengxianbin.github.io/cs-notes/site/Others/Editors/VS%20Code/snippets) /
[markdown](https://mengxianbin.github.io/cs-notes/site/Others/Editors/VS%20Code/snippets/markdown)

```json
{
	"java": {
		"prefix": "j",
		"body": [
			"```java",
			"$1",
			"```",
			"$0"
		],
		"description": "Java Code Block"
	},
	"c": {
		"prefix": "c",
		"body": [
			"```c",
			"$1",
			"```",
			"$0"
		],
		"description": "C Code Block"
	},
	"uml": {
		"prefix": "u",
		"body": [
			"```puml",
			"@startuml",
			"",
			"$0",
			"",
			"@enduml",
			"```",
			""
		],
		"description": "PlantUML"
	},
	"realize": {
		"prefix": "r",
		"body": [
			".up.|> "
		],
		"description": "PlantUML Realize"
	},
	"generalize": {
		"prefix": "g",
		"body": [
			"-up-|> "
		],
		"description": "PlantUML Generalize"
	},
	"rectangle": {
		"prefix": "t",
		"body": [
			"rectangle $1 {",
			"\t$2",
			"}",
			""
		],
		"description": "PlantUML Rectangle"
	},
	"storage": {
		"prefix": "s",
		"body": [
			"storage $1 {",
			"\t$2",
			"}",
			""
		],
		"description": "PlantUML Storage"
	},
	"hidden": {
		"prefix": "h",
		"body": [
			"-[hidden]- ",
		],
		"description": "PlantUML Hidden"
	},
	"http": {
		"prefix": "ht",
		"body": [
			"```http",
			"$1",
			"```",
			"$0"
		],
		"description": "HTTP"
	},
	"left": {
		"prefix": "l",
		"body": [
			"<- "
		],
		"description": "Left Arrow"
	},
	"arrow with comment": {
		"prefix": "t",
		"body": "-$1-> "
	}
}
```
