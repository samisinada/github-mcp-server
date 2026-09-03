# Family medicine clinical workflows

Operational playbooks for an FQHC family-medicine panel. Use GitHub for
protocol, staffing, QI, and credential work only. Keep the EHR as the
system of record for patients.

Suggested issue labels: `clinic-ops`, `qi`, `protocol`, `credentials`,
`no-phi`.

## 1. Start-of-session guard

Before any GitHub write:

1. Confirm the work is process or protocol, not a patient.
2. Search existing issues (`search_issues`) so the same standing-order or
   inbox-rule change is not opened twice.
3. If the request contains a name, DOB, MRN, or visit detail, stop and
   rewrite it as a de-identified process issue, or do not file it.

## 2. Daily clinic flow (Helping Hands / FQHC)

Run in the EHR and huddle. File a GitHub issue only when a *system*
breaks (template missing, standing order outdated, referral queue rule
wrong).

| Slot | Owner | Done when |
|---|---|---|
| Huddle | MA + clinician | High-risk outreach, interpreter needs, and same-day slots are named without writing identifiers into chat logs or issues |
| Rooming | MA | Vitals, med rec start, PHQ-2/9 and substance screen per standing order, gap list visible |
| Visit | Clinician | Agenda set, chronic + preventive work bundled, after-visit summary in EHR |
| Close | Clinician + MA | Orders signed, refill rules applied, referrals routed, inbox emptied of today |

Common system issues to file (templates, not patients):

- Rooming template missing PHQ-2, tobacco, or BP recheck.
- Huddle list cannot show care gaps without exporting PHI.
- Same-day slot rules block hospital-follow-up or newborn visits.

## 3. Preventive care (bundle in the visit)

Close gaps in the room. Track *measure definitions* and *template
defects* here; track *who is due* only in the EHR/population tool.

Priority adult bundles (confirm current USPSTF/ACIP/org protocol before
changing a standing order):

- BP, tobacco, BMI, depression, intimate-partner safety as indicated
- Cervical, colorectal, breast, lung screening by age/risk
- Diabetes / lipid / HIV / HCV / immunization status
- Prenatal and pediatric periodicity when the clinic session includes
  those panels

QI issue pattern: one measure, one defect, one owner, one acceptance
check (for example “Pap standing order fires for ages 21–65 with no Pap
in 3 years and no hysterectomy flag”).

## 4. Chronic disease panels

Do not export panel CSVs into this repo. File issues against the
*protocol* or *registry rule*.

| Condition | Visit bundle | Escalation (process, not a chart) |
|---|---|---|
| Hypertension | Confirm home vs office technique; med rec; labs per protocol | Issue if BP recheck or home-BP order set is missing |
| Type 2 diabetes | A1c, urine albumin, statin/ACE discussion, foot/eye flags | Issue if registry cannot show last A1c without a manual hunt |
| Asthma / COPD | Control screen, inhaler technique, action-plan template | Issue if action-plan smartphrase is stale |
| Depression / anxiety | PHQ-9 / GAD-7, safety, follow-up interval | Issue if positive PHQ-9 has no warm-handoff path documented as a clinic rule |
| Heart failure / CAD | GDMT checklist, weight, renal labs | Issue if the order set omits a required lab |

## 5. Inbox, refills, and results

Triage in the EHR. Promote a GitHub issue only for a broken rule.

Refill default (clinic policy to encode in the EHR, not in GitHub):

- Chronic med, stable, seen within the protocol window, no controlled
  substance → refill + schedule
- Controlled substance → in-person or video per clinic policy; never
  decide from a GitHub issue
- Abnormal result needing a new plan → schedule, do not “refill around”
  it

Inbox SLA to keep in clinic ops (adjust to org policy):

- Critical results: same day
- Abnormal non-critical: 2 business days
- Routine patient messages: 2 business days
- Referral updates: 5 business days

## 6. Referrals and care coordination

Track *queues and templates*, not named patients.

A healthy referral issue looks like: “GI open-access order set still
asks for a paper fax cover that our FQHC abandoned.”

A prohibited issue looks like: “Please call Ms. X about her colonoscopy.”

## 7. Standing orders and protocols

Use `family_medicine_qi_protocol` so a protocol change is an issue with
acceptance criteria, then a PR against the markdown or order-set note in
this repo.

Required sections on every protocol PR:

1. What changes
2. Who may act (MA, RN, clinician)
3. Inclusion / exclusion
4. When to stop and get a clinician
5. EHR location of the live order set
6. Review date

## 8. Credential and payer watch

Use `family_medicine_credential_watch` after each public-source refresh.

Watch items (public only):

- IL 036 / 336 status and expiration
- NPPES address vs CMS Doctors address (keep both; do not merge)
- PECOS enrollment and Order/Referring flags
- OIG LEIE absence

Open or update a `credentials` issue when a value changes or a renewal
window starts (typically 90 days before IL expiration). Do not paste
DEA or the full CS number into the issue.

## 9. Prompt → tool map

| Prompt | First tools | Stop if |
|---|---|---|
| `family_medicine_clinic_ops` | `search_issues`, then `issue_write` / granular create | Request contains PHI |
| `family_medicine_qi_protocol` | `search_issues`, `issue_write`, then `create_pull_request` after a branch exists | Change is a patient-specific exception |
| `family_medicine_credential_watch` | Read `knowledgebase/credentials.md`, `search_issues` | Issue would need a secret or full CS number |
