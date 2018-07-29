[![wercker status](https://app.wercker.com/status/9f53c8c0c9f0d4f1b97046d28b280bf5/s/master "wercker status")](https://app.wercker.com/project/byKey/9f53c8c0c9f0d4f1b97046d28b280bf5)

CI / CD on Oracle Managed Kubernetes Cluster - OKE & Wercker 
============================================================

This tutorial shows you how to create a production quality Kubernetes cluster
on Oracle Cloud and then setup full CI/CD (continuous integration/continuous deployment)
into the Kubernetes cluster.

This project is executed using Oracle Kubernetes Engine and Oracle Container Registry Service powering the infrastructure
while Wercker and this Git repo powers the CI/CD pipeline. 

Getting started 
===============
* Clone/ Fork Repo. Please feel free to contact me for any questions. 
* Refer article on Medium for detailed information - https://medium.com/jsonlovesyaml/ci-cd-simplified-with-wercker-kubernetes-on-oracle-cloud-infrastructure-a2332497d36f

Repo Structure
===============
1) Kubernetes Directory Contains the deployment templates and the ingress controller
2) Wercker.yml file contains the information for Build and deployment
3) The main directory contains a go program that exposes a webservice on port 5000

Wercker Custom Env Variables
============================
* These values need to be changed in the wercker account where the workflow is configured. 
* These variables are not pipeline specific but workflow specific and globally accessible across pipelines. 
* The first three env variables are declared specific to Oracle Cloud Container Registry and Mangaed OKE Parallems. 
* Please modify the placeholders for the env variables in wercker.yml and the corresponding Kubernetes templates and the env variables in your wercker account. 

| Name of Env Variable        | Explanation                           |
| ----------------------------|---------------------------------------|
| CONTAINER_REGISTRY_USERNAME | User name for Registry                |
| API_USER_TOKEN              | Authentication for User               |
| TENANCY_NAME                | Name of Tenancy where OKE is hosted   |
| PRIVATE_REGISTRY_PATH       | Enpoint of Private Registry           | 
| REPO_NAME                   | Name of  Repository within registry   |
| APP_NAME                    | Name of Application                   |
| SLACK_URL                   | Url of slack webhook for publication  |
| SLACK_CHANNEL               | Channel ID to post notifications      | 
| SLACK_USERNAME              | Slack User name for auth              |
| SLACK_TOKEN                 | Slack Token for Auth                  | 
| KUBE_ENDPOINT               | Cluster End Point - refr kubeconfig   |
| KUBE_USER_TOKEN             | Token provided in Kubeconfig          |

Feature Set 
==================================
* Added Random Cities Generation while accessing webservice.
* Removed Kubeconfig from repo and configured it as an Env Variable in Wercker
* Added Slack Chatbot Integration 
