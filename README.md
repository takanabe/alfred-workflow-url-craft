# alfred-workflow-url-craft

![URL Craft logo](logo.png)

A workflow that transforms a url into new one that allows some formats such as "Github Flavored Markdown link" or "shorten url" and such and such.

## Install

[Download the workflow file from here](https://github.com/takanabe/alfred-workflow-url-craft/raw/master/url-craft.alfredworkflow)

## Usage

Type following commands on Alfred main window.

* `um <URL>` changes the specified URL to a markdown link


## Development

Create a symlink to the workflow directory so Alfred can use information stored there.

```
alfred link
```

Build bianry after code changes

```
make
```

Export workflow file named `url-craft` to the directory root.
