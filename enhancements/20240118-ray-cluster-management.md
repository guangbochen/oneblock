# Ray Cluster Management

## Summary

This proposal aims to add support for Ray cluster management in OneBlock.

### Related Issues

- https://github.com/oneblock-ai/oneblock/issues/17

## Motivation

### Goals
- add Ray cluster management support
- add Ray cluster autoscaling support
- add built-in GCS fault tolerant support using Redis
- add built-in gang scheduling support
- add built-in Ray cluster monitoring & logging support

### Non-goals [optional]
- N/A

## Proposal
There are three types of users in regards to Ray cluster management:
- Cluster admin: manages all Ray clusters in the k8s cluster
- Ray cluster admin(Namespace owner): manages the Ray cluster in his namespace
- Ray user(Namespace user): submits Ray jobs to the Ray cluster in his namespace


### User Stories

#### Story 1
As a cluster admin, I want to be able to mange all Ray clusters in my k8s cluster, where I can:
- List all Ray clusters
- Create, update, delete Ray clusters
- Manage redis service for GCS fault tolerance
- Manage gang scheduler configs

#### Story 2
As a Ray cluster admin, I want to be able to manage my Ray cluster, where I can:
- View and edit Ray cluster configs
  - autoscaling configs
  - head node configs(enable/disable GCS)
  - worker node configs
- Select the gang scheduler by its allowed namespaces
- View Ray cluster status and check its monitoring & logging info

#### Story 3
As a Ray cluster user, I want to be able to:
- Submit Ray jobs to the Ray cluster
- View Ray jobs status and check its monitoring & logging info

### Implementation Details

Add `kuberay-operator` as a built-in system component in OneBlock. The `kuberay-operator` will be deployed as a Deployment in the `oneblock-system` namespace by default and helps to list/watch the Ray clusters.

#### API changes
- CRDs introduced by the [KubeRay](https://github.com/ray-project/kuberay), and we will be focusing on the `RayCluster` in this proposal:
  - `RayCluster`: KubeRay fully manages the lifecycle of RayCluster, including cluster creation/deletion, autoscaling, and ensuring fault tolerance.
  - `RayJob`: With RayJob, KubeRay automatically creates a RayCluster and submits a job when the cluster is ready. You can also configure RayJob to automatically delete the RayCluster once the job finishes.
  - `RayService`: RayService is made up of two parts: a RayCluster and a Ray Serve deployment graph. RayService offers zero-downtime upgrades for RayCluster and high availability.
- CRDs introduced by the [Volcano](https://volcano.sh/en/docs/):
  - `Queue`: Queue is a collection of PodGroups, which adopts FIFO. It is also used as the basis for resource division. Queue is a cluster-scope resource.
  - `PodGroup`: PodGroup is a group of pods with strong association and is mainly used in batch scheduling, for example, ps and worker tasks in TensorFlow. PodGroup is a namespace-scope resource.

## Design

### RayCluster

## Gang Scheduling

Gang scheduling is a scheduling policy that allows a group of pods to be scheduled together or not at all.

## GCS fault tolerance

Global Control Service (GCS) manages cluster-level metadata. By default, the GCS lacks fault tolerance as it stores all data in-memory, and a failure can cause the entire Ray cluster to fail. To make the GCS fault tolerant, you must have a high-availability Redis. This way, in the event of a GCS restart, it retrieves all the data from the Redis instance and resumes its regular functions.

#### Redis

Redis will be deployed as a built-in service in OneBlock. The Redis service will be deployed as a StatefulSet with 3 replicas. The Redis service will be deployed in the `oneblokc-system` namespace by default.
- The password is auto generated and stored in the `oneblock-system` secret.
- You can access the Redis service using redis-cli within the master redis pod.

### Implementation Overview


### Test plan

### Upgrade strategy

N/A

## Note

N/A
