# Check if a Docker image with a specific tag exists in ECR

This action returns whether a Docker image with a specific tag exists in your Amazon ECR repository.

## Usage

```yaml
name: ci

on:
  push:
    branches:
      - 'main'

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - id: check-ecr-image
        uses: mnmandahalf/check-ecr-image-exists@v0.1.5
        with:
          region: us-east-1
          repository-name: "my-repository/name"
          image-tag: "1.2.3"
          access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
      - name: Perform other action based on the output
        if: ${{ steps.check-ecr-image.outputs.image-exists == '0' }}
```

see [action.yml](https://github.com/mnmandahalf/github-actions-ecr-image-exists/blob/main/action.yml)

## inputs

Following inputs can be used as `steps[*].with` keys

| Name                | Required | Description                                                                                                                                                                                                                |
| ------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `region`            | true     | AWS Region Example: `us-east-1`                                                                                                                                                                                            |
| `repository-name`   | true     | AWS ECR Repository Name Example: `my-repository/name`. Note: Do not enter the whole ECR repository url, i.e.  `ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/my-repository/name`, just the last bit, i.e. my-repository/name` |
| `image-tag`         | true     | ECR Image Tag, Example: `1.2.3`                                                                                                                                                                                            |  |
| `access-key-id`     | false    | AWS access key id                                                                                                                                                                                                          |
| `secret-access-key` | false    | AWS secret access key                                                                                                                                                                                                      |


## Output
| Name | Description | Default |
| - | - | - |
| `image-exists` | Flag indicating whether the docker image exists or not. Returns `'0'` or `'1'` | N/A |
