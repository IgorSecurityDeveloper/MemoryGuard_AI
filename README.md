# MemoryGuard_AI

# MemoryGuard AI

**MemoryGuard AI** is a security-oriented monitoring project designed to collect, process, and analyze memory-related events from distributed environments. The project is being developed as a hands-on study and portfolio initiative focused on **Support Engineering, Cybersecurity, backend systems, observability, and event-driven architecture**.

The main goal of MemoryGuard AI is to simulate a modern monitoring and security pipeline where a lightweight agent collects system memory metrics, sends them to a messaging platform, and allows a backend service to consume, process, and eventually store or analyze these events.

The current architecture includes a **Go-based collector**, responsible for gathering memory usage data, and a **Java Quarkus backend**, designed to consume and process events. Communication between services is handled through **Apache Kafka**, while **PostgreSQL** is used as the relational database layer. The environment is containerized with **Docker Compose**, making it easier to reproduce, test, and evolve the system locally.

## Project Purpose

MemoryGuard AI was created to explore how cybersecurity, troubleshooting, and backend engineering can work together in a practical project. It is not intended to be a finished commercial product at this stage, but rather a growing technical foundation for learning and demonstrating skills related to:

* Event-driven systems
* System monitoring
* Memory usage collection
* Kafka-based communication
* Java backend development with Quarkus
* Go-based lightweight agents
* Dockerized development environments
* Technical troubleshooting and infrastructure setup
* Security-oriented architecture

## Current Architecture

The current project flow is:

```text
Go Collector → Kafka Topic → Quarkus API → PostgreSQL
```

The collector publishes memory usage events to a Kafka topic called `memory-events`. The backend API is responsible for consuming these events and will later be extended to persist, analyze, and expose the collected data through APIs and dashboards.

## Technologies Used

* **Java 21**
* **Quarkus**
* **Go**
* **Apache Kafka**
* **Zookeeper**
* **PostgreSQL**
* **PgAdmin**
* **Docker**
* **Docker Compose**
* **WSL / Linux development environment**

## Roadmap

Planned improvements include:

* Persisting memory events in PostgreSQL
* Creating REST endpoints to query collected data
* Adding basic alert rules for abnormal memory usage
* Building a frontend dashboard
* Improving collector capabilities with CPU, process, and disk metrics
* Adding authentication and tenant-aware architecture
* Exploring anomaly detection and AI-assisted analysis
* Improving observability with logs, metrics, and dashboards

## Professional Context

This project reflects my interest in combining **Support Engineering**, **Cybersecurity**, **troubleshooting**, and **software development**. It is part of my learning path toward building stronger skills in secure backend systems, incident-oriented monitoring, cloud-ready architecture, and real-world infrastructure troubleshooting.
