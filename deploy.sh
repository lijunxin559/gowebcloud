#! /bin/sh
kill -9 $ (pgrep webserver)
cd ~/newweb/
git pull git@github.com:lijunxin559/gowebcloud.git
cd webserver/
./webserver &