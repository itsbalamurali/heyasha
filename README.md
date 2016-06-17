Asha: Artificial Intelligent Friend
====================================
[![GitHub release](https://img.shields.io/github/release/itsbalamurali/heyasha.svg?maxAge=2592000)](https://github.com/itsbalamurali/heyasha)
[![GitHub tag](https://img.shields.io/github/tag/itsbalamurali/heyasha.svg?maxAge=2592000)](https://github.com/itsbalamurali/heyasha)
[![Build Status](https://travis-ci.com/itsbalamurali/heyasha.svg?token=C7gvCDMuFC47B18pMTgy&branch=master)](https://travis-ci.com/itsbalamurali/heyasha)
[![Build Status](https://semaphoreci.com/api/v1/projects/b0cffd25-cf7f-44b5-8dbf-4f365e0eccd0/836471/badge.svg)](https://semaphoreci.com/itsbalamurali/heyasha)

### Overview

A Virtual & Extensible AI Friend(Bot) available on multiple platforms like:

* Facebook Messenger
* Telegram
* Skype
* Email
* Kik
* Android App
* iOS App

& Developer REST API for Developers to integrate asha into their applications

## Installation

Make sure you have a working Go environment. Go version 1.5+ is required. See the install instructions.

To install asha, simply run:

```
$ go get github.com/itsbalamurali/heyasha
```

Make sure your `PATH` includes to the `$GOPATH/bin` directory so your commands can be easily used:

```
export PATH=$PATH:$GOPATH/bin
```

To start Core Rest API server on port `80`:

```
$ heyasha
```

To Create/Import Intents & Aiml Knowledge base

```
$ asha-cli
```

## Rest API Reference

REST API Documentation can be found here : https://docs.heyasha.com 

## Tests
To run tests run 

```
$ go test
```

## License

Copyright (C) 2016 Balamurali Pandranki - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.