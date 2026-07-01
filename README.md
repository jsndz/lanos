# LANOS

LANOS is a modular local-first application platform for physical spaces.

It provides networking, identity, communication, access control, and application infrastructure as a reusable core that can be extended through modules.

Instead of building separate applications for cafés, hotels, campuses, events, and other environments, developers build modules on top of a shared platform.

---

## What LANOS Provides

### Identity & Access

* Authentication
* User management
* Session management
* Role-based access control (RBAC)
* Device identification
* Permission management

---

### Network Infrastructure

* QR onboarding
* Wi-Fi onboarding
* Captive portal management
* Local DNS routing
* Network policies
* VLAN assignment
* Device isolation
* Firewall policies

---

### Communication

* Event bus
* WebSocket infrastructure
* Real-time messaging
* Notifications
* Broadcast channels

---

### Application Runtime

* Module system
* Service registration
* Configuration management
* Shared storage
* Lifecycle management
* SDK for module development

---

### Local-First Operation

* Offline operation
* Local application hosting
* Edge deployment
* Synchronization engine
* Local data storage

---

## User Flow

```text
Scan QR
    ↓
Connect to Network
    ↓
Captive Portal
    ↓
Authentication
    ↓
Role Assignment
    ↓
Network Policies Applied
    ↓
Module Launched
```

---

## Architecture

```text
+-----------------------------------+
|             Modules               |
|-----------------------------------|
| Cafe | Hotel | Retail | Campus    |
+-----------------------------------+
                 ↑
+-----------------------------------+
|          LANOS Runtime            |
|-----------------------------------|
| Module Loader                     |
| Service Registry                  |
| Event System                      |
| Synchronization                   |
+-----------------------------------+
                 ↑
+-----------------------------------+
|      Identity & Access Layer      |
|-----------------------------------|
| Auth                              |
| Users                             |
| Sessions                          |
| RBAC                              |
+-----------------------------------+
                 ↑
+-----------------------------------+
|        Network Control Layer      |
|-----------------------------------|
| QR Onboarding                     |
| Captive Portal                    |
| VLAN Management                   |
| Firewall Policies                 |
| DNS Routing                       |
+-----------------------------------+
                 ↑
+-----------------------------------+
|         Physical Network          |
|-----------------------------------|
| Wi-Fi | Ethernet | Switches       |
+-----------------------------------+
```

---

## Module Philosophy

The LANOS core contains no business logic.

Business functionality is implemented through modules.

Examples:

* Cafe Module
* Hotel Module
* Retail Module
* Event Module
* Campus Module
* IoT Module

Modules consume platform services instead of implementing infrastructure themselves.

---

## Example

A Café module should not implement:

* Login
* Sessions
* Device management
* WebSockets
* Event infrastructure
* Network onboarding

Those capabilities are already provided by LANOS.

The module only implements:

* Menu management
* Ordering
* Tables
* Kitchen workflow

---

## Goals

* Local-first operation
* Network-aware applications
* Modular architecture
* Offline capability
* Event-driven communication
* Real-time interactions
* Dynamic network control
* Reusable platform infrastructure

---

## Vision

LANOS is a platform where applications, users, devices, and networks operate as a single system.

Developers build business modules while LANOS provides identity, networking, communication, and runtime infrastructure.

**Tagline**

A local-first operating system for physical spaces.
