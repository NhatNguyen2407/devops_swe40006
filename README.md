# DevOps SWE40006 - AWS Workshop

This repository contains a small static web application and the container/deployment configuration used for an AWS-oriented DevOps exercise.

## What's included

- `index.html` - a standalone, styled workshop page
- `AWS_logo.png` - image displayed by the page
- `Dockerfile` - packages the page in an Nginx container
- `buildspec.yml` - AWS CodeBuild commands for building and pushing the image to Amazon ECR
- `taskdef.json` - Amazon ECS Fargate task definition with a container listening on port 80
- `appspec.yaml` - CodeDeploy ECS service deployment template

## Run locally with Docker

From the repository root:

```bash
docker build -t devops-workshop .
docker run --rm -p 8080:80 devops-workshop
```

Open http://localhost:8080 to view the web page. The site itself is static and does not require a database or application backend.

## AWS deployment configuration

The checked-in build specification logs in to Amazon ECR, builds and tags a container image, pushes it, and exports `imageDetail.json` together with the ECS deployment templates. The task definition requests Fargate with 256 CPU units and 512 MiB of memory. `appspec.yaml` maps the `webapp` container to port 80.

The configuration contains AWS account-specific repository and role references. These are examples tied to the original workshop setup, not reusable defaults. Anyone deploying in another account must update the ECR repository, task roles, ECS resources, permissions, and deployment settings.

This repository demonstrates the intended pipeline configuration. The presence of those files does not prove that a cloud deployment is currently running or that a full pipeline execution has succeeded.

## Attribution

This is workshop coursework. I keep the application source and its deployment definitions together so I can revisit the build and deployment workflow.
