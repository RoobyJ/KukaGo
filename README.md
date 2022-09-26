# KukaGo
KukaGo is a tool for kuka robot programmers to help them in boring tasks. For now it just get data from $config.dat to excel, there is to a version in python [Kuka-Data]([https://www.google.com](https://github.com/RoobyJ/Kuka-Data/)) written works same as this but when comparing both working time program written in Golang works about 5 times faster ( I checked it on 32 kuka robot backups, python procced in ~2.9s Golang needed ~0.4s ) so i decided to continue this project in Golang

## Instalation 

download just all files ( needed Golang 1.19 )

## Usage
KukaGo-Excel
to use it just compile the the whole package and when executing pass your path of backups file this -p C:\Backups

## Roadmap

In future i want fo add collision check from .jt from generated in process simulate and cleaning src files from process simulate
