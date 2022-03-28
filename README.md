# installd

## Motivation

People are lazy as hell. And I am too.
I want to be able to install software using one single command.
Ideally I can use `.run` tlds for that.

Like:

```shell
curl https://installd.run | bash
```

And it should give me a bash script based on my os.

This is what `installd` is about.
It parses your curl/wget/what have you User Agent and returns smart stuff based on that.

Thank me later

## Installation

```shell

go install github.com/4thel00z/installd/...@latest

```

## License

This project is licensed under the GPL-3 license.
