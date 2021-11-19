# Cell structure

## Properties

### `id` - Cell identity
- Mandatory: `true`
- Format: `/[a-zA-Z]{1}[\w\d]{0,49}/g` or `$ref`

### `cell_type` - Type of cell
- Mandatory: true
- Values:
  - string
  - boolean
  - integer
  - double
  - date
  - array
  - enum
