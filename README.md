
# Project 1 — Local Café Network System

Goal:
Build a usable real-world product.

Focus:
- captive portal
- local hosting
- QR onboarding
- real-time dashboards
- offline-first design

Tech:
- React PWA
- Go backend
- SQLite
- hostapd
- dnsmasq
- WebSockets

You learn:
- networking basics
- local infra
- distributed UI updates
- edge deployment

Portfolio signal:
“Built local-first hotspot-based business system.”

---

Yes. That is the correct approach.

Build the core platform first, then abstract deployment modes later.

Recommended architecture:

```text id="jlwm1r"
                CafeOS Core
        ┌───────────────────────┐
        │ Go Backend            │
        │ React Frontend        │
        │ SQLite                │
        │ WebSocket Engine      │
        │ Auth/RBAC             │
        └───────────────────────┘
                   ↓
        Deployment Adapters
   ┌──────────┬──────────┬──────────┐
   │ Mobile   │ PC       │ Edge     │
   │ Mode     │ Mode     │ Mode     │
   └──────────┴──────────┴──────────┘
```

Build order:

# Phase 1 — Core Backend

Implement:

* auth
* menu management
* ordering
* billing
* roles
* websocket updates

Do NOT think about networking yet.

---

# Phase 2 — React Frontend

Build:

* customer UI
* waiter UI
* kitchen UI
* admin dashboard

Make it:

* responsive
* PWA-ready

---

# Phase 3 — Local Deployment

Run:

* Go server locally
* SQLite locally
* local Wi-Fi access

Test:

* multiple phones on same LAN

---

# Phase 4 — Mobile Hosting Mode

Experiment with:

* Android hotspot
* Go server on Android
* local access through hotspot IP

This itself becomes impressive.

---

# Phase 5 — Captive Portal Integration

Add:

* QR onboarding
* local DNS redirect
* auto-open flows

---

# Phase 6 — Advanced Networking

Only later:

* VLANs
* firewall rules
* identity networking
* LANOS concepts

Important architectural principle:

Everything should work even without:

* captive portals
* VLANs
* hotspot control

Those should become optional infrastructure layers.

This keeps the system portable and much easier to evolve.
