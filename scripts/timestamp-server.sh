#!/bin/sh

/usr/bin/timestamp-server serve --host=0.0.0.0 --port=3000 --timestamp-signer=memory
