CI / CD on Oracle Managed Kubernetes Cluster - OKE & Wercker 
============================================================

This tutorial shows you how to create a production quality Kubernetes cluster
on Oracle Cloud and then setup full CI/CD (continuous integration/continuous deployment)
into the Kubernetes cluster.

This project is executed using Oracle Kubernetes Engine and Oracle Container Registry Service powering the infrastructure
while Wercker and this Git repo powers the CI/CD pipeline. 

Repo Structure
===============
1) Kubernetes Directory Contains the deployment templates and the ingress controller
2) Wercker.yml file contains the information for Build and deployment
3) The main directory contains a go program that exposes a webservice on port 5000


Feature Set 
==================================
Added Random Cities Generation while accessing webservice
Removed Kubeconfig from repo and configured it as an Env Variable in Wercker
Added Slack Chatbot Integration 
