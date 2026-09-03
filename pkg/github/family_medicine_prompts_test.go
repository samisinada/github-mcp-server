package github

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func promptText(t *testing.T, result *mcp.GetPromptResult) string {
	t.Helper()
	var b string
	for _, msg := range result.Messages {
		text, ok := msg.Content.(*mcp.TextContent)
		require.True(t, ok)
		b += text.Text + "\n"
	}
	return b
}

func TestFamilyMedicineClinicOpsPrompt(t *testing.T) {
	p := FamilyMedicineClinicOpsPrompt(stubTranslation)
	assert.Equal(t, "family_medicine_clinic_ops", p.Prompt.Name)
	assert.Equal(t, ToolsetMetadataIssues.ID, p.Toolset.ID)
	require.NotNil(t, p.Handler)

	result, err := p.Handler(context.Background(), &mcp.GetPromptRequest{
		Params: &mcp.GetPromptParams{
			Name: p.Prompt.Name,
			Arguments: map[string]string{
				"owner": "samisinada",
				"repo":  "github-mcp-server",
				"area":  "rooming",
				"need":  "PHQ-2 missing from MA rooming template",
			},
		},
	})
	require.NoError(t, err)
	body := promptText(t, result)
	assert.Contains(t, body, "Never write patient names")
	assert.Contains(t, body, "samisinada/github-mcp-server")
	assert.Contains(t, body, "rooming")
	assert.Contains(t, body, "PHQ-2 missing from MA rooming template")
	assert.Contains(t, body, "clinic-ops,no-phi")
}

func TestFamilyMedicineQIProtocolPrompt(t *testing.T) {
	p := FamilyMedicineQIProtocolPrompt(stubTranslation)
	assert.Equal(t, "family_medicine_qi_protocol", p.Prompt.Name)

	result, err := p.Handler(context.Background(), &mcp.GetPromptRequest{
		Params: &mcp.GetPromptParams{
			Name: p.Prompt.Name,
			Arguments: map[string]string{
				"owner":  "samisinada",
				"repo":   "github-mcp-server",
				"title":  "Colorectal screening standing order ages",
				"change": "Align FIT interval with current org protocol",
				"labels": "qi,protocol",
			},
		},
	})
	require.NoError(t, err)
	body := promptText(t, result)
	assert.Contains(t, body, "who may act")
	assert.Contains(t, body, "Align FIT interval with current org protocol")
	assert.Contains(t, body, "qi,protocol")
	assert.NotContains(t, body, "DEA")
}

func TestFamilyMedicineCredentialWatchPrompt(t *testing.T) {
	p := FamilyMedicineCredentialWatchPrompt(stubTranslation)
	assert.Equal(t, "family_medicine_credential_watch", p.Prompt.Name)

	result, err := p.Handler(context.Background(), &mcp.GetPromptRequest{
		Params: &mcp.GetPromptParams{
			Name: p.Prompt.Name,
			Arguments: map[string]string{
				"owner":   "samisinada",
				"repo":    "github-mcp-server",
				"summary": "NPPES active; IL 036/336 ACTIVE through 2029-07-31; LEIE no match",
			},
		},
	})
	require.NoError(t, err)
	body := promptText(t, result)
	assert.Contains(t, body, "knowledgebase/credentials.md")
	assert.Contains(t, body, "full Illinois controlled-substance number")
	assert.Contains(t, body, "do not open a noisy issue")
	assert.Contains(t, body, "NPPES active")
}

func TestFamilyMedicinePromptsRegistered(t *testing.T) {
	prompts := AllPrompts(stubTranslation)
	names := map[string]bool{}
	for _, p := range prompts {
		names[p.Prompt.Name] = true
	}
	assert.True(t, names["family_medicine_clinic_ops"])
	assert.True(t, names["family_medicine_qi_protocol"])
	assert.True(t, names["family_medicine_credential_watch"])
}

func TestPromptArgHelpers(t *testing.T) {
	assert.Equal(t, "", promptArg(nil, "owner"))
	assert.Equal(t, "", promptArg(&mcp.GetPromptRequest{}, "owner"))
	req := &mcp.GetPromptRequest{Params: &mcp.GetPromptParams{Arguments: map[string]string{"owner": "  acme  "}}}
	assert.Equal(t, "acme", promptArg(req, "owner"))
	assert.Equal(t, "fallback", defaultPromptArg(&mcp.GetPromptRequest{}, "labels", "fallback"))
}
