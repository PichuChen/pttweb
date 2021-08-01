pttweb
======

PTT BBS Web Frontend.

In production on http://www.ptt.cc/

[![Build Status](https://travis-ci.org/ptt/pttweb.svg?branch=master)](https://travis-ci.org/ptt/pttweb)

Features
--------

 - List board index
 - Show articles
 - Render ANSI colors
 - Templating support
 - Ask user if he/she is over age 18 when entering some areas.

Configuration
-------------

See `config.go`.
Put them into a JSON-encoded file.

    $ ./pttweb -conf config.json


Compile
-------

If you want to compile boardd in pttbbs in Ubuntu, you may want to install `libgrpc++-dev` and `protobuf-compiler-grpc` first

Template
--------

To be documented.

License
-------

MIT
