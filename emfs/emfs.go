//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package emfs

import (
	"embed"
	"os"
	"strings"
	"sync"
)

const (
	FSName = "embed"
)

type IEmbedFS interface {
	// ReadFile 读取文件
	//  EmbedFS: dirname/filename or filename
	//  MemoryFS: filename
	ReadFile(name string) ([]byte, error)
}

// IResourceProvider 资源提供者接口
type IResourceProvider interface {
	IEmbedFS
	// Name 返回提供者名称
	Name() string
	// HasFile 检查文件是否存在
	HasFile(name string) bool
}

// ResourceManager 资源管理器
type ResourceManager struct {
	mu        sync.RWMutex
	providers map[string]IResourceProvider
}

var (
	gResManager = &ResourceManager{
		providers: make(map[string]IResourceProvider),
	}
)

// EmbedFSProvider embed.FS 提供者
type EmbedFSProvider struct {
	name string
	fs   embed.FS
}

func NewEmbedFSProvider(name string, fs embed.FS) *EmbedFSProvider {
	return &EmbedFSProvider{
		name: name,
		fs:   fs,
	}
}

func (p *EmbedFSProvider) Name() string {
	return p.name
}

func (p *EmbedFSProvider) ReadFile(name string) ([]byte, error) {
	if name == "" {
		return nil, os.ErrNotExist
	}
	if !strings.HasPrefix(name, p.Name()) {
		name = p.Name() + "/" + name
	}
	return p.fs.ReadFile(name)
}

func (p *EmbedFSProvider) HasFile(name string) bool {
	f, err := p.fs.Open(name)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}

// MemoryFSProvider 内存字节数据提供者
type MemoryFSProvider struct {
	name  string
	files map[string][]byte
}

func NewMemoryFSProvider(name string) *MemoryFSProvider {
	return &MemoryFSProvider{
		name:  name,
		files: make(map[string][]byte),
	}
}

func (p *MemoryFSProvider) Name() string {
	return p.name
}

func (p *MemoryFSProvider) AddFile(name string, data []byte) {
	p.files[name] = data
}

func (p *MemoryFSProvider) AddFiles(files map[string][]byte) {
	for k, v := range files {
		p.files[k] = v
	}
}

func (p *MemoryFSProvider) ReadFile(name string) ([]byte, error) {
	if name == "" {
		return nil, os.ErrNotExist
	}
	if data, ok := p.files[name]; ok {
		return data, nil
	}
	return nil, os.ErrNotExist
}

func (p *MemoryFSProvider) HasFile(name string) bool {
	_, ok := p.files[name]
	return ok
}

// RegisterEmbedFS 注册 embed.FS 资源
func RegisterEmbedFS(name string, fs embed.FS) {
	provider := NewEmbedFSProvider(name, fs)
	RegisterProvider(provider)
}

// RegisterMemoryFS 注册内存资源
func RegisterMemoryFS(name string) *MemoryFSProvider {
	provider := NewMemoryFSProvider(name)
	RegisterProvider(provider)
	return provider
}

// RegisterProvider 注册资源提供者
func RegisterProvider(provider IResourceProvider) {
	gResManager.mu.Lock()
	defer gResManager.mu.Unlock()
	gResManager.providers[provider.Name()] = provider
}

// UnregisterProvider 移除资源提供者
func UnregisterProvider(name string) {
	gResManager.mu.Lock()
	defer gResManager.mu.Unlock()
	delete(gResManager.providers, name)
}

// GetProvider 获取指定的资源提供者
func GetProvider(name string) (IResourceProvider, error) {
	gResManager.mu.RLock()
	defer gResManager.mu.RUnlock()
	provider, ok := gResManager.providers[name]
	if ok {
		return provider, nil
	}
	return nil, os.ErrNotExist
}

// GetResource 获取资源文件内容
//
//	需要知道明确的资源提供者名称和文件名
func GetResource(name, filename string) ([]byte, error) {
	provider, err := GetProvider(name)
	if err != nil {
		return nil, err
	}
	return provider.ReadFile(filename)
}

// HasProviders 检查资源提供者是否存在
func HasProviders(name string) bool {
	gResManager.mu.RLock()
	defer gResManager.mu.RUnlock()
	_, ok := gResManager.providers[name]
	return ok
}

// GetProviders 获取所有提供者名称
func GetProviders() []string {
	gResManager.mu.RLock()
	defer gResManager.mu.RUnlock()
	names := make([]string, 0, len(gResManager.providers))
	for _, provider := range gResManager.providers {
		names = append(names, provider.Name())
	}
	return names
}

// ClearProviders 清空所有提供者
func ClearProviders() {
	gResManager.mu.Lock()
	defer gResManager.mu.Unlock()
	gResManager.providers = make(map[string]IResourceProvider)
}
