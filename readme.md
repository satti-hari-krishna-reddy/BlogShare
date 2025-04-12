# BlogShare

> A unified platform to publish and share blog content across multiple social media platforms with AI-generated posts. Built for developers, creators, and content teams who want to amplify their reach without repetitive manual work.

---

## 🚀 What is BlogShare?

**BlogShare** is a smart publishing tool that helps users:

- **Connect** their blog (e.g. Hashnode, Dev.to, Medium) and social media accounts (Twitter, LinkedIn, etc.)
- **Share** blog posts automatically or manually with a custom message.
- **Generate** high-quality social share messages using AI.

It's like turning a single blog post into a multichannel marketing weapon with one click.

---

## 🧠 Core Features

### 1. **Platform Connections**
- OAuth2 integration with social media platforms.
- Blog APIs (Hashnode, Dev.to, etc.) for fetching published articles.

### 2. **AI-Powered Messaging**
- Uses OpenAI under the hood.
- Smart prompt engineering to match tone & platform context (e.g. tweets are punchy, LinkedIn posts more professional).

### 3. **One-Click Sharing**
- Select platforms.
- Preview generated content.
- Click share — boom, it’s live.

### 4. **User Dashboard**
- View blog history.
- Manage connected accounts.
- Check share status per platform.

### 5. **Dev-Friendly Backend**
- Clean Docker-based setup.
- Azure Functions for async tasks.
- CI/CD with GitHub Actions.

---

## 📸 Visual Architecture

> _(Refer to the attached images in the repo for the complete system design and workflows.)_

- `architecture.png`: Full app system design.
- `workflow.png`: End-to-end blog-to-social flow.
- `ci-pipeline.png`: GitHub CI/CD flow.

---

## ⚙️ Tech Stack

| Layer        | Tech Used                         |
|--------------|-----------------------------------|
| Frontend     | React + Tailwind CSS              |
| Backend      | Go + Azure Functions |
| AI Service   | Gemini API                        |
| Cache         | Redis                     |
| DB           | MongoDB                            |
| Deployment   | Docker + Azure Cloud      |
| CI/CD        | GitHub Actions                    |

---

## 🧭 Workflow

### 1. **User connects blog & socials**
- OAuth tokens are securely stored.

### 2. **Blog is fetched automatically**
- Periodic jobs or manual sync.

### 3. **AI generates a message**
- Blog metadata passed to Gemini AI.
- Platform-specific tone.

### 4. **User reviews & customizes**
- Can edit or regenerate content.

### 5. **Share is triggered**
- Each platform is called via its API.
- Posting status is tracked.

### 6. **Dashboard reflects results**
- Logs, previews, post links.

---

## 📦 Local Setup

```bash
git clone git@github.com:satti-hari-krishna-reddy/BlogShare.git
cd BlogShare
docker-compose up --build
```

Visit: `http://localhost:3000`

> Make sure to set `.env` file correctly. Template is available as `.env.example`.

---

## 📈 Future Roadmap

- [ ] Analytics dashboard
- [ ] Scheduling shares
- [ ] Custom AI tone templates
- [ ] Instagram & Facebook support
- [ ] Team collaboration features

---

## 👨‍💻 Author

Built with 💻 & ☕ by [@satti-hari-krishna-reddy](https://github.com/satti-hari-krishna-reddy)

---

## ⭐️ Show some love

If you find this useful, feel free to star the repo and share it! Let’s automate the boring parts of content creation 🔥

