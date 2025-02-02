# Event Trigger Platform

A Golang-based API to create and manage event triggers (scheduled & API-based).

## 🚀 Features
- Create, Edit, Delete, and Execute Triggers
- Schedule events using cron
- API triggers with JSON payloads
- Event logs with 48-hour lifecycle
- Redis caching for logs
- **Dockerized & Deployed on Fly.io**

---

## 🛠 Tech Stack
- **Golang + Gin** (API)
- **PostgreSQL** (Database)
- **Redis** (Caching)
- **Docker & Fly.io** (Deployment)

---

## 🛠 Local Setup

### 1️⃣ Prerequisites
- Install **Docker**: [Docker Install Guide](https://docs.docker.com/get-docker/)
- Install **Docker Compose**

### 2️⃣ Clone Repository
```sh
git clone https://github.com/shivamsouravjha/segwise.git
cd segwise
```

### 3️⃣ Run with Docker
```sh
docker-compose up --build
```

The API will run on `http://localhost:4000`.

---

## 📡 API Endpoints

### **1️⃣ Create a Scheduled Trigger**
```sh
curl -X POST "http://localhost:4000/api/triggers"      -H "Content-Type: application/json"      -d '{
           "type": "scheduled",
           "schedule": "*/10 * * * *", 
           "endpoint": null,
           "payload": null
         }'
```

### **2️⃣ Create an API Trigger**
```sh
curl -X POST "http://localhost:4000/api/triggers"      -H "Content-Type: application/json"      -d '{
           "type": "api",
           "schedule": null,
           "endpoint": "http://example.com/webhook",
           "payload": "{ \"message\": \"Hello, world!\" }"
         }'
```

### **3️⃣ Execute a Trigger Manually**
```sh
curl -X POST "http://localhost:4000/api/triggers/{trigger_id}/execute"
```

### **4️⃣ Get All Triggers**
```sh
curl -X GET "http://localhost:4000/api/triggers"
```

### **5️⃣ Delete a Trigger**
```sh
curl -X DELETE "http://localhost:4000/api/triggers/{trigger_id}"
```

---

## 🚀 Deployment (Fly.io)

The project is deployed at:  
🔗 **[Event Trigger App](https://event-trigger-app-bitter-bird-4607.fly.dev/swagger/index.html)**

---

## 💰 Cost Estimation
Fly.io Free Tier Usage:

* Compute (256MB RAM, 1 vCPU)	Free
* PostgreSQL (LiteDB)	Free
* Redis (Upstash)	Free
* Bandwidth (5GB limit)	Free
* Estimated Cost for 30 Days (24x7, 5 queries/day): $0 (Free Tier)

