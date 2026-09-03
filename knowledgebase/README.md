# Family medicine knowledge base

Public-directory credentials and clinic-operations playbooks for this
GitHub MCP Server fork. Nothing here is a patient record.

## What lives here

| File | Purpose |
|---|---|
| [credentials.md](credentials.md) | Public professional credentials (NPPES, Illinois open data, CMS, OIG) |
| [changelog.md](changelog.md) | Daily refresh log — rewrite only when a value or action changes |
| [family-medicine-workflows.md](family-medicine-workflows.md) | FQHC family-medicine operational playbooks mapped to GitHub issues/PRs |

## PHI boundary

Do not put names, DOB, MRN, visit notes, lab values, addresses, or phone
numbers that identify a patient into issues, pull requests, comments, or
this folder. Track work as process, protocol, or de-identified QI only.

## MCP prompts

The `issues` toolset registers three prompts that turn these playbooks
into GitHub work:

- `family_medicine_clinic_ops`
- `family_medicine_qi_protocol`
- `family_medicine_credential_watch`

## Credential refresh rules

Daily cron re-queries NPPES, CMS Doctors and Clinicians (`mj5m-pzi6`),
Illinois Professional Licensing (`pzzh-kp68`), CMS PPEF, Order and
Referring, and OIG LEIE. Official IDFPR License Lookup UI stays
`unconfirmed` unless a human actually submits the form.

Never store DEA numbers, passwords, tokens, DOB, SSN, or the full
Illinois controlled-substance number.
