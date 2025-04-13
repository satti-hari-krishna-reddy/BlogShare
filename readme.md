# BlogShare

> Automate, Summarize, and Schedule Your Blogs to X and LinkedIn with One Click — Powered by Gemini

---

## 🚀 What Problem Does BlogShare Solve?

A lot of individuals, especially developers and bloggers, have a larger following on **social media** (like X/Twitter, LinkedIn) than on traditional blogging platforms like Hashnode.

After writing a blog, people often manually go to each platform, write a small summary/glimpse, and post the blog link. This repetitive manual effort is:
- **Time-consuming**
- **Inconsistent** across platforms
- **Misses optimal post timing**

**BlogShare** solves this by:
- **Fetching** your blog post automatically.
- **Generating a social-ready summary** using AI (Gemini).
- **Letting you share it with a single click** to both X and LinkedIn.
- **Scheduling** your shares with a centralized UI.

No more copy-pasting, rewording, and switching tabs. Write once. Share smart.

---

## 🧠 Core Features

### 1. **Platform Connections**
- OAuth2 integration with X (Twitter), LinkedIn
- API key-based integration with Hashnode (Blog source)

### 2. **AI-Powered Summaries**
- Uses **Gemini API** under the hood
- Tailored tone depending on platform: short for X, professional for LinkedIn

### 3. **One-Click Multi-Platform Sharing**
- Select where to share
- Hit Share. Done.

### 4. **Centralized Scheduling Engine**
- Built with heap-based priority queue and Go routines
- Schedule now or later

### 5. **User Dashboard**
- Shows shared blogs, scheduled blogs
---

## 📸 Architecture Overview

![Cloud Architecture](https://github.com/user-attachments/assets/162dc992-f23d-471b-ac4a-b98275d4363e)

### 🔁 Scheduling Engine

- Built using **Go routines + heap-based priority queue**
- Worker-Agent architecture
- Handles retries, backoffs, and concurrent posting

![Scheduler Architecture](https://github.com/user-attachments/assets/5f320c4d-68c7-446c-b929-adce099fbddb)

### 🔄 CI/CD Pipeline (GitHub Actions)

![CI/CD Diagram](https://github.com/user-attachments/assets/3b83b2ff-5a02-4b12-80dd-802ca4e86909)

- **Monorepo setup** with `/frontend` and `/backend`
- Conditional CD: deploys only changed part
- Auto-builds Docker images and pushes to registry
- Azure App Services pull and auto-restart

---

## 🛡 Security & Infra

| Aspect        | Detail                                              |
|---------------|------------------------------------------------------|
| CSRF/XSS      | Sanitized inputs, anti-CSRF tokens, CORS control    |
| Auth          | OAuth2 for social, OTP via Mailgun for email verify |
| Rate Limiting | Custom IP-based limiters to prevent abuse           |
| Cloudflare    | Protects from DDoS, adds caching/CDN layer          |
| Secrets       | Managed via GitHub Actions + Azure Environment Variables     |

---

## ⚙️ Tech Stack

| Layer        | Tech Used                         |
|--------------|-----------------------------------|
| Frontend     | React + Tailwind CSS              |
| Backend      | Go + Azure Functions              |
| AI Service   | Gemini API                        |
| Cache        | Redis                             |
| DB           | MongoDB                           |
| Deployment   | Docker + Azure Cloud              |
| CI/CD        | GitHub Actions                    |

---

## 🧭 Workflow in Action

### 1. **User connects blog & socials**
- Hashnode API Key, X and LinkedIn OAuth

### 2. **Blog is fetched automatically**
- Polling or on-demand fetch

### 3. **AI generates content**
- Gemini API creates platform-specific summaries

### 4. **User reviews or edits content**
- One-click approve/share

### 5. **Scheduling**
- Queue-based execution with timestamp-based triggers

### 6. **Share is triggered**
- Handled via platform-specific APIs

### 7. **Dashboard reflects state**
- Success/failure, preview, and post links

---

## 📦 Local Setup

```bash
git clone git@github.com:satti-hari-krishna-reddy/BlogShare.git
cd BlogShare
docker-compose up
```

Visit: `http://localhost:3000`

> Ensure `.env` is configured.

---

## 📈 Roadmap

- [ ] Add Webhook support
- [ ] More platform support (Medium, Facebook)
- [ ] Custom AI tone templates

---

## 👨‍💻 Built With

Crafted with 💻 + ☕ by [@satti-hari-krishna-reddy](https://github.com/satti-hari-krishna-reddy)

> Trying to solve a real-world problem with sharp tech. If you’re hiring for backend or full-stack roles — let’s chat 😉


---

## ⭐️ Show some love

If you find this useful, feel free to star the repo and share it! Let’s automate the boring parts of content creation 🔥

