# Model file structure

## Properties

### `id` - Data identity

- Mandatory: `true`
- Format: `/[a-zA-Z]{1}[\w\d]{0,49}/g`

### `cells` - Data cells collection

- Mandatory: `true`
- Format: **[Cell structure](./model_file_cell.md)**

## Template

```json
{
  "id": "DATA_ID",
  "cells": []
}
```
