package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/github/github-mcp-server/pkg/inventory"
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const familyMedicinePHIBoundary = "Never write patient names, DOB, MRN, addresses, phone numbers, visit notes, or other identifiers into GitHub issues, pull requests, comments, or the knowledgebase. The EHR is the system of record. If the user request contains PHI, stop and rewrite it as a de-identified process or protocol issue, or refuse to file it."

func promptArg(request *mcp.GetPromptRequest, name string) string {
	if request == nil || request.Params == nil || request.Params.Arguments == nil {
		return ""
	}
	return strings.TrimSpace(request.Params.Arguments[name])
}

func defaultPromptArg(request *mcp.GetPromptRequest, name, fallback string) string {
	if v := promptArg(request, name); v != "" {
		return v
	}
	return fallback
}

// FamilyMedicineClinicOpsPrompt turns a clinic-operations request into a de-identified GitHub issue.
func FamilyMedicineClinicOpsPrompt(t translations.TranslationHelperFunc) inventory.ServerPrompt {
	return inventory.NewServerPrompt(
		ToolsetMetadataIssues,
		mcp.Prompt{
			Name:        "family_medicine_clinic_ops",
			Description: t("PROMPT_FAMILY_MEDICINE_CLINIC_OPS_DESCRIPTION", "File a de-identified FQHC family-medicine clinic-operations issue (rooming, inbox, refill rules, referrals) without PHI"),
			Arguments: []*mcp.PromptArgument{
				{
					Name:        "owner",
					Description: "Repository owner",
					Required:    true,
				},
				{
					Name:        "repo",
					Description: "Repository name",
					Required:    true,
				},
				{
					Name:        "area",
					Description: "Workflow area: huddle, rooming, inbox, refills, referrals, standing-orders, or other",
					Required:    true,
				},
				{
					Name:        "need",
					Description: "What is broken or missing in the clinic process (no patient identifiers)",
					Required:    true,
				},
				{
					Name:        "labels",
					Description: "Comma-separated labels (optional; defaults include clinic-ops and no-phi)",
					Required:    false,
				},
			},
		},
		func(_ context.Context, request *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			owner := promptArg(request, "owner")
			repo := promptArg(request, "repo")
			area := promptArg(request, "area")
			need := promptArg(request, "need")
			labels := defaultPromptArg(request, "labels", "clinic-ops,no-phi")

			messages := []*mcp.PromptMessage{
				{
					Role:    "user",
					Content: &mcp.TextContent{Text: "You are a family-medicine clinic operations assistant using GitHub issues. " + familyMedicinePHIBoundary + " Follow knowledgebase/family-medicine-workflows.md. Search for duplicates before creating an issue."},
				},
				{
					Role:    "user",
					Content: &mcp.TextContent{Text: fmt.Sprintf("Open or update a clinic-ops issue in %s/%s for area %q. Need: %s\nLabels: %s\nIssue body must include: problem (process only), expected clinic behavior, acceptance check, and owner role (MA, RN, clinician, or ops). Do not include any patient identifiers.", owner, repo, area, need, labels)},
				},
				{
					Role:    "assistant",
					Content: &mcp.TextContent{Text: fmt.Sprintf("I will search %s/%s for an existing %s clinic-ops issue, refuse if the need looks like PHI, then create or update a de-identified issue with labels %s.", owner, repo, area, labels)},
				},
			}
			return &mcp.GetPromptResult{Messages: messages}, nil
		},
	)
}

// FamilyMedicineQIProtocolPrompt turns a protocol or QI change into an issue and a follow-on PR.
func FamilyMedicineQIProtocolPrompt(t translations.TranslationHelperFunc) inventory.ServerPrompt {
	return inventory.NewServerPrompt(
		ToolsetMetadataIssues,
		mcp.Prompt{
			Name:        "family_medicine_qi_protocol",
			Description: t("PROMPT_FAMILY_MEDICINE_QI_PROTOCOL_DESCRIPTION", "Turn a family-medicine protocol or QI defect into a GitHub issue and a documentation pull request, without PHI"),
			Arguments: []*mcp.PromptArgument{
				{
					Name:        "owner",
					Description: "Repository owner",
					Required:    true,
				},
				{
					Name:        "repo",
					Description: "Repository name",
					Required:    true,
				},
				{
					Name:        "title",
					Description: "Protocol or QI issue title",
					Required:    true,
				},
				{
					Name:        "change",
					Description: "What should change in the standing order, template, or measure definition",
					Required:    true,
				},
				{
					Name:        "labels",
					Description: "Comma-separated labels (optional; defaults include qi,protocol,no-phi)",
					Required:    false,
				},
			},
		},
		func(_ context.Context, request *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			owner := promptArg(request, "owner")
			repo := promptArg(request, "repo")
			title := promptArg(request, "title")
			change := promptArg(request, "change")
			labels := defaultPromptArg(request, "labels", "qi,protocol,no-phi")

			messages := []*mcp.PromptMessage{
				{
					Role:    "user",
					Content: &mcp.TextContent{Text: "You are a family-medicine QI assistant. " + familyMedicinePHIBoundary + " Protocol PRs update knowledgebase/family-medicine-workflows.md or a linked protocol note. Patient-specific exceptions do not belong in GitHub."},
				},
				{
					Role:    "user",
					Content: &mcp.TextContent{Text: fmt.Sprintf("In %s/%s, create issue %q and then a PR that documents the protocol change.\nChange: %s\nLabels: %s\nIssue and PR must include: what changes, who may act, inclusion/exclusion, when to stop and get a clinician, EHR location of the live order set, and a review date.", owner, repo, title, change, labels)},
				},
				{
					Role:    "assistant",
					Content: &mcp.TextContent{Text: fmt.Sprintf("I will search %s/%s for a duplicate QI issue, create %q if needed, then open a documentation PR that records the standing-order or measure change without PHI.", owner, repo, title)},
				},
			}
			return &mcp.GetPromptResult{Messages: messages}, nil
		},
	)
}

// FamilyMedicineCredentialWatchPrompt compares public credential sources to the local knowledgebase.
func FamilyMedicineCredentialWatchPrompt(t translations.TranslationHelperFunc) inventory.ServerPrompt {
	return inventory.NewServerPrompt(
		ToolsetMetadataIssues,
		mcp.Prompt{
			Name:        "family_medicine_credential_watch",
			Description: t("PROMPT_FAMILY_MEDICINE_CREDENTIAL_WATCH_DESCRIPTION", "Compare public credential sources to knowledgebase/credentials.md and open or update a credentials issue when something changed"),
			Arguments: []*mcp.PromptArgument{
				{
					Name:        "owner",
					Description: "Repository owner",
					Required:    true,
				},
				{
					Name:        "repo",
					Description: "Repository name",
					Required:    true,
				},
				{
					Name:        "summary",
					Description: "What the latest public-source refresh found (status, dates, unmatched LEIE). No secrets, DEA, or full controlled-substance number",
					Required:    true,
				},
			},
		},
		func(_ context.Context, request *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			owner := promptArg(request, "owner")
			repo := promptArg(request, "repo")
			summary := promptArg(request, "summary")

			messages := []*mcp.PromptMessage{
				{
					Role:    "user",
					Content: &mcp.TextContent{Text: "You maintain public professional credentials for a family physician. Read knowledgebase/credentials.md and knowledgebase/changelog.md. " + familyMedicinePHIBoundary + " Never store DEA numbers, passwords, tokens, DOB, SSN, or the full Illinois controlled-substance number. Leave IDFPR UI status unconfirmed unless a human submitted License Lookup."},
				},
				{
					Role:    "user",
					Content: &mcp.TextContent{Text: fmt.Sprintf("Latest public-source refresh for %s/%s:\n%s\nIf any watched value changed, a renewal window opened, or OIG/LEIE status changed, search for an existing credentials issue and create or update one. If nothing material changed, say so and do not open a noisy issue.", owner, repo, summary)},
				},
				{
					Role:    "assistant",
					Content: &mcp.TextContent{Text: fmt.Sprintf("I will compare that summary to knowledgebase/credentials.md in %s/%s, update changelog only if a value or action changed, and file a credentials issue only for a real delta.", owner, repo)},
				},
			}
			return &mcp.GetPromptResult{Messages: messages}, nil
		},
	)
}
