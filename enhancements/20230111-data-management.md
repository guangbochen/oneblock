# Data Processing & Management

## Summary

Data pre-processing & management is a key component in ML. This enhancement is to provide a way to allow users to upload files to OneBlock and process the data in the file to generate datasets.

### Related Issues

https://github.com/oneblock-ai/oneblock/issues/6

## Motivation

Data pre-processing is usually done via a data pipeline. OneBlock platform should provide a user-friendly way to allow users to create a scalable data pipeline to process the data in the file to generate datasets.

### Goals

- Allowing users to upload files to OneBlock and process the data in the file to generate datasets.
- Allowing users to upload files using remote url:
  - HTTP/HTTPs URL
  - NFS endpoint with path.
  - S3 endpoint with bucket name and path.
- List and view generated datasets.
- Be able to modify and labeling the datasets.
- Support multiple file formats(e.g., parquet, images, text, cvs, binary and TFRecords).
- Users should be able to select and store the dataset to different destinations(i.e., s3-compatible storage and VectorDB(e.g., Qdrant)).

### Non-goals [optional]

- online data batch inference.

What is out of scope for this enhancement? Listing non-goals helps to focus discussion and make progress.

## Proposal

This is where we get down to the nitty gritty of what the proposal actually is.

### User Stories
Detail the things that people will be able to do if this enhancement is implemented. A good practise is including a comparsion of what user cannot do before the enhancement implemented, why user would want an enhancement and what user need to do after, to make it clear why the enhancement beneficial to the user.

The experience details should be in the `User Experience In Detail` later.

#### Story 1
#### Story 2

### User Experience In Detail

Detail what user need to do to use this enhancement. Include as much detail as possible so that people can understand the "how" of the system. The goal here is to make this feel real for users without getting bogged down.

### API changes

A Dataset is a distributed data collection for data loading and processing.

```YAML
apiVersion: core.oneblock.ai/v1alpha1
kind: DataSet
metadata:
  name: pdf-data-sample
  namespace: default
spec:
  source:
    type: text # options are csv, parquet, image, text, binary, tfrecords
    url: https://oneblock.ai/sample.pdf
  destination:
    type: s3 # options are local or cloud storage, nfs 
    format: parquet # options are parquet, image, text, binary, tfrecords
    url: s3://oneblock-ai/dataset
  entrypoint: python /home/ray/samples/sample_code.py
  image: 
    name: oneblockai/ray-data:main-head
    pullPolicy: Always
  rayJobConfig:
    distributeOps: "map" # options are map, filter, repartition
    parallelism: 2
    gpu: 0
    batch_size: 100
    scheduling: "DEFAULT" # support DEFAULT or SPREAD
  rayClusterConfig:
```

## Design

### Implementation Overview

Overview on how the enhancement will be implemented.

### Test plan

Integration test plan.

### Upgrade strategy

Anything that requires if user want to upgrade to this enhancement

## Note [optional]

Additional nodes.
