// SPDX-FileCopyrightText: Copyright 2021 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testWorkspaceClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const want = `{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"],"normalizesLineEndings":true,"changeAnnotationSupport":{"groupsOnLabel":true}},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true,"semanticTokens":{"refreshSupport":true},"codeLens":{"refreshSupport":true},"fileOperations":{"dynamicRegistration":true,"didCreate":true,"willCreate":true,"didRename":true,"willRename":true,"didDelete":true,"willDelete":true}}`
	wantType := WorkspaceClientCapabilities{
		ApplyEdit: true,
		WorkspaceEdit: &WorkspaceClientCapabilitiesWorkspaceEdit{
			DocumentChanges:       true,
			FailureHandling:       "FailureHandling",
			ResourceOperations:    []string{"ResourceOperations"},
			NormalizesLineEndings: true,
			ChangeAnnotationSupport: &WorkspaceClientCapabilitiesWorkspaceEditChangeAnnotationSupport{
				GroupsOnLabel: true,
			},
		},
		DidChangeConfiguration: &DidChangeConfigurationWorkspaceClientCapabilities{
			DynamicRegistration: true,
		},
		DidChangeWatchedFiles: &DidChangeWatchedFilesWorkspaceClientCapabilities{
			DynamicRegistration: true,
		},
		Symbol: &WorkspaceSymbolClientCapabilities{
			DynamicRegistration: true,
			SymbolKind: &WorkspaceSymbolSymbolKind{
				ValueSet: []SymbolKind{
					SymbolKindFile,
					SymbolKindModule,
					SymbolKindNamespace,
					SymbolKindPackage,
					SymbolKindClass,
					SymbolKindMethod,
				},
			},
		},
		ExecuteCommand: &ExecuteCommandClientCapabilities{
			DynamicRegistration: true,
		},
		WorkspaceFolders: true,
		Configuration:    true,
		SemanticTokens: &SemanticTokensWorkspaceClientCapabilities{
			RefreshSupport: true,
		},
		CodeLens: &CodeLensWorkspaceClientCapabilities{
			RefreshSupport: true,
		},
		FileOperations: &WorkspaceClientCapabilitiesFileOperations{
			DynamicRegistration: true,
			DidCreate:           true,
			WillCreate:          true,
			DidRename:           true,
			WillRename:          true,
			DidDelete:           true,
			WillDelete:          true,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          WorkspaceClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             WorkspaceClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got WorkspaceClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentSyncClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true,"didSave":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentSyncClientCapabilities{
		DynamicRegistration: true,
		WillSave:            true,
		WillSaveWaitUntil:   true,
		DidSave:             true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentSyncClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentSyncClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentSyncClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentSyncClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentSyncClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testCompletionTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true,"tagSupport":{"valueSet":[1]},"insertReplaceSupport":true,"resolveSupport":{"properties":["test","properties"]},"insertTextModeSupport":{"valueSet":[1,2]}},"completionItemKind":{"valueSet":[1]},"contextSupport":true}`
		wantNil = `{}`
	)
	wantType := CompletionTextDocumentClientCapabilities{
		DynamicRegistration: true,
		CompletionItem: &CompletionTextDocumentClientCapabilitiesItem{
			SnippetSupport:          true,
			CommitCharactersSupport: true,
			DocumentationFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
			DeprecatedSupport: true,
			PreselectSupport:  true,
			TagSupport: &CompletionTextDocumentClientCapabilitiesItemTagSupport{
				ValueSet: []CompletionItemTag{
					CompletionItemTagDeprecated,
				},
			},
			InsertReplaceSupport: true,
			ResolveSupport: &CompletionTextDocumentClientCapabilitiesItemResolveSupport{
				Properties: []string{"test", "properties"},
			},
			InsertTextModeSupport: &CompletionTextDocumentClientCapabilitiesItemInsertTextModeSupport{
				ValueSet: []InsertTextMode{
					InsertTextModeAsIs,
					InsertTextModeAdjustIndentation,
				},
			},
		},
		CompletionItemKind: &CompletionTextDocumentClientCapabilitiesItemKind{
			ValueSet: []CompletionItemKind{CompletionItemKindText},
		},
		ContextSupport: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          CompletionTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          CompletionTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             CompletionTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             CompletionTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CompletionTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testHoverTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]}`
		wantNil = `{}`
	)
	wantType := HoverTextDocumentClientCapabilities{
		DynamicRegistration: true,
		ContentFormat: []MarkupKind{
			PlainText,
			Markdown,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          HoverTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          HoverTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             HoverTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             HoverTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got HoverTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testSignatureHelpTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"],"parameterInformation":{"labelOffsetSupport":true},"activeParameterSupport":true},"contextSupport":true}`
		wantNil = `{}`
	)
	wantType := SignatureHelpTextDocumentClientCapabilities{
		DynamicRegistration: true,
		SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
			DocumentationFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
			ParameterInformation: &TextDocumentClientCapabilitiesParameterInformation{
				LabelOffsetSupport: true,
			},
			ActiveParameterSupport: true,
		},
		ContextSupport: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          SignatureHelpTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          SignatureHelpTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             SignatureHelpTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             SignatureHelpTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got SignatureHelpTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testReferencesTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := ReferencesTextDocumentClientCapabilities{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          ReferencesTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          ReferencesTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             ReferencesTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             ReferencesTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ReferencesTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testDocumentHighlightClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := DocumentHighlightClientCapabilities{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DocumentHighlightClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DocumentHighlightClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             DocumentHighlightClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             DocumentHighlightClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentHighlightClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testDocumentSymbolClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true,"tagSupport":{"valueSet":[1]},"labelSupport":true}`
		wantNil = `{}`
	)
	wantType := DocumentSymbolClientCapabilities{
		DynamicRegistration: true,
		SymbolKind: &WorkspaceSymbolClientCapabilitiesKind{
			ValueSet: []SymbolKind{
				SymbolKindFile,
				SymbolKindModule,
				SymbolKindNamespace,
				SymbolKindPackage,
				SymbolKindClass,
				SymbolKindMethod,
			},
		},
		HierarchicalDocumentSymbolSupport: true,
		TagSupport: &DocumentSymbolClientCapabilitiesTagSupport{
			ValueSet: []SymbolTag{
				SymbolTagDeprecated,
			},
		},
		LabelSupport: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DocumentSymbolClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DocumentSymbolClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             DocumentSymbolClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             DocumentSymbolClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentSymbolClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesFormatting(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesFormatting{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesFormatting
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesFormatting{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesFormatting{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesFormatting
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesRangeFormatting(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesRangeFormatting{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesRangeFormatting
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesRangeFormatting{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesRangeFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesRangeFormatting{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesRangeFormatting
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilitiesOnTypeFormatting(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilitiesOnTypeFormatting{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilitiesOnTypeFormatting
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilitiesOnTypeFormatting{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilitiesOnTypeFormatting
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilitiesOnTypeFormatting{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilitiesOnTypeFormatting
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testDeclarationTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := DeclarationTextDocumentClientCapabilities{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DeclarationTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DeclarationTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             DeclarationTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             DeclarationTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DeclarationTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testDefinitionTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := DefinitionTextDocumentClientCapabilities{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DefinitionTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DefinitionTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             DefinitionTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             DefinitionTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DefinitionTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTypeDefinitionTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := TypeDefinitionTextDocumentClientCapabilities{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TypeDefinitionTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TypeDefinitionTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TypeDefinitionTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TypeDefinitionTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TypeDefinitionTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testImplementationTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"linkSupport":true}`
		wantNil = `{}`
	)
	wantType := ImplementationTextDocumentClientCapabilities{
		DynamicRegistration: true,
		LinkSupport:         true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          ImplementationTextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          ImplementationTextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             ImplementationTextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             ImplementationTextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ImplementationTextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testCodeActionClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}},"isPreferredSupport":true,"disabledSupport":true,"dataSupport":true,"resolveSupport":{"properties":["testProperties"]},"honorsChangeAnnotations":true}`
		wantNil = `{}`
	)
	wantType := CodeActionClientCapabilities{
		DynamicRegistration: true,
		CodeActionLiteralSupport: &CodeActionClientCapabilitiesLiteralSupport{
			CodeActionKind: &CodeActionClientCapabilitiesKind{
				ValueSet: []CodeActionKind{
					QuickFix,
					Refactor,
					RefactorExtract,
					RefactorRewrite,
					Source,
					SourceOrganizeImports,
				},
			},
		},
		IsPreferredSupport: true,
		DisabledSupport:    true,
		DataSupport:        true,
		ResolveSupport: &CodeActionClientCapabilitiesResolveSupport{
			Properties: []string{"testProperties"},
		},
		HonorsChangeAnnotations: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          CodeActionClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          CodeActionClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             CodeActionClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             CodeActionClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeActionClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testCodeLensClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := CodeLensClientCapabilities{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          CodeLensClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          CodeLensClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             CodeLensClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             CodeLensClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got CodeLensClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testDocumentLinkClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"tooltipSupport":true}`
		wantNil = `{}`
	)
	wantType := DocumentLinkClientCapabilities{
		DynamicRegistration: true,
		TooltipSupport:      true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DocumentLinkClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DocumentLinkClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             DocumentLinkClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             DocumentLinkClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentLinkClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testDocumentColorClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true}`
		wantNil = `{}`
	)
	wantType := DocumentColorClientCapabilities{
		DynamicRegistration: true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          DocumentColorClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          DocumentColorClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             DocumentColorClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             DocumentColorClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got DocumentColorClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testRenameClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"prepareSupport":true,"prepareSupportDefaultBehavior":1,"honorsChangeAnnotations":true}`
		wantNil = `{}`
	)
	wantType := RenameClientCapabilities{
		DynamicRegistration:           true,
		PrepareSupport:                true,
		PrepareSupportDefaultBehavior: PrepareSupportDefaultBehaviorIdentifier,
		HonorsChangeAnnotations:       true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          RenameClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          RenameClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             RenameClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             RenameClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got RenameClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testPublishDiagnosticsClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"relatedInformation":true,"tagSupport":{"valueSet":[2,1]},"versionSupport":true,"codeDescriptionSupport":true,"dataSupport":true}`
		wantNil = `{}`
	)
	wantType := PublishDiagnosticsClientCapabilities{
		RelatedInformation: true,
		TagSupport: &PublishDiagnosticsClientCapabilitiesTagSupport{
			ValueSet: []DiagnosticTag{
				DiagnosticTagDeprecated,
				DiagnosticTagUnnecessary,
			},
		},
		VersionSupport:         true,
		CodeDescriptionSupport: true,
		DataSupport:            true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          PublishDiagnosticsClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          PublishDiagnosticsClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             PublishDiagnosticsClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             PublishDiagnosticsClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got PublishDiagnosticsClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testFoldingRangeClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"dynamicRegistration":true,"rangeLimit":5,"lineFoldingOnly":true}`
		wantNil = `{}`
	)
	wantType := FoldingRangeClientCapabilities{
		DynamicRegistration: true,
		RangeLimit:          uint32(5),
		LineFoldingOnly:     true,
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          FoldingRangeClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          FoldingRangeClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             FoldingRangeClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             FoldingRangeClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got FoldingRangeClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testTextDocumentClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"synchronization":{"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true,"didSave":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"publishDiagnostics":{"relatedInformation":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":5,"lineFoldingOnly":true},"selectionRange":{"dynamicRegistration":true},"callHierarchy":{"dynamicRegistration":true},"semanticTokens":{"dynamicRegistration":true,"requests":{"range":true,"full":true},"tokenTypes":["test","tokenTypes"],"tokenModifiers":["test","tokenModifiers"],"formats":["relative"],"overlappingTokenSupport":true,"multilineTokenSupport":true},"linkedEditingRange":{"dynamicRegistration":true},"moniker":{"dynamicRegistration":true}}`
		wantNil = `{}`
	)
	wantType := TextDocumentClientCapabilities{
		Synchronization: &TextDocumentSyncClientCapabilities{
			DynamicRegistration: true,
			WillSave:            true,
			WillSaveWaitUntil:   true,
			DidSave:             true,
		},
		Completion: &CompletionTextDocumentClientCapabilities{
			DynamicRegistration: true,
			CompletionItem: &CompletionTextDocumentClientCapabilitiesItem{
				SnippetSupport:          true,
				CommitCharactersSupport: true,
				DocumentationFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
				DeprecatedSupport: true,
				PreselectSupport:  true,
			},
			CompletionItemKind: &CompletionTextDocumentClientCapabilitiesItemKind{
				ValueSet: []CompletionItemKind{CompletionItemKindText},
			},
			ContextSupport: true,
		},
		Hover: &HoverTextDocumentClientCapabilities{
			DynamicRegistration: true,
			ContentFormat: []MarkupKind{
				PlainText,
				Markdown,
			},
		},
		SignatureHelp: &SignatureHelpTextDocumentClientCapabilities{
			DynamicRegistration: true,
			SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
				DocumentationFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
			},
		},
		Declaration: &DeclarationTextDocumentClientCapabilities{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		Definition: &DefinitionTextDocumentClientCapabilities{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		TypeDefinition: &TypeDefinitionTextDocumentClientCapabilities{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		Implementation: &ImplementationTextDocumentClientCapabilities{
			DynamicRegistration: true,
			LinkSupport:         true,
		},
		References: &ReferencesTextDocumentClientCapabilities{
			DynamicRegistration: true,
		},
		DocumentHighlight: &DocumentHighlightClientCapabilities{
			DynamicRegistration: true,
		},
		DocumentSymbol: &DocumentSymbolClientCapabilities{
			DynamicRegistration: true,
			SymbolKind: &WorkspaceSymbolClientCapabilitiesKind{
				ValueSet: []SymbolKind{
					SymbolKindFile,
					SymbolKindModule,
					SymbolKindNamespace,
					SymbolKindPackage,
					SymbolKindClass,
					SymbolKindMethod,
				},
			},
			HierarchicalDocumentSymbolSupport: true,
		},
		CodeAction: &CodeActionClientCapabilities{
			DynamicRegistration: true,
			CodeActionLiteralSupport: &CodeActionClientCapabilitiesLiteralSupport{
				CodeActionKind: &CodeActionClientCapabilitiesKind{
					ValueSet: []CodeActionKind{
						QuickFix,
						Refactor,
						RefactorExtract,
						RefactorRewrite,
						Source,
						SourceOrganizeImports,
					},
				},
			},
		},
		CodeLens: &CodeLensClientCapabilities{
			DynamicRegistration: true,
		},
		DocumentLink: &DocumentLinkClientCapabilities{
			DynamicRegistration: true,
		},
		ColorProvider: &DocumentColorClientCapabilities{
			DynamicRegistration: true,
		},
		Formatting: &DocumentFormattingClientCapabilities{
			DynamicRegistration: true,
		},
		RangeFormatting: &DocumentRangeFormattingClientCapabilities{
			DynamicRegistration: true,
		},
		OnTypeFormatting: &DocumentOnTypeFormattingClientCapabilities{
			DynamicRegistration: true,
		},
		PublishDiagnostics: &PublishDiagnosticsClientCapabilities{
			RelatedInformation: true,
		},
		Rename: &RenameClientCapabilities{
			DynamicRegistration: true,
			PrepareSupport:      true,
		},
		FoldingRange: &FoldingRangeClientCapabilities{
			DynamicRegistration: true,
			RangeLimit:          uint32(5),
			LineFoldingOnly:     true,
		},
		SelectionRange: &SelectionRangeClientCapabilities{
			DynamicRegistration: true,
		},
		CallHierarchy: &CallHierarchyClientCapabilities{
			DynamicRegistration: true,
		},
		SemanticTokens: &SemanticTokensClientCapabilities{
			DynamicRegistration: true,
			Requests: SemanticTokensWorkspaceClientCapabilitiesRequests{
				Range: true,
				Full:  true,
			},
			TokenTypes:     []string{"test", "tokenTypes"},
			TokenModifiers: []string{"test", "tokenModifiers"},
			Formats: []TokenFormat{
				TokenFormatRelative,
			},
			OverlappingTokenSupport: true,
			MultilineTokenSupport:   true,
		},
		LinkedEditingRange: &LinkedEditingRangeClientCapabilities{
			DynamicRegistration: true,
		},
		Moniker: &MonikerClientCapabilities{
			DynamicRegistration: true,
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          TextDocumentClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          TextDocumentClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             TextDocumentClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             TextDocumentClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got TextDocumentClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}

func testClientCapabilities(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	const (
		want    = `{"workspace":{"applyEdit":true,"workspaceEdit":{"documentChanges":true,"failureHandling":"FailureHandling","resourceOperations":["ResourceOperations"]},"didChangeConfiguration":{"dynamicRegistration":true},"didChangeWatchedFiles":{"dynamicRegistration":true},"symbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]}},"executeCommand":{"dynamicRegistration":true},"workspaceFolders":true,"configuration":true},"textDocument":{"synchronization":{"dynamicRegistration":true,"willSave":true,"willSaveWaitUntil":true,"didSave":true},"completion":{"dynamicRegistration":true,"completionItem":{"snippetSupport":true,"commitCharactersSupport":true,"documentationFormat":["plaintext","markdown"],"deprecatedSupport":true,"preselectSupport":true},"completionItemKind":{"valueSet":[1]},"contextSupport":true},"hover":{"dynamicRegistration":true,"contentFormat":["plaintext","markdown"]},"signatureHelp":{"dynamicRegistration":true,"signatureInformation":{"documentationFormat":["plaintext","markdown"]}},"declaration":{"dynamicRegistration":true,"linkSupport":true},"definition":{"dynamicRegistration":true,"linkSupport":true},"typeDefinition":{"dynamicRegistration":true,"linkSupport":true},"implementation":{"dynamicRegistration":true,"linkSupport":true},"references":{"dynamicRegistration":true},"documentHighlight":{"dynamicRegistration":true},"documentSymbol":{"dynamicRegistration":true,"symbolKind":{"valueSet":[1,2,3,4,5,6]},"hierarchicalDocumentSymbolSupport":true},"codeAction":{"dynamicRegistration":true,"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"documentLink":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":true},"formatting":{"dynamicRegistration":true},"rangeFormatting":{"dynamicRegistration":true},"onTypeFormatting":{"dynamicRegistration":true},"publishDiagnostics":{"relatedInformation":true},"rename":{"dynamicRegistration":true,"prepareSupport":true},"foldingRange":{"dynamicRegistration":true,"rangeLimit":5,"lineFoldingOnly":true},"selectionRange":{"dynamicRegistration":true},"callHierarchy":{"dynamicRegistration":true},"semanticTokens":{"dynamicRegistration":true,"requests":{"range":true,"full":true},"tokenTypes":["test","tokenTypes"],"tokenModifiers":["test","tokenModifiers"],"formats":["relative"],"overlappingTokenSupport":true,"multilineTokenSupport":true},"linkedEditingRange":{"dynamicRegistration":true},"moniker":{"dynamicRegistration":true}},"window":{"workDoneProgress":true,"showMessage":{"messageActionItem":{"additionalPropertiesSupport":true}},"showDocument":{"support":true}},"general":{"regularExpressions":{"engine":"ECMAScript","version":"ES2020"},"markdown":{"parser":"marked","version":"1.1.0"}},"experimental":"testExperimental"}`
		wantNil = `{}`
	)
	wantType := ClientCapabilities{
		Workspace: &WorkspaceClientCapabilities{
			ApplyEdit: true,
			WorkspaceEdit: &WorkspaceClientCapabilitiesWorkspaceEdit{
				DocumentChanges:    true,
				FailureHandling:    "FailureHandling",
				ResourceOperations: []string{"ResourceOperations"},
			},
			DidChangeConfiguration: &DidChangeConfigurationWorkspaceClientCapabilities{
				DynamicRegistration: true,
			},
			DidChangeWatchedFiles: &DidChangeWatchedFilesWorkspaceClientCapabilities{
				DynamicRegistration: true,
			},
			Symbol: &WorkspaceSymbolClientCapabilities{
				DynamicRegistration: true,
				SymbolKind: &WorkspaceSymbolSymbolKind{
					ValueSet: []SymbolKind{
						SymbolKindFile,
						SymbolKindModule,
						SymbolKindNamespace,
						SymbolKindPackage,
						SymbolKindClass,
						SymbolKindMethod,
					},
				},
			},
			ExecuteCommand: &ExecuteCommandClientCapabilities{
				DynamicRegistration: true,
			},
			WorkspaceFolders: true,
			Configuration:    true,
		},
		TextDocument: &TextDocumentClientCapabilities{
			Synchronization: &TextDocumentSyncClientCapabilities{
				DynamicRegistration: true,
				WillSave:            true,
				WillSaveWaitUntil:   true,
				DidSave:             true,
			},
			Completion: &CompletionTextDocumentClientCapabilities{
				DynamicRegistration: true,
				CompletionItem: &CompletionTextDocumentClientCapabilitiesItem{
					SnippetSupport:          true,
					CommitCharactersSupport: true,
					DocumentationFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
					DeprecatedSupport: true,
					PreselectSupport:  true,
				},
				CompletionItemKind: &CompletionTextDocumentClientCapabilitiesItemKind{
					ValueSet: []CompletionItemKind{CompletionItemKindText},
				},
				ContextSupport: true,
			},
			Hover: &HoverTextDocumentClientCapabilities{
				DynamicRegistration: true,
				ContentFormat: []MarkupKind{
					PlainText,
					Markdown,
				},
			},
			SignatureHelp: &SignatureHelpTextDocumentClientCapabilities{
				DynamicRegistration: true,
				SignatureInformation: &TextDocumentClientCapabilitiesSignatureInformation{
					DocumentationFormat: []MarkupKind{
						PlainText,
						Markdown,
					},
				},
			},
			Declaration: &DeclarationTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			Definition: &DefinitionTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			TypeDefinition: &TypeDefinitionTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			Implementation: &ImplementationTextDocumentClientCapabilities{
				DynamicRegistration: true,
				LinkSupport:         true,
			},
			References: &ReferencesTextDocumentClientCapabilities{
				DynamicRegistration: true,
			},
			DocumentHighlight: &DocumentHighlightClientCapabilities{
				DynamicRegistration: true,
			},
			DocumentSymbol: &DocumentSymbolClientCapabilities{
				DynamicRegistration: true,
				SymbolKind: &WorkspaceSymbolClientCapabilitiesKind{
					ValueSet: []SymbolKind{
						SymbolKindFile,
						SymbolKindModule,
						SymbolKindNamespace,
						SymbolKindPackage,
						SymbolKindClass,
						SymbolKindMethod,
					},
				},
				HierarchicalDocumentSymbolSupport: true,
			},
			CodeAction: &CodeActionClientCapabilities{
				DynamicRegistration: true,
				CodeActionLiteralSupport: &CodeActionClientCapabilitiesLiteralSupport{
					CodeActionKind: &CodeActionClientCapabilitiesKind{
						ValueSet: []CodeActionKind{
							QuickFix,
							Refactor,
							RefactorExtract,
							RefactorRewrite,
							Source,
							SourceOrganizeImports,
						},
					},
				},
			},
			CodeLens: &CodeLensClientCapabilities{
				DynamicRegistration: true,
			},
			DocumentLink: &DocumentLinkClientCapabilities{
				DynamicRegistration: true,
			},
			ColorProvider: &DocumentColorClientCapabilities{
				DynamicRegistration: true,
			},
			Formatting: &DocumentFormattingClientCapabilities{
				DynamicRegistration: true,
			},
			RangeFormatting: &DocumentRangeFormattingClientCapabilities{
				DynamicRegistration: true,
			},
			OnTypeFormatting: &DocumentOnTypeFormattingClientCapabilities{
				DynamicRegistration: true,
			},
			PublishDiagnostics: &PublishDiagnosticsClientCapabilities{
				RelatedInformation: true,
			},
			Rename: &RenameClientCapabilities{
				DynamicRegistration: true,
				PrepareSupport:      true,
			},
			FoldingRange: &FoldingRangeClientCapabilities{
				DynamicRegistration: true,
				RangeLimit:          uint32(5),
				LineFoldingOnly:     true,
			},
			SelectionRange: &SelectionRangeClientCapabilities{
				DynamicRegistration: true,
			},
			CallHierarchy: &CallHierarchyClientCapabilities{
				DynamicRegistration: true,
			},
			SemanticTokens: &SemanticTokensClientCapabilities{
				DynamicRegistration: true,
				Requests: SemanticTokensWorkspaceClientCapabilitiesRequests{
					Range: true,
					Full:  true,
				},
				TokenTypes:     []string{"test", "tokenTypes"},
				TokenModifiers: []string{"test", "tokenModifiers"},
				Formats: []TokenFormat{
					TokenFormatRelative,
				},
				OverlappingTokenSupport: true,
				MultilineTokenSupport:   true,
			},
			LinkedEditingRange: &LinkedEditingRangeClientCapabilities{
				DynamicRegistration: true,
			},
			Moniker: &MonikerClientCapabilities{
				DynamicRegistration: true,
			},
		},
		Window: &WindowClientCapabilities{
			WorkDoneProgress: true,
			ShowMessage: &ShowMessageRequestClientCapabilities{
				MessageActionItem: &ShowMessageRequestClientCapabilitiesMessageActionItem{
					AdditionalPropertiesSupport: true,
				},
			},
			ShowDocument: &ShowDocumentClientCapabilities{
				Support: true,
			},
		},
		General: &GeneralClientCapabilities{
			RegularExpressions: &RegularExpressionsClientCapabilities{
				Engine:  "ECMAScript",
				Version: "ES2020",
			},
			Markdown: &MarkdownClientCapabilities{
				Parser:  "marked",
				Version: "1.1.0",
			},
		},
		Experimental: "testExperimental",
	}

	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			field          ClientCapabilities
			want           string
			wantMarshalErr bool
			wantErr        bool
		}{
			{
				name:           "Valid",
				field:          wantType,
				want:           want,
				wantMarshalErr: false,
				wantErr:        false,
			},
			{
				name:           "ValidNilAll",
				field:          ClientCapabilities{},
				want:           wantNil,
				wantMarshalErr: false,
				wantErr:        false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := marshal(&tt.field)
				if (err != nil) != tt.wantMarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(string(got), tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name             string
			field            string
			want             ClientCapabilities
			wantUnmarshalErr bool
			wantErr          bool
		}{
			{
				name:             "Valid",
				field:            want,
				want:             wantType,
				wantUnmarshalErr: false,
				wantErr:          false,
			},
			{
				name:             "ValidNilAll",
				field:            wantNil,
				want:             ClientCapabilities{},
				wantUnmarshalErr: false,
				wantErr:          false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got ClientCapabilities
				if err := unmarshal([]byte(tt.field), &got); (err != nil) != tt.wantUnmarshalErr {
					t.Fatal(err)
				}

				if diff := cmp.Diff(got, tt.want); (diff != "") != tt.wantErr {
					t.Errorf("%s: wantErr: %t\n(-got, +want)\n%s", tt.name, tt.wantErr, diff)
				}
			})
		}
	})
}
