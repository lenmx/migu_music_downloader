#!/bin/bash

uNames=`uname -s`
osName=${uNames: 0: 4}

function buildFromMac() {
  echo "打包编译中($1)..."
  case $1 in
    'darwin')
      fyne package -os darwin -icon fav.png
    ;;
    'windows')
      # mac 交叉编译到 windows 需要安装 brew install mingw-w64，并设置环境变量 CC="x86_64-w64-mingw32-gcc"
      env CC="x86_64-w64-mingw32-gcc" fyne package -os windows -icon fav.png
    ;;
    *) echo '当前支持的交叉编译目标系统有 darwin|windows'
    ;;
  esac
  echo "完成"
}

function buildFromWin() {
  echo "todo"
}

function build() {
  case $osName in
    'Darw')
      buildFromMac $1
    ;;
    'Linu')
      echo 'linux系统暂不支持'
    ;;
    *)
      buildFromWin $1
    ;;
  esac
}

case "$1" in
	'build')
	build $2
	;;
	*)
	echo "usage: $0 {build target_os}"
	exit 1
	;;
esac