# Credential knowledgebase changelog

## 2026-09-03 — daily refresh, no material change

Sources queried:

- NPPES API v2.1 `number=1205568904` — active; last update 2026-03-31; Kedzie address unchanged.
- CMS Doctors and Clinicians `mj5m-pzi6` (catalog modified 2026-07-31) — PAC `4486130093`, enrollment `I20260114001217`, Erie Family Health Center Inc at 1701 W Superior St; 1 row; ind/grp assignment Y.
- CMS PPEF public enrollment API (dataset modified 2026-07-27) — same PAC/enrollment; IL family practice `14-08`; multiple-NPI flag N.
- CMS Order and Referring latest file `OrderReferring_20260831.csv` (dataset modified 2026-09-01) — PARTB/DME/HHA/PMD/HOSPICE all Y.
- Illinois Professional Licensing `pzzh-kp68` (rows still updated 2026-09-02 09:42 UTC) — 036 and 336 ACTIVE through 2029-07-31; temp permit 125 cancelled.
- IDFPR License Lookup UI — CAPTCHA present; UI status unconfirmed.
- OIG LEIE `UPDATED.csv` 2026-08-10 — no NPI or last-name SINADA match.

Material changes vs 2026-09-02 snapshot: none.

This run also added family-medicine clinic-ops playbooks and MCP prompts so credential watch and QI work stay in GitHub without PHI.

## 2026-09-02 — first snapshot

Sources queried:

- NPPES API v2.1 `number=1205568904` — active; last update 2026-03-31; Kedzie address unchanged.
- CMS Doctors and Clinicians `mj5m-pzi6` — PAC `4486130093`, enrollment `I20260114001217`, Erie Family Health Center Inc at 1701 W Superior St.
- CMS PPEF latest extract 2026-07-17 — same PAC/enrollment; IL family practice.
- CMS Order and Referring 2026-08-31 — order/refer flags all Y.
- Illinois Professional Licensing `pzzh-kp68` (updated 2026-09-02) — 036 and 336 ACTIVE through 2029-07-31; 2026 renewal posted; temp permit 125 cancelled.
- IDFPR License Lookup UI — CAPTCHA (AWS WAF); UI status unconfirmed.
- OIG LEIE `UPDATED.csv` 2026-08-10 — no match.

Material changes vs prior memory notes:

- IL 036/336 renewal confirmed in open data (was an open action).
- IL 036 expiration now 2029-07-31; effective 2026-05-16.
- IL 336 public redaction `33*****72`; last modified 2026-07-26; schedules II–V.
- Historical temp permit 125.080634 recorded as cancelled.
- Second public practice address added from CMS Doctors (Superior St).
