# Privateer Plugin osps-baseline

This wireframe is designed to quickly get your service pack repository up to speed!

Privateer's plugin architecture relies on some key elements being present at the top level
of the service pack, all of which are provided along with examle code in this repo.

Simply fork or clone this repo and start adjusting the tests to build your own service pack!

Based on the Open Source Project Security (OSPS) Baseline, here is a consolidated checklist of all rules across various categories, including descriptions of what each baseline implies to check.

# OSPS Rules Implementation Status – Combined with Current Status

| Category | Rule | Description | Status (GraphQL) | Status (REST) | Implemented |
|-------------------------|------------|-----------------------------------------------------------------------------------------------------------------------|--------------------|--------------------|-------------|
| **Access Control (AC)** | OSPS-AC-01 | Require multi-factor authentication for collaborators modifying repository settings or accessing sensitive data. | ✅ Implemented | N/A | true |
| | OSPS-AC-02 | Restrict collaborator permissions to the lowest available privileges by default. | ❌ Not Implemented | N/A | false |
| | OSPS-AC-03 | Prevent unintentional direct commits against the primary branch. | ✅ Implemented | N/A | true |
| | OSPS-AC-04 | Prevent unintentional deletion of the primary branch. | ✅ Implemented | N/A | true |
| | OSPS-AC-05 | Configure project CI/CD pipeline permissions to the lowest available privileges. | ❌ Not Implemented | N/A | false |
| | OSPS-AC-07 | Require multi-factor authentication (non-SMS) for users modifying repository settings or accessing sensitive data. | ❌ Not Implemented | ❌ Not Implemented | false |
| **Build and Release (BR)** | OSPS-BR-01 | Do not execute arbitrary code from outside the build script. | ❌ Not Implemented | N/A | false |
| | OSPS-BR-02 | Assign a unique version identifier to each release. | ✅ Implemented | ✅ Implemented | true |
| | OSPS-BR-03 | Use encrypted channels (SSH, HTTPS) for websites, API responses, and other project development and release services. | N/A | ✅ Implemented | true |
| | OSPS-BR-04 | Create all released software assets with consistent, automated build and release pipelines. | ❌ Not Implemented | N/A | false |
| | OSPS-BR-05 | Use standardized tooling to ingest dependencies at build time. | ❌ Not Implemented | N/A | false |
| | OSPS-BR-06 | Provide a descriptive change log for all releases. | ✅ Implemented | N/A | true |
| | OSPS-BR-08 | Sign or account for all released software assets in a signed manifest including cryptographic hashes. | N/A | ✅ Implemented | true |
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
