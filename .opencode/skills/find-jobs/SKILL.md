---
name: find-jobs
description: This skill is dedicated for finding and listing job offers in IT. Use when the user asks to find jobs.
---

# How to find jobs
You should use MCP server "czyjesteldorado" to find jobs.

## Instructions
1. Find job offers for a Go developer, seniority declared in the offer should be: Mid or Senior or Principal or Architect.
2. Filter out offers typical for Frontend and DevOps and SRE and Management.
3. Include following information in the results: offer title, company, required technologies, salary (if provided), location (city), work mode (remote/hybrid), link to the offer.
4. Format and print offers as a table in Markdown format.

## Expected outcome
A table with all relevant job offers is printed out.

## Example
User request: "znajdź prace" or "find job" or "find jobs"

Agent behavior:
- Contact MCP server "czyjesteldorado"  with phrase "Go" to fetch current job offers for Golang dvelopers
- Filter out offers unrelated to Go like frontend, DevOps, SRE, Management
- Collect information about: offer title, company, required technologies, salary, location (city), work mode (remote/hybrid/office), link to the offer.
- Format the results as a table in Markdown, keep original offer language (don't translate)

## Expected output format - Markdown table, first column is item ordinal number, '-' means salary not provided in the offer
# Golang Job Offers 15.04.2026

| No. | Title | Company | Technologies | Salary (PLN) | Location | Work Mode | Link |
|---|-------|---------|--------------|--------------|----------|-----------|------|
| 1. | Senior Engineer (Go) | Ericsson | C, Golang, Java, Python, LTE/4G/5G | 20000-25000 | Kraków, Łódź | hybrid | [Link](https://czyjesteldorado.pl/praca/326686-senior-engineer-c-or-go-ericsson) |
| 2. | Senior Backend Engineer (Ruby and/or Go), Tenant Scale | GitLab | Backend, Ruby, Go, GitLab, Security, AI, Architecture | - | - | remote | [Link](https://czyjesteldorado.pl/praca/300805-senior-backend-engineer-ruby-and-or-go-tenant-scale-cells-infrastructure-gitlab) |
| 3. | DevOps Engineer | ConnectPoint | Kubernetes, Terraform, Ansible, CI/CD, Python, Golang, PostgreSQL, SQL Server, Azure | 15000-20000 | Warszawa | hybrid, remote | [Link](https://czyjesteldorado.pl/praca/315328-devops-engineer-connectpoint) |
| 4. | Golang Architect | ITEAMLY | Golang, SQL, Python, Java, Pub/Sub, Kafka, NewRelic, DataDog | 26000-36000 | Kraków | remote | [Link](https://czyjesteldorado.pl/praca/323872-golang-architect-iteamly-spolka-z-ograniczona-odpowiedzialnoscia) |
| 5. | Senior Software Engineer II (Golang, Order) | SPOTON POLAND | Golang, PHP, React, JavaScript, AWS, Terraform | 23800-29800 | Kraków | hybrid | [Link](https://czyjesteldorado.pl/praca/316861-senior-software-engineer-ii-golang-order-spoton-poland-spolka-z-ograniczona-odpowiedzialnoscia) |
