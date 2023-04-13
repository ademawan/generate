package main

import (
	"path/to/my/pluginmanager"
	"plugin"
)

func main() {
	plugin.Open("exampleplugin/exampleplugin.so")
	pluginmanager.Load("exampleplugin/exampleplugin.so")
}
