# Privateer Plugin for GitHub Repos

This wireframe is designed to quickly get your service pack repository up to speed!

Privateer's plugin architecture relies on some key elements being present at the top level
of the service pack, all of which are provided along with examle code in this repo.

Simply fork or clone this repo and start adjusting the tests to build your own service pack!

Based on the Open Source Project Security (OSPS) Baseline, here is a consolidated checklist of all rules across various categories, including descriptions of what each baseline implies to check.

# OSPS Rules Implementation Status – Combined with Current Status

| Category | Rule | Description | Status (GraphQL) | Status (REST) | Implemented |
|-------------------------|------------|-----------------------------------------------------------------------------------------------------------------------|--------------------|--------------------|-------------|
| **Access Control (AC)** | OSPS-AC-01 | The project's version control system MUST require multi-factor authentication for collaborators modifying the project repository settings or accessing sensitive data. | ✅ | N/A | true |
| | OSPS-AC-02 | The project's version control system MUST restrict collaborator permissions to the lowest available privileges by default. | ✅ | N/A | true |
| | OSPS-AC-03 | The project's version control system MUST prevent unintentional direct commits against the primary branch. | ✅ | N/A | true |
| | OSPS-AC-04 | The project's version control system MUST prevent unintentional deletion of the primary branch. | ✅ | N/A | true |
| | OSPS-AC-05 | The project's permissions in CI/CD pipelines MUST be configured to the lowest available privileges except when explicitly elevated. | ✅ | N/A | false |
| | OSPS-AC-07 | The project's version control system MUST require multi-factor authentication that does not include SMS for users when modifying the project repository settings or accessing sensitive data. | N/A | N/A | false |
| **Build and Release (BR)** | OSPS-BR-01 | The project’s build and release pipelines MUST NOT permit untrusted input that allows access to privileged resources. | ❌ Not Implemented | N/A | false |
| | OSPS-BR-02 | All releases and released software assets MUST be assigned a unique version identifier for each release intended to be used by users. | ✅ Implemented | ✅ Implemented | true |
| | OSPS-BR-03 | Any websites and version control systems involved in the project development MUST be delivered using SSH, HTTPS, or other encrypted channels. | N/A | ✅ Implemented | true |
| | OSPS-BR-04 | All released software assets MUST be created with consistent, automated build and release pipelines. | ❌ Not Implemented | N/A | false |
| | OSPS-BR-05 | All build and release pipelines MUST use standardized tooling where available to ingest dependencies at build time. | ❌ Not Implemented | N/A | false |
| | OSPS-BR-06 | All releases MUST provide a descriptive log of functional and security modifications. | ✅ Implemented | N/A | true |
| | OSPS-BR-08 | All released software assets MUST be signed or accounted for in a signed manifest including each asset’s cryptographic hashes. | N/A | ✅ Implemented | true |
| | OSPS-BR-09 | Any websites or other services involved in the distribution of released software assets MUST be delivered using HTTPS or other encrypted channels. | N/A | ❌ Not Implemented | false |
| | OSPS-BR-10 | Any websites, API responses or other services involved in release pipelines MUST be fetched using SSH, HTTPS or other encrypted channels. | N/A | ❌ Not Implemented | false |
| **Documentation (DO)** | OSPS-DO-03 | Provide user guides for all basic functionalities. | N/A | ✅ Implemented | true |
| | OSPS-DO-05 | Include a mechanism for reporting defects in project documentation. | ✅ Implemented | N/A | true |
| | OSPS-DO-12 | Include instructions to verify the integrity and authenticity of release assets, including the expected signer identity. | ❌ Not Implemented | N/A | false |
| | OSPS-DO-13 | Include a statement about the scope and duration of support. | ❌ Not Implemented | N/A | false |
| | OSPS-DO-14 | Describe when releases or versions are no longer supported and won't receive security updates. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-DO-15 | Describe how the project selects, obtains, and tracks dependencies. | ❌ Not Implemented | N/A | false |
| **Governance (GV)** | OSPS-GV-01 | Include roles and responsibilities for project members in the documentation. | ❌ Not Implemented | N/A | false |
| | OSPS-GV-02 | Establish one or more mechanisms for public discussions about proposed changes and usage obstacles. | ✅ Implemented | N/A | true |
| | OSPS-GV-03 | Include an explanation of the contribution process in the documentation. | ✅ Implemented | N/A | true |
| | OSPS-GV-04 | Provide a guide for code contributors outlining requirements for acceptable contributions. | ❌ Not Implemented | N/A | false |
| | OSPS-GV-05 | Have a policy requiring code contributor review before granting escalated permissions to sensitive resources. | ❌ Not Implemented | ❌ Not Implemented | false |
| **Legal (LE)** | OSPS-LE-01 | Require all code contributors to assert they are legally authorized to commit their contributions. | ✅ Implemented | ❌ Not Implemented | true |
| | OSPS-LE-02 | Use a license that meets the OSI Open Source Definition or the FSF Free Software Definition. | ✅ Implemented | N/A | true |
| | OSPS-LE-03 | Maintain the source code license in a standard location within the project's repository. | ✅ Implemented | N/A | true |
| | OSPS-LE-04 | Ensure the released software assets use a license that meets the OSI Open Source Definition or the FSF Free Software Definition. | ❌ Not Implemented | N/A | false |
| **Quality (QA)** | OSPS-QA-01 | Make the project's source code publicly readable with a static URL. | N/A | ✅ Implemented | true |
| | OSPS-QA-02 | Maintain a publicly readable commit history with author and timestamp information. | ❌ Not Implemented | N/A | false |
| | OSPS-QA-03 | Deliver all released software assets with a machine-readable list of dependencies and their versions. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-QA-04 | Require all automated status checks for commits to pass or require manual acknowledgement before merge. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-QA-05 | Enforce security requirements on additional subproject code repositories as applicable. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-QA-06 | Do not store generated executable artifacts in version control. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-QA-08 | Use at least one automated test suite with clear documentation on when and how tests are run. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-QA-09 | Include a policy that major changes add or update tests in an automated test suite. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-QA-10 | Require at least one non-author approval of changes before merging into the release or primary branch. | ❌ Not Implemented | ❌ Not Implemented | false |
| **Security Assessment (SA)** | OSPS-SA-01 | Provide design documentation demonstrating all actions and actors within the system. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-SA-02 | Include descriptions of all external input and output interfaces of the released software assets. | ❌ Not Implemented | N/A | false |
| | OSPS-SA-03 | Perform threat modeling and attack surface analysis. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-SA-04 | Perform a security assessment to understand potential security problems. | ❌ Not Implemented | ❌ Not Implemented | false |
| **Vulnerability Management (VM)** | OSPS-VM-01 | The project MUST use automated vulnerability scanning tools on a regular basis. | ❌ Not Implemented | ❌ Not Implemented | false |
| | OSPS-VM-02 | The project MUST have a documented vulnerability response plan. | ❌ Not Implemented | N/A | false |
| | OSPS-VM-03 | The project MUST track and document any identified vulnerabilities using CVE identifiers where applicable. | ❌ Not Implemented | N/A | false |
| | OSPS-VM-04 | The project MUST publish security advisories for fixed vulnerabilities. | ❌ Not Implemented | N/A | false |

For detailed information on each rule, refer to the [Open Source Project Security Baseline](https://eddieknight.dev/security-baseline/). 
