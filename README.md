## gif_anime_creator_wasm

## discription

This is an experimental web assembly program.  
Create animation gif from png file on web browser by web assembly.  
Creating animation gif by golang is based on this repository.  
[from-unknown/gif_anime_creator_go](https://github.com/from-unknown/gif_anime_creator_go)  
Referenced repository: [agnivade/shimmer](https://github.com/agnivade/shimmer)  

## how to test

1. go to ./go directory
1. build by `make` command
1. move `gifanimecreator.wasm` file to top directory
1. run server by `go run ./go/server/server/go`
1. access to `http://localhost:3434`
1. drag & drop png file to `Drop files here`
1. wait until amnimated gif is processed

*caution: big size picture file potentially takes very long time!*  
 
 ---
 
## 概要

[発表資料](https://speakerdeck.com/fromunknown/golangdezuo-tutawebassemblydehua-xiang-jia-gong)  
これは実験的なWeb Assemblyのプログラムです。  
Web Assemblyって何？という方はこちらの[スライド](https://speakerdeck.com/fromunknown/godewebassembly)を見ると雰囲気がつかめるかもしれません。  
ブラウザ上でWeb Assemblyを使ってpngファイルからアニメーションgifを生成します。  
アニメーションgifを作るGolangのプログラムは以下のレポジトリを元にしています。  
[from-unknown/gif_anime_creator_go](https://github.com/from-unknown/gif_anime_creator_go)  
参考にしたレポジトリ: [agnivade/shimmer](https://github.com/agnivade/shimmer)  

## how to test

1. `./go` ディレクトリに遷移する
1. `make` コマンドを使ってビルドする
1. `gifanimecreator.wasm` を上位のディレクトリに動かす
1. `go run ./go/server/server/go` でサーバーを起動する
1. `http://localhost:3434` にアクセスする
1. `Drop files here` にpngファイルをドラッグ＆ドロップする
1. gifアニメが生成されるまで待つ

*注意：大きなサイズの画像だととても時間がかかる可能性があります*  
