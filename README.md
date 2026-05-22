# ⚽ World Cup 2026 - Betting Pool App

A complete World Cup 2026 betting pool application featuring live match updates, real-time rankings, private groups, and a fully mobile-first experience.

---

# ✨ Features

## 🔐 Authentication

* JWT authentication
* Login and registration flow
* Validation using Zod + PrimeVue Forms
* Route guards
* Smooth SPA transitions

## 👥 Groups

* Create private groups
* Join groups using invite codes
* Multiple groups per user
* Group-based rankings

## ⚽ Matches

* Automatic match import via external API
* Live match auto refresh
* Automatic scroll to highlighted match
* Skeleton loading states
* Stage filters
* Smart refresh system

## 🎯 Predictions

* Match predictions
* Real-time updates
* Stage-based scoring system
* Custom stage weights
* Prediction points display

## 🏆 Rankings

* Group rankings
* Best user position on dashboard
* Prediction count
* Total points

## 📱 Mobile First

* iPhone safe area support
* Fixed bottom navigation
* Smooth transitions
* Sports-app inspired UX
* Pinch zoom disabled
* Real-time visual feedback

---

# 🧠 Scoring System

| Event                      | Points |
| -------------------------- | ------ |
| Correct winner             | 3 pts  |
| Correct draw               | 3 pts  |
| Correct score for one team | 1 pt   |
| Exact score prediction     | 5 pts  |

Competition stages use different multipliers.

Example:

| Stage          | Weight |
| -------------- | ------ |
| Group Stage    | 1x     |
| Round of 16    | 2x     |
| Quarter-finals | 3x     |
| Semi-finals    | 4x     |
| Final          | 5x     |

---

# 🏗️ Architecture

## Backend

* Golang
* Clean Architecture
* PostgreSQL
* SQLC
* JWT Authentication
* Cronjobs with urfave/cli

## Frontend

* Vue 3
* PrimeVue
* TailwindCSS
* Pinia
* Vue Router
* Motion animations

---

# 📂 Project Structure

```txt
internal/
 ├── app/
 │   ├── domain/
 │   │ └── usecase/
 │   ├── gateway/
 │   ├── handler/
 │   └── config/
 │
frontend/
 ├── src/
 │   ├── components/
 │   ├── pages/
 │   ├── stores/
 │   ├── services/
 │   └── router/
```

---

# ⚙️ Cronjobs

## Import Matches

```bash
go run cmd/cronjob/main.go import-matches
```

## Calculate Prediction Points

```bash
go run cmd/cronjob/main.go calculate-prediction-points
```

---

# 🚀 Running the Project

## Backend

```bash
go mod tidy
go run cmd/api/main.go
```

## Frontend

```bash
npm install
npm run dev
```

---

# 🎨 UI/UX

## Fonts

* Outfit
* Rajdhani

## Design

* Glassmorphism
* Mobile-first
* App-like experience
* Smooth animations
* Live score transitions

---

# 📸 Main Features

✅ Smart dashboard

✅ Real-time rankings

✅ Automatic live match updates

✅ Group system

✅ Invite codes

✅ Live scores

✅ Mobile-friendly predictions

✅ Automatic match scrolling

✅ Individual card auto refresh

---

# 🔮 Roadmap

* PWA support
* Push notifications
* Global rankings
* Admin panel
* Advanced statistics
* Prediction sharing
* WebSocket live updates

---

# 👨‍💻 Author

Developed by Leonardo Muller.

Created for the 2026 FIFA World Cup with a strong focus on performance, mobile UX, and scalable architecture.
