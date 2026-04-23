# Project Overview

## Project name
Multi-Tenant SaaS Backend Platform

## Purpose
This project demonstrates a backend service for managing organizations, users, and project memberships in a multi-tenant SaaS environment.

## What problem it solves
Multi-tenant systems need clear separation between organizations while still running on shared infrastructure. This project models that pattern with a simple but realistic API.

## Core capabilities
- user registration
- JWT-based login
- organization creation and listing
- project membership creation and listing
- optional organization manifest upload to S3/R2-compatible storage

## Target use cases
- SaaS backend prototype
- auth and organization management demo
- portfolio project for backend or platform roles

## Primary technology stack
- Go
- Echo framework
- PostgreSQL
- GORM
- JWT
- Docker / Docker Compose
- AWS S3 / Cloudflare R2 compatible object storage
