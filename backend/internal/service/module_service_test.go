package service

import (
	"context"
	"testing"
	"time"

	"github.com/Wei-Shaw/LightBridge/internal/modules"
)

type moduleServiceMemoryStore struct {
	items []modules.InstalledModule
}

func (s moduleServiceMemoryStore) ListInstalled(context.Context) ([]modules.InstalledModule, error) {
	return append([]modules.InstalledModule(nil), s.items...), nil
}

func (s moduleServiceMemoryStore) GetInstalled(context.Context, string) (*modules.InstalledModule, error) {
	panic("not implemented")
}

func (s moduleServiceMemoryStore) SaveInstalled(context.Context, modules.InstalledModule) error {
	panic("not implemented")
}

func (s moduleServiceMemoryStore) SavePermissions(context.Context, string, []modules.PermissionRecord) error {
	panic("not implemented")
}

func (s moduleServiceMemoryStore) ListPermissions(context.Context, string) ([]modules.PermissionRecord, error) {
	panic("not implemented")
}

func (s moduleServiceMemoryStore) ApprovePermissions(context.Context, string) error {
	panic("not implemented")
}

func (s moduleServiceMemoryStore) SetStatus(context.Context, string, modules.ModuleStatus, string) error {
	panic("not implemented")
}

func TestDecodeMarketplaceRegistryPreservesLocalizedModuleText(t *testing.T) {
	registry, err := decodeMarketplaceRegistry([]byte(`{
		"modules": [{
			"id": "openai",
			"version": "0.1.1",
			"type": "provider",
			"name": "OpenAI OAuth Provider",
			"name_i18n": {
				"en": "OpenAI OAuth Provider",
				"zh-CN": "OpenAI OAuth 提供商"
			},
			"description": "OpenAI provider module",
			"description_i18n": {
				"en": "OpenAI provider module",
				"zh-CN": "OpenAI OAuth 提供商模块"
			},
			"downloadUrl": "https://example.test/lightbridge-module-openai-0.1.1.tar.zst",
			"core": ">=0.1.0 <0.2.0",
			"capabilities": ["provider.adapter"]
		}]
	}`))
	if err != nil {
		t.Fatalf("decode marketplace registry: %v", err)
	}
	if got := registry.Modules[0].NameI18n["zh-CN"]; got != "OpenAI OAuth 提供商" {
		t.Fatalf("expected zh-CN module name, got %q", got)
	}
	if got := registry.Modules[0].DescriptionI18n["zh-CN"]; got != "OpenAI OAuth 提供商模块" {
		t.Fatalf("expected zh-CN module description, got %q", got)
	}
	if err := validateMarketplaceEntry(registry.Modules[0]); err != nil {
		t.Fatalf("localized marketplace entry should validate: %v", err)
	}
}

func TestUIManifestIncludesLocalizedRouteMenuAndAccountFormText(t *testing.T) {
	now := time.Now()
	svc := NewModuleService(moduleServiceMemoryStore{items: []modules.InstalledModule{{
		ID:      "anthropic-oauth",
		Name:    "Anthropic OAuth Provider",
		Type:    modules.ModuleTypeProvider,
		Version: "0.1.0",
		Status:  modules.ModuleStatusEnabled,
		Manifest: modules.Manifest{
			NameI18n: modules.LocalizedText{
				"en":    "Anthropic OAuth Provider",
				"zh-CN": "Anthropic OAuth 提供商",
			},
			Frontend: &modules.FrontendSpec{
				Entry: "frontend/remoteEntry.js",
				Routes: []modules.FrontendRouteSpec{{
					Path:  "/admin/providers/anthropic-oauth-module",
					Title: "Anthropic OAuth Provider",
					TitleI18n: modules.LocalizedText{
						"zh-CN": "Anthropic OAuth 提供商",
					},
					ExposedModule: "./AnthropicOAuthProviderSettings",
				}},
				Menu: []modules.FrontendMenuSpec{{
					Title: "Anthropic OAuth Provider",
					TitleI18n: modules.LocalizedText{
						"zh-CN": "Anthropic OAuth 提供商",
					},
					Path: "/admin/providers/anthropic-oauth-module",
				}},
				AccountForms: []modules.FrontendAccountFormSpec{{
					ProviderID:    "anthropic-oauth",
					ExposedModule: "./AnthropicOAuthAccountForm",
				}},
			},
		},
		InstalledAt: now,
	}}})

	items, err := svc.UIManifest(context.Background())
	if err != nil {
		t.Fatalf("load UI manifest: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected one UI manifest item, got %d", len(items))
	}
	if got := items[0].ModuleNameI18n["zh-CN"]; got != "Anthropic OAuth 提供商" {
		t.Fatalf("expected localized module name, got %q", got)
	}
	if got := items[0].Routes[0].TitleI18n["zh-CN"]; got != "Anthropic OAuth 提供商" {
		t.Fatalf("expected localized route title, got %q", got)
	}
	if got := items[0].Menu[0].TitleI18n["zh-CN"]; got != "Anthropic OAuth 提供商" {
		t.Fatalf("expected localized menu title, got %q", got)
	}
	if got := items[0].AccountForms[0].ProviderNameI18n["zh-CN"]; got != "Anthropic OAuth 提供商" {
		t.Fatalf("expected localized account form provider name, got %q", got)
	}
}
