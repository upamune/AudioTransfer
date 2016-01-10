# AudioTransfer

## Description

AudioTransfer is [Content Transfer](https://www.sony.jp/support/audiosoftware/contenttransfer/download/) only Audio file on CLI.

Read an audio tag and copy it to ` `dst`/Artist/Album/Track-Title.Ext`

## Usage

`at -d `destination` `audio files` `

## Example

```bash
$ tree Music
Music
├── i☆Ris - Believe in.mp3
├── i☆Ris - Believer's HEAVEN.mp3
├── i☆Ris - Color.mp3
├── i☆Ris - Defy the fate.mp3
├── i☆Ris - Love Magic.mp3
├── i☆Ris - Make it!.mp3
├── i☆Ris - Realize!.mp3
├── i☆Ris - ayatsunagi.mp3
├── i☆Ris - §Rainbow.mp3
├── i☆Ris - 幻想曲WONDERLAND.mp3
└── i☆Ris - 流星.mp3

$ cd Music && mkdir dst
$ at -d /run/media/upamune/WALKMAN/Music *.mp3
$ tree /run/media/upamune/WALKMAN/Music

Music
└── i☆Ris
 └── We are i☆Ris!!!
  ├── 01-Make it!.mp3
  ├── 02-幻想曲WONDERLAND.mp3
  ├── 03-Defy the fate.mp3
  ├── 04-Believe in.mp3
  ├── 05-流星.mp3
  ├── 06-Love Magic.mp3
  ├── 07-Color.mp3
  ├── 08-§Rainbow.mp3
  ├── 09-Believer's HEAVEN.mp3
  ├── 10-Realize!.mp3
  └── 11-ayatsunagi.mp3

```

## Install

### Dependencies
You must have the static [taglib](http://taglib.github.io/) libraries installed in order to compile go-taglib.

```bash
$ go get -d github.com/upamune/AudioTransfer
$ cd $GOPATH/github.com/upamune/AudioTransfer
$ go build -o at && mv ./at $GOPATH/bin/at
```

## Contribution

1. Fork ([https://github.com/upamune/AudioTransfer/fork](https://github.com/upamune/AudioTransfer/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[upamune](https://github.com/upamune)
