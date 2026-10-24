Yes. Since you have **3 days left**, I would change our approach from “learn DSA broadly” to **interview-focused preparation across all three company-name areas**:

1. **Coding — 40%**
2. **Infrastructure design — 30%**
3. **SRE / operations — 25%**
4. **Hiring-manager / behavioral — 5%**

You can use the following as a **master prompt/context in a separate ChatGPT conversation**. It is intentionally written to make the other chat understand your current level and not throw advanced DSA or generic SRE theory at you.

# company-name SRE III — 3-Day Interview Preparation Master Prompt

## Candidate Context

I am Farhaan Shamsee, a Senior DevOps / Platform Engineer with 7–8+ years of experience.

My strongest areas are:

- Kubernetes
- AWS / EKS
- Terraform
- GitOps
- ArgoCD / FluxCD
- CI/CD
- Platform Engineering
- Observability
- Infrastructure automation
- Docker
- Helm
- Prometheus / Grafana / Loki / OpenTelemetry
- Cloud infrastructure
- Incident troubleshooting

My application/software coding experience is comparatively less deep than my infrastructure/platform experience.

I am currently preparing for a **company-name Senior Platform Site Reliability Engineer (SRE III)** interview.

---

# Interview Context

I already completed the first technical round.

The architecture/conceptual portion went well.

Topics discussed included:

- EKS
- ArgoCD vs Helm vs FluxCD
- Terraform vs ArgoCD
- Monitoring / observability

One weakness identified during the interview was coding.

The coding problem involved string manipulation. I struggled to derive the algorithm live, and the interviewer had to help me.

I subsequently passed the round, but HR explicitly indicated that the coding feedback was not very positive and advised me to brush up on coding.

The next stage is an **in-person pairing/interview day in Bangalore**.

Expected sessions:

### Session 1 — Infrastructure Design & Development
Approximately 90 minutes.

### Hiring Manager
Approximately 45 minutes.

### Session 2 — SRE Design / Operations
Approximately 90 minutes.

Therefore, this preparation is NOT just a coding preparation.

---

# My Current Coding Level

I am currently rebuilding my coding fundamentals from the ground up.

Do NOT assume I am already comfortable with DSA.

I want to learn the underlying reasoning before memorizing problems.

We have already covered:

- Function vs loop
- Time complexity
- Space complexity
- Big-O
- O(1), O(n), O(n²)
- Traversal
- State
- Accumulators
- Early exit
- `break` vs `return`
- Two pointers
- Palindrome
- Reverse
- Building a result
- Maps
- Frequency counting
- Seen maps
- First duplicate
- Character frequency
- First non-repeating character

Current language: **Go**

I need explanations that are beginner-friendly but interview-oriented.

---

# Important Teaching Style

Do NOT dump large amounts of DSA theory.

Teach interactively.

For each concept:

1. Explain the mental model.
2. Give a tiny example.
3. Ask me to reason about it.
4. Let me answer.
5. Correct my answer.
6. Give the Go implementation.
7. Ask a slightly harder variation.
8. Explain time and space complexity.

Do not immediately give me the solution unless I am genuinely stuck.

The goal is to develop:

> Problem → Pattern → Algorithm → Code → Complexity

rather than:

> Memorize → Copy code.

Before coding, make me explain the algorithm in plain English.

This is particularly important because my Round 1 weakness was deriving the algorithm live.

---

# Three-Day Constraint

I have only **3 days remaining** before the interview.

Do NOT try to teach me all of DSA.

Prioritize patterns that are realistic for a Senior SRE / Platform Engineer coding interview.

Focus on:

### Coding

1. Arrays / slices
2. Strings
3. Maps / frequency counting
4. Traversal
5. State / accumulator
6. Two pointers
7. Sliding window
8. Stack
9. Queue
10. Binary search
11. Basic sorting concepts
12. Simple recursion / tree traversal if time permits
13. Go syntax and common coding patterns
14. Complexity analysis

Avoid spending significant time on:

- Dynamic programming
- Advanced graph algorithms
- Advanced backtracking
- Competitive-programming tricks
- Very difficult LeetCode problems

Only introduce those if there is clear evidence they are relevant.

---

# INFRASTRUCTURE DESIGN PREPARATION

This is equally important.

My strongest technical background is infrastructure/platform engineering.

Prepare me for senior-level infrastructure design questions involving:

## AWS

Focus heavily on:

- AWS Organizations
- AWS Accounts
- Organizational Units
- SCP
- Resource Control Policies / RCP
- IAM
- IAM roles
- Cross-account access
- IAM Identity Center
- Permission boundaries
- Resource policies
- AWS account vending / account management
- Multi-account architecture
- Centralized logging
- Centralized security
- Guardrails
- AWS Control Tower concepts
- VPC
- Subnets
- Route tables
- Internet Gateway
- NAT Gateway
- VPC endpoints
- Security Groups
- Network ACLs
- Transit Gateway
- DNS / Route 53
- EKS
- IAM Roles for Service Accounts / Pod Identity
- ECR
- RDS
- S3
- CloudWatch
- CloudTrail
- KMS
- Secrets Manager

Pay particular attention to:

> SCP vs IAM policy vs resource policy vs permission boundary

and:

> AWS Organizations / OU / account architecture

and:

> How to securely manage infrastructure across many AWS accounts.

---

# TERRAFORM PREPARATION

Terraform is explicitly important for this interview.

Prepare me deeply on:

- Terraform fundamentals
- Providers
- Resources
- Data sources
- Variables
- Locals
- Outputs
- Modules
- Module design
- `for_each`
- `count`
- Dynamic blocks
- Expressions
- Dependency graph
- Implicit dependencies
- Explicit `depends_on`
- Terraform state
- Remote state
- State locking
- State isolation
- Terraform workspaces
- Import
- Drift
- Plan vs apply
- Refresh
- Lifecycle
- `create_before_destroy`
- `prevent_destroy`
- `ignore_changes`
- Provider aliases
- Multi-account AWS Terraform
- Cross-account providers
- AssumeRole
- Secrets handling
- Terraform CI/CD
- Policy as code
- Security scanning
- Module versioning
- State recovery
- Handling concurrent changes
- Terraform failure scenarios

I need to be able to answer senior-level scenario questions such as:

> "How would you design Terraform for 100 AWS accounts?"

> "How would you manage state?"

> "How would you prevent developers from creating insecure infrastructure?"

> "What happens if Terraform state is corrupted?"

> "How do you handle provider authentication across accounts?"

> "How do you structure reusable modules?"

> "How do you deal with drift?"

> "How do you safely roll out Terraform changes?"

---

# SRE / OPERATIONS PREPARATION

The previous interview exposed a weakness in my monitoring answer.

I discussed:

- HTTP status codes
- latency
- tracing

but missed:

- external dependencies
- database-layer monitoring

Therefore, explicitly train me to think about monitoring across the complete dependency chain.

Use this mental model:

```text
User
 ↓
Load Balancer / Ingress
 ↓
Application
 ↓
Kubernetes
 ↓
Database
 ↓
Cache
 ↓
Message Queue
 ↓
External APIs / Services
 ↓
Cloud Infrastructure
```

For every layer ask:

- What can fail?
- What should I measure?
- What should I alert on?
- What logs do I need?
- What traces do I need?
- What symptoms will the customer see?
- What is the likely root cause?

---

# Observability

Prepare deeply on:

### Metrics

- RED
- USE
- Golden signals
- Request rate
- Error rate
- Latency
- Saturation
- CPU
- Memory
- Disk
- Network
- Queue depth
- Connection pools
- DB performance

### Logs

- Structured logging
- Correlation IDs
- Request IDs
- Tenant IDs
- Log aggregation
- Loki / Elasticsearch / OpenSearch concepts

### Tracing

- Distributed tracing
- Trace IDs
- Span IDs
- Service dependencies
- OpenTelemetry
- Sampling

### Alerting

Teach me to distinguish:

- symptom alerts
- cause alerts
- noisy alerts
- actionable alerts

Focus heavily on:

> Alert on user impact / SLO symptoms rather than every possible infrastructure metric.

---

# SLO / SLI / SLA

Prepare me to confidently explain:

- SLI
- SLO
- SLA
- Error budget
- Availability
- Latency SLOs
- Burn rate
- Multi-tenant SLOs

Example:

> If availability SLO is 99.9%, how much downtime is allowed?

Make me calculate and explain it.

---

# INCIDENT MANAGEMENT

Prepare scenario-based questions.

Examples:

### Scenario 1

> Production latency suddenly increases.

Make me reason through:

```text
Detect
 ↓
Scope
 ↓
Impact
 ↓
Recent changes
 ↓
Application
 ↓
Kubernetes
 ↓
Database
 ↓
Dependencies
 ↓
Mitigation
 ↓
Recovery
 ↓
RCA
 ↓
Prevention
```

### Scenario 2

> EKS workloads suddenly cannot reach an external API.

Make me investigate:

- Pod
- Node
- Security Group
- Network policy
- Route table
- NAT Gateway
- DNS
- VPC endpoint
- External service
- TLS

### Scenario 3

> Application error rate increases after a deployment.

Make me investigate:

- Deployment history
- Git commit
- ArgoCD
- Kubernetes events
- Pod logs
- Metrics
- Traces
- Rollback

### Scenario 4

> Database latency increases.

Make me consider:

- DB CPU
- connections
- connection pool
- locks
- slow queries
- indexes
- storage IOPS
- network
- replication
- cache
- application behavior

Do not let me answer only with generic "check logs and metrics."

---

# KUBERNETES / EKS

Because this is an SRE role, prepare:

- EKS architecture
- Control plane vs worker nodes
- Node groups
- Karpenter
- Scheduling
- Requests / limits
- HPA
- VPA concepts
- Cluster Autoscaler
- Pod disruption budgets
- Readiness / liveness / startup probes
- NetworkPolicy
- Service networking
- Ingress
- DNS
- CoreDNS
- CNI
- IAM
- Pod Identity / IRSA
- Secrets
- ConfigMaps
- RBAC
- StatefulSets
- Deployments
- Jobs / CronJobs
- Persistent volumes
- EBS / EFS
- Cluster upgrades
- Failure scenarios
- Observability

Focus on troubleshooting rather than memorizing definitions.

---

# MULTI-TENANCY

company-name operates a large SaaS platform, so prepare multi-tenant SRE concepts.

Topics:

- Tenant isolation
- Noisy neighbors
- Resource quotas
- Limit ranges
- Network isolation
- IAM isolation
- Namespace isolation
- Data isolation
- Tenant-aware observability
- Tenant-level SLOs
- Rate limiting
- Fair scheduling
- Capacity planning
- Failure isolation
- Blast radius
- Per-tenant metrics
- Per-tenant logs
- Per-tenant tracing

Ask scenario questions like:

> One tenant suddenly consumes 10x the normal resources. How do you prevent it from affecting other tenants?

---

# INFRASTRUCTURE DESIGN FRAMEWORK

Whenever you give me a design question, teach me to answer using:

```text
1. Requirements
2. Constraints
3. Scale
4. High-level architecture
5. Networking
6. Compute
7. Storage
8. Security
9. IAM
10. Observability
11. Reliability
12. Scaling
13. Failure scenarios
14. Disaster recovery
15. Cost
16. Trade-offs
```

Don't let me jump directly into services.

Make me clarify requirements first.

---

# SENIOR-LEVEL EXPECTATION

Do not only ask:

> "What AWS service would you use?"

Ask:

> "Why?"

> "What happens if it fails?"

> "How does it scale?"

> "What is the blast radius?"

> "How would you secure it?"

> "How would you observe it?"

> "How would you recover it?"

> "What would you automate?"

> "What trade-off are you making?"

---

# 3-DAY STUDY STRATEGY

## Day 1

### Coding — ~2 hours

- Frequency maps
- Strings
- Arrays
- Two pointers
- Sliding window
- Stack
- Basic Go implementation

### Infrastructure — ~2 hours

AWS Organizations:

- Accounts
- OU
- SCP
- RCP
- IAM
- Cross-account access
- Multi-account architecture

Terraform:

- State
- Modules
- Provider aliases
- `for_each`
- lifecycle
- drift

### SRE — ~1 hour

Observability:

- Metrics
- Logs
- Traces
- DB
- Dependencies
- SLOs

---

## Day 2

### Coding — ~2 hours

- Mixed problems
- 30–45 minute live coding sessions
- Explain algorithm before coding
- Complexity after every problem

### Infrastructure — ~2 hours

Design scenarios:

- Multi-account AWS platform
- EKS platform
- Terraform platform
- Secure cross-account infrastructure
- Centralized logging/security

### SRE — ~1.5 hours

Incident scenarios:

- High latency
- High error rate
- EKS outage
- DB degradation
- External dependency failure
- Deployment failure

---

## Day 3

Do NOT learn large amounts of new material.

### Morning

Coding mock:

```text
45 minutes
2–3 problems
Explain before coding
```

### Midday

Infrastructure design mock:

```text
45–60 minutes
```

### Afternoon

SRE operations mock:

```text
45–60 minutes
```

### Evening

Review:

- AWS Organizations
- SCP
- RCP
- IAM
- Terraform state
- Terraform modules
- EKS
- Observability
- SLOs
- Incident response
- Multi-tenancy

Then do a final mixed interview simulation.

---

# How You Should Interact With Me

Treat this like an intensive interview coaching session.

Do not just provide notes.

Ask me questions.

Let me answer.

Challenge incomplete answers.

If I give a technically correct but shallow answer, ask:

> "What happens next?"

or:

> "What would you check?"

or:

> "Why?"

or:

> "What is the failure mode?"

Correct me when necessary.

Keep explanations concise enough that I can retain them.

Use diagrams and tables when they genuinely help.

For coding, prefer Go.

For AWS/Terraform/SRE, use real-world senior-engineer scenarios.

---

# Most Important Goal

I do NOT need to become an expert in all of DSA or AWS in 3 days.

I need to be able to demonstrate during the interview that I can:

1. Think logically.
2. Break problems into steps.
3. Write clean, understandable Go.
4. Explain complexity.
5. Design reliable infrastructure.
6. Reason about AWS security and account architecture.
7. Use Terraform safely at scale.
8. Troubleshoot production systems systematically.
9. Think about dependencies, databases, and external systems.
10. Communicate like a senior SRE/platform engineer.

When uncertain between breadth and depth, prioritize **interview performance and reasoning ability** over covering more topics.

### How I suggest we use this

You **don't need to switch away from this conversation**. We can use this context ourselves and keep the coding teaching exactly as we've been doing.

For the next **3 days**, I'd run the preparation roughly like this:

**Coding:** interactive teaching + increasingly realistic problems  
**AWS/Terraform:** senior scenario questions + architecture  
**SRE:** incident scenarios + observability + SLOs

And importantly, **I would not stop the coding fundamentals abruptly**. We're close to finishing the foundation. Once frequency maps are solid, we'll move through the highest-value patterns quickly and then start mixing them into **company-name-style live coding problems**.

For the infrastructure side, your existing AWS/Terraform experience means we should spend much less time on basic EC2/S3 definitions and much more time on **Organizations, SCP/RCP, IAM, multi-account architecture, Terraform state/modules, and failure scenarios**.


Yes. **We have enough time — but not enough time to prepare everything equally.** The right move now is to switch from *learning mode* to **interview-priority mode**.

You have just over 50 hours, with today + Sunday as the real preparation window and Monday mostly for consolidation/travel. Given the JD and what we already know about your Round 1 feedback, I would **not spend the remaining time trying to finish DSA comprehensively**.

### Our priority stack

| Priority | Area | Target |
|---|---|---|
| 🔴 1 | **Coding / Python-Go problem solving** | Become comfortable solving unfamiliar medium/easy problems |
| 🔴 2 | **EKS + Kubernetes troubleshooting/design** | Strong senior-level reasoning |
| 🔴 3 | **AWS + Terraform** | Deep, practical answers |
| 🟠 4 | **SRE / Observability / SLOs / incidents** | Scenario-based answers |
| 🟠 5 | **Linux + Networking** | Troubleshooting fundamentals |
| 🟡 6 | **IAM / Security / SSO/OAuth/SAML** | Interview-level understanding |
| 🟡 7 | Kafka/SQS + RDS/Aurora | Know architecture/troubleshooting basics |
| 🟢 8 | CI/CD/GitOps | Already one of your stronger areas |
| 🟢 9 | Crossplane/KubeVela | Already a differentiator for you |
| 🟢 10 | Behavioral / leadership | Final polish |

---

# How I would use the remaining time

## 🟢 TODAY — finish coding foundation + start infra

### Block 1 — DSA: ~2 hours

We've already covered:

- Complexity
- Traversal
- State
- Early exit
- Two pointers
- Hash maps
- Two Sum
- Sliding window
- Stack
- Monotonic stack

We should **not keep learning DSA endlessly**.

Next:

1. Binary search
2. Queue
3. Sorting/common array patterns
4. 3–4 mixed problems where **I don't tell you the pattern**

That last part is crucial.

You'll have to look at:

> "Here's the problem."

and decide:

> "Is this a map? Two pointers? Window? Stack?"

That's much closer to the actual interview.

---

# 🟠 Then switch hard to Guidewire-specific technical prep

This is where your JD becomes extremely useful.

### EKS/Kubernetes

We need to be able to answer scenarios such as:

> A pod is running but the application isn't reachable. Walk me through troubleshooting.

> EKS nodes are healthy but pods are pending.

> One tenant is consuming excessive resources.

> An application suddenly has high latency.

> A deployment succeeds but traffic returns 503.

> DNS works intermittently.

> How does traffic flow from an internet-facing load balancer to a pod?

> How would you secure workloads in EKS?

We should cover:

```text
EKS
├── Control plane
├── Nodes
├── Pods
├── Deployments
├── Services
├── Ingress
├── CNI
├── DNS
├── RBAC
├── NetworkPolicy
├── Secrets
├── IRSA
└── Observability
```

You already know a lot of this practically. Our job is to make your answers **structured and senior-level**.

---

# 🔴 AWS + Terraform

This deserves a large chunk of tomorrow because the JD explicitly says:

> **Deep experience with AWS**

and:

> **Terraform/Terragrunt**

We'll go deep on:

### AWS

- IAM
- Roles vs policies
- STS
- AssumeRole
- Cross-account access
- VPC
- Subnets
- Route tables
- NAT Gateway
- Internet Gateway
- Security Groups
- NACL
- Load Balancers
- Route 53
- EKS
- RDS/Aurora
- S3
- KMS
- CloudWatch
- CloudTrail

### Terraform

- State
- Remote state
- Locking
- Modules
- `for_each`
- `count`
- `dynamic`
- `depends_on`
- Provider aliases
- Multiple AWS accounts
- Drift
- Import
- Lifecycle
- Secrets
- CI/CD
- Terraform plan/apply failure scenarios

And importantly, we'll do **scenario questions**, not just definitions.

---

# 🟠 SRE — very important

This JD is explicitly SRE rather than simply DevOps.

We need you to be extremely comfortable with:

### Observability

```text
Metrics
Logs
Traces
    ↓
Correlation
    ↓
Diagnosis
```

And the four golden signals:

```text
Latency
Traffic
Errors
Saturation
```

But we also need to fix the weakness from your Round 1.

You previously covered:

- status codes
- latency
- tracing

but missed:

> **external dependencies + database layer**

So we'll specifically drill:

```text
User
 ↓
API
 ↓
Service
 ↓
DB
 ↓
External dependency
```

and ask:

> "Where could latency/error be introduced?"

---

# 🟠 SLO / SLI / SLA

You should be able to explain these naturally:

```text
SLI → what we measure
SLO → target
SLA → contractual commitment
```

Then:

- error budget
- burn rate
- alerting
- availability
- latency objectives

And most importantly:

> **How would you design an SLO for a multi-tenant SaaS API?**

That's exactly the sort of senior SRE reasoning the JD points toward.

---

# 🟠 Linux + Networking

Don't try to become a Linux expert in 50 hours.

We need the **troubleshooting toolkit**:

```text
CPU
Memory
Disk
IO
Processes
Sockets
Connections
DNS
TCP
HTTP
TLS
```

Commands:

```bash
top
ps
free
df
du
iostat
ss
netstat
curl
dig
nslookup
traceroute
tcpdump
journalctl
```

And questions like:

> "The service works locally but not from another pod. What do you check?"

---

# 🟡 Security / Identity

Because the JD specifically mentions:

- SSO
- SAML
- OAuth
- Okta
- IAM
- IRSA
- RBAC
- NetworkPolicy
- Pod Security
- Secrets

We need **architecture-level understanding**, not implementation mastery.

For example:

```text
User
 ↓
SSO / IdP
 ↓
OAuth/SAML
 ↓
Application
 ↓
Authorization
```

And:

```text
Pod
 ↓
IRSA / Pod Identity
 ↓
IAM Role
 ↓
AWS API
```

You should be able to explain those flows clearly.

---

# 🟢 Your existing strengths

Don't waste preparation time relearning things you're already strong in.

You already have strong material around:

- GitOps
- FluxCD
- ArgoCD
- Terraform
- Kubernetes
- AWS
- Prometheus/Grafana/Loki
- OpenTelemetry
- Crossplane
- KubeVela
- platform engineering
- Backstage
- CI/CD

Instead, we'll focus on turning those into **interview answers**.

For example, rather than:

> "I used FluxCD."

we want:

> "What problem were you solving? Why GitOps? What failure modes did you consider? How did reconciliation behave? How did you detect drift? What happened when the cluster couldn't reach Git?"

That's senior-level discussion.

---

# 📅 The actual schedule I'd use

### **Today**

**Now → ~2h**

DSA:

- Binary search
- Queue
- sorting
- mixed problems

**Break**

**~2–3h**

EKS/Kubernetes:

- architecture
- networking
- troubleshooting
- security

**Break**

**~2h**

AWS fundamentals + IAM + VPC

**~1h**

SRE observability/SLO revision

Then **stop**. Sleep.

---

### **Sunday = BIG DAY**

This is our most important preparation day.

**Morning**

AWS deep dive:

```text
IAM
STS
Organizations
cross-account
VPC
EKS
RDS
S3
KMS
CloudWatch
```

**Afternoon**

Terraform:

```text
state
modules
providers
multi-account
drift
import
CI/CD
failure scenarios
```

**Evening**

SRE:

```text
observability
SLO/SLI
incident response
RCA
capacity
multi-tenancy
```

Then:

### 🔥 Mock interview

I'll become the interviewer.

No teaching during the first pass.

I'll give you:

- coding problem
- EKS scenario
- AWS scenario
- Terraform scenario
- SRE incident
- architecture question

You answer.

Then I'll tell you exactly where your answer is strong/weak.

---

# Monday

**NO heavy learning.**

This is important.

Morning:

### Final revision

Your own:

```text
AWS cheat sheet
Terraform cheat sheet
EKS troubleshooting tree
SRE cheat sheet
Coding patterns
```

Then maybe **2–3 coding problems**, not 20.

And before your flight:

> **Stop studying.**

You don't want to arrive in Bangalore mentally exhausted.

---

# Most important change from now

Until now we've been in **learning mode**.

From this point:

### Learn → Practice → Interview

Not:

### Learn → Learn → Learn → Learn

Because you've already demonstrated something important today:

You can learn a pattern and **derive the solution yourself**.

So yes:

## **We have enough time.**

But we're going to be ruthless about prioritization.

And based on the JD, your biggest remaining preparation risk isn't "knowing another 20 DSA problems." It's being asked a **senior SRE scenario** and giving a technically correct but insufficiently structured answer.

That's what we'll attack tomorrow.

**For right now, let's finish Binary Search, then we'll switch into the Guidewire technical material.**


how about this???


## JD
Job Description for Senior Platform Site Reliability Engineer (EKS and Python/Golang)
Job Description Summary
Site Reliability Engineering (SRE) brings together software and systems engineering to design and operate large-scale, highly distributed, and fault-tolerant
systems. As a Site Reliability Engineer III, on the Platform team, you will play a key role in designing, building, and operating the foundational
infrastructure that powers Guidewire's SaaS platform. You will focus on automation, reliability, scalability, and operability — ensuring our multi-tenant
cloud platform consistently meets both functional and non-functional requirements. You will work closely with product development teams to influence
system design, improve resilience, and ensure production readiness, while building the tools and frameworks that enable efficient, global, follow-the-sun
operations.
Job Description
What You'll Do
Drive Reliability, Automation & Scale:
Design, build, and operate highly reliable, scalable infrastructure for a multi-tenant SaaS platform.
Automate deployment, provisioning, and operational workflows across cloud infrastructure and applications.
Develop internal tools, services, and frameworks to improve efficiency and reduce manual effort.
Participate in a 24x7 follow-the-sun on-call rotation to support critical production systems.
Improve Platform & Infrastructure:
Contribute to core platform systems by building features, resolving issues, and enhancing reliability.
Partner with development teams to ensure systems meet availability, performance, and scalability requirements.
Proactively identify risks, bottlenecks, and failure modes, and implement solutions before they impact customers.
Observability, Incident Management & Resilience:
Build and maintain observability systems (metrics, logging, tracing, dashboards).
Define and track Service Level Objectives (SLOs) and reliability metrics.
Lead or contribute to incident response, root cause analysis, and blameless postmortems.
Drive improvements toward self-healing systems and reduced operational toil.
Security & Identity:
Design and support secure access patterns, including SSO, SAML, and OAuth-based authentication systems.
Ensure platform services meet security and compliance standards.
Enablement & Collaboration:
Collaborate across engineering teams, providing guidance, feedback, and hands-on contributions.
Create and maintain documentation, runbooks, and training materials.
Mentor engineers and promote best practices in reliability engineering and automation.
Who You Are
Infrastructure Development:
8–12 years of hands-on experience in Site Reliability Engineering (SRE), DevOps, Cloud Infrastructure, or a related platform engineering role, with
a proven track record of designing, building, and operating highly available, scalable, and reliable production systems.
Strong programming skills in Python or Go (Java/Spring Boot is a plus).
Deep experience with AWS and building/operating production systems at scale.
Hands-on expertise with Kubernetes (EKS), Docker, Helm, CNI, and Ingress networking.
Strong understanding of Kubernetes primitives and patterns (deployments, services, operators, etc.).
Experience with Infrastructure as Code (Terraform, Terragrunt, or similar).
Solid understanding of Linux systems and networking fundamentals.
Observability & Operations:
Experience with observability platforms such as Datadog, Prometheus, OpenTelemetry, or CloudWatch.
Familiarity with incident management practices and production support in a microservices environment.
Experience with messaging/streaming systems (e.g., Kafka, SQS) and relational databases (e.g., Aurora, RDS) is a plus.
Security & Identity:
Working knowledge of SSO, SAML, OAuth, and identity providers (Okta is a plus).
Experience with AWS IAM (roles, policies, IRSA), VPC security groups, and Kubernetes security primitives (RBAC, network policies, pod
security standards, secrets management).
DevOps & Delivery:
Experience with CI/CD and GitOps tools such as GitHub Actions, TeamCity, Jenkins, FluxCD, or Bitbucket.
Comfortable working in agile environments (Scrum, Kanban).
Mindset & Collaboration:
Strong troubleshooting and problem-solving skills with a proactive, systems-thinking mindset.
Passion for automation: “If you have to do it more than once, automate it.”
Excellent communication skills and ability to work across distributed teams.
A collaborative team player who can influence, mentor, and lead through technical expertise.
Demonstrated ability to leverage AI and data-driven insights to improve productivity and outcomes.
Preferred Qualifications
Bachelor's degree in Computer Science or related field, or equivalent experience
Experience supporting large-scale SaaS platforms
AWS or Kubernetes certifications
Exposure to modern platform frameworks such as KubeVela (OAM) or Crossplane
Contributions to open-source projects
Why Guidewire?
Work on a mission-critical global platform used by leading P&C insurers worldwide
Solve complex, real-world infrastructure problems at genuine scale
Be part of a collaborative, high-impact engineering culture grounded in integrity, rationality, and collegiality
Opportunity to shape the future of a rapidly evolving cloud platform
A culture of curiosity and innovation where engineers are empowered to leverage AI and emerging technologies.