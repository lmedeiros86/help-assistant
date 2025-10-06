# Help Assistant 

# 🤖 AI-Powered Customer Support Assistant (MVP)

This project is a **Minimum Viable Product (MVP)** of an AI-powered diagnostic tool, built to directly align with the mission of the New Relic Help Experience team.

The core goal is to provide **context-aware, reliable answers** to customer technical troubleshooting questions by using **Retrieval-Augmented Generation (RAG)**, which grounds the LLM's response in a curated knowledge base.

***

## ✨ Technology Stack Showcase

This project demonstrates fluency across the required stack, with clear separation of concerns:

| Component | Technology | Job Requirement Alignment |
| :--- | :--- | :--- |
| **Frontend (UI)** | **React** / **TypeScript** | *Fluency in TypeScript and React* for a production-ready UI. |
| **Backend (API Gateway)** | **Go** (`net/http`) | *Experience in backend services/API development* in Go; handles orchestration and core SQL logging. |
| **AI/RAG Service** | **Python** (LangChain, OpenAI) | *Experience with LLMs, RAG, and machine learning* (Bonus Point). |
| **Data Persistence** | **SQLite** (Go Driver) | *Database Fluency (SQL, data hygiene)*. Architected for production swap to AWS RDS/PostgreSQL. |

***

## 💡 Architectural Decisions and Role Alignment

The design choices reflect a customer-focused mindset and efficient engineering:

* **Polyglot Architecture:** Using **Go** for the API gateway and **Python** for the RAG service demonstrates the ability to manage cross-service communication and choose the best language for the task.
* **Retrieval-Augmented Generation (RAG):** RAG is implemented to retrieve relevant documents from the Vector Database, ensuring the diagnostic advice is fact-based, which is critical for driving **customer resolution success**.
* **Database Fluency:** **SQLite** was chosen for zero-config, rapid prototyping, but the structured SQL schema proves knowledge of **data hygiene** and scalability paths (to AWS RDS/PostgreSQL).

***

## 🚀 Local Setup Instructions

### Prerequisites

* **Go** (1.18+)
* **Node.js** & **npm**
* **Python 3.10+**
* An **OpenAI API Key**

### 1. Configuration & Dependencies

Create a `.env` file in the **`knowledge-base`** directory, and add your API key:
