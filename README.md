# grpcbuild2

Sample for gRPC build tool.


# Usage

以下でサーバーを立ち上げ。  

    $ go build ./server
    $ mkdir tmp
    $ cd tmp
    $ ../server

以下でそれぞれの client プログラムを実行。  

    $ go build ./client
    $ go build ./client-exec
    $ go build ./client-send
    $ go build ./client-recv

実行方法は以下の通り。

## client

gRPC のサンプルと同じで、引数に対し Greeting メッセージを返します。

    $ client
    2019/01/16 20:21:26 Greeting: Hello world

    $ client Hello
    2019/01/16 20:21:33 Greeting: Hello Hello

    $ client sago35
    2019/01/16 20:21:37 Greeting: Hello sago35


## client-exec

引数で指定したプログラムをサーバー側で実行します。

    $ client-exec go version
    go version go1.11.4 windows/amd64

    $ client-exec perl -E "say 'hello gRPC'"
    hello gRPC

    $ client-exec perl -e "die"
    2019/01/15 23:40:21 error: rpc error: code = Unknown desc = exit status 255 : Died at -e line 1.

## client-send

引数で指定したファイルをサーバー側に送信します。

    $ client-exec cmd /c dir /b

    $ echo hello > hello.txt

    $ client-send hello.txt

    $ client-exec cmd /c dir /b
    hello.txt

## client-recv

引数で指定したファイルをサーバー側から受信します。

    $ dir /b hello.txt
    ファイルが見つかりません

    $ client-recv hello.txt
    hello.txt

    $ dir /b hello.txt
    hello.txt






