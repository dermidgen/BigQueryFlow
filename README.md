# BigQueryFlow
BigQuery Job/Query Execution Flows - Create data pipelines in BigQuery!

## Goals
 * Lowest possible cost at rest
 * Declarative pipelines
 * REST API interface
 * DAG Based 
 * Event driven
 * Saga support

## Proposed Design
BigQueryFlow seeks to provide the lowest cost footprint possible by using Cloud Functions, Pub/Sub and Stackdriver integrations. The `bqflow` CLI allows you to apply configuration changes, which will interactively provision the cloud functions required to execute your flow.

### Setup/Installation
In order to get Pub/Sub events about BigQuery Jobs, we need to perform an some initial setup of Stackdriver. Google may expose native BigQuery events in the future, but until then the accepted workaround is to tweak Stackdriver integration.

### DAG by Digraph
DOT format, supporting label/query based selectors and support standard BigQuery job options.

```dot
digraph {
  node [selector[label, queryName], jobOptions] -> node [selector[label, queryName], jobOptions]
}
```

### Flows by YAML/JSON
```yaml
kind: BigQueryFlow
metadata:
  name: myflow
spec:
  selector:
    matchLabels:
      name: myflow
  template:
    metadata:
      labels:
        name: myflow
    spec:
      dags:
        - name: mydag
          valueFrom:
            dagRef: mydag
      endpoint:
```

### Usage
```bash
bqflow create dag mydag --from-dot ./mydag.dot
bqflow create flow myflow
bqflow apply -f ./myflow.yml
bgflow get dag mydag
bqflow get flow myflow
bqflow logs -f myflow
```
