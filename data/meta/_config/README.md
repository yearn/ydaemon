# Yearn metadata storage

[![Validation](https://github.com/yearn/ydaemon/workflows/Validation/badge.svg)](https://github.com/yearn/ydaemon/actions?query=workflow%3AValidation)
[![Styled With Prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg)](https://prettier.io/)

## What?

This repo contains all the metadata of the yearn ecosystem. Contents of the
[`meta`](./data/meta) directory are synced to IPFS for storage, accessible through
our gateway [meta.yearn.network](https://meta.yearn.network). Consistency of
the stored data is verified by smalls scripts and schemas.

## Adding documents

Any document can be added to the [`meta`](./data/meta) directory, but there are some
special checks to ensure consistency and ease of accessibility:

- All JSON files that share names with the files in the [`schema`](./data/meta/_config/schema)
  directory must follow the defined schema, otherwise verification will fail.
- Any folder that begins with `0x` is considered as an address. The address must
  be checksummed, otherwise verification will fail.
- All files named `index.json` will be ignored by git and will be
  overwritten by the indexing process. (see [indexing](#indexes))


## Adding schemas

Schemas can be created in the root of the [`schema`](./data/meta/_config/schema) folder. For
syntax you can take a look at the [JSON schema specs](https://json-schema.org).
The [AJV](https://github.com/ajv-validator/ajv) library is used to validate the
data with the provided schemas.

## Syncing with IPFS

After each commit to main, direct or as a result of a merged pull request, a
sync to IPFS is triggered.

## Indexes

Before each deployment the [`meta`](./data/meta) directory is scanned and an
`index.json` file is generated inside each directory (root included). The file
follows the [`index.json` schema](./data/meta/_config/schema/index.json) and will contain
information about the files and folders stored in that directory. For an example
see [meta.yearn.network/index.json](https://meta.yearn.network/json)

## Translations

Anything under protocols, strategies, and tokens are able to be translated. In the json files listed, we have locale codes and the English text to be translated. The name (if applicable) and description are what should be translated. If you dont see your locale code, make an issue and it will be added manually. To update new json's with the locale information run the python script `toLocale.py` located in the scripts folder.

## Helpful links

- 🌐 [Live site](https://yearn.network)
- ⚖️ [Governance forum](https://gov.yearn.finance)
- 📑 [Documentation](https://docs.yearn.finance)

## Contributing

Code style follows prettier conventions (`yarn format`). Commit messages follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).
