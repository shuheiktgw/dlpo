# dlpo
dlpo is a lightweight CLI tool that makes Terraform plan output from the `datadog_logs_pipeline_order` resource more human-readable.

## Installation

```shell
go install github.com/shuheiktgw/dlpo@latest
```

## Prerequisites
- (Optional) pbpaste – Recommended for easier usage.
- (Required) `DD_API_KEY`, `DD_APP_KEY`, and `DD_SITE` environment variables. 

## Usage

Copy the Terraform plan output to your clipboard and pass it to dlpo as follows:

```shell
pbpaste | dlpo
```

This enhances the plan output, transforming something like this:

```
  ~ resource "datadog_logs_pipeline_order" "test" {
        id        = "test"
        name      = "test"
      ~ pipelines = [
            "L_nrR7jTjOCPfjiojfisDg",
          + "No2kfjedFtuUTMMAX_wVxA",
            "dYLK_fImoap3QN4yIvIkaI",
          - "No2kfjedFtuUTMMAX_wVxA",
        ]
    }

Plan: 0 to add, 1 to change, 0 to destroy.
```

into a more readable format like this:

```
  ~ resource "datadog_logs_pipeline_order" "test" {
        id        = "test"
        name      = "test"
      ~ pipelines = [
            "Some Pipeline 1 (L_nrR7jTjOCPfjiojfisDg)",
          + "Some Pipeline 2 (No2kfjedFtuUTMMAX_wVxA)",
            "Some Pipeline 3 (dYLK_fImoap3QN4yIvIkaI)",
          - "Some Pipeline 2 (No2kfjedFtuUTMMAX_wVxA)",
        ]
    }

Plan: 0 to add, 1 to change, 0 to destroy.
```

This makes it much easier to understand which pipelines are being modified.
