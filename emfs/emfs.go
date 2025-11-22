//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package emfs

import (
	"os"
)

var (
	resourcesFS IEmbedFS
	libsFS      IEmbedFS
)

type IEmbedFS interface {
	ReadFile(name string) ([]byte, error)
}

func SetLibsFS(lib IEmbedFS) {
	libsFS = lib
}

func SetResourcesFS(resource IEmbedFS) {
	resourcesFS = resource
}

func GetLibsFS() IEmbedFS {
	return libsFS
}

func GetResourcesFS() IEmbedFS {
	return resourcesFS
}

func GetResources(file string) ([]byte, error) {
	if GetResourcesFS() != nil {
		return GetResourcesFS().ReadFile(file)
	} else {
		return os.ReadFile(file)
	}
}

func GetLibs(fileName string) ([]byte, error) {
	if GetLibsFS() != nil {
		return GetLibsFS().ReadFile(fileName)
	} else {
		return os.ReadFile(fileName)
	}
}

func SetEMFS(libs IEmbedFS, resources IEmbedFS) {
	SetLibsFS(libs)
	SetResourcesFS(resources)
}
