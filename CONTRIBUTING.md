# Contributing

## Structure of the project

  - src/ (Golang / Wails)
  - frontend/ ( Typescript / Sveltekit - All the UI for the Desktop APP & the Web version)
  - docs/ (All the development documentation)
  - frontend/src/lib/i18n (All the translations of the project)

## Requisites (only apply to code contributions)
  - If you want to contribute first you need to check the issues, them if you like any of the open issues work on it and merge it to the project (obviously you can open a new issue to enhance the features or correct any bug you found to work on it later)
  - Try to do self-explanatory code (if cannot be you can comment to enhance the comprehension)

## Resources of interest

  - [LibreRemotePlay Docs](./docs/README.md)
  - [How to run the project](./README.md#run-dev)
  - [How to build the project](./README.md#build)

## How to

  1. Fork this repository
  2. Clone it
  3. Work on the issue
  4. When you have finished make a pull request to merge it with the main branch
  5. Wait for merge (maybe it will not be merged at first because of bad code)
  6. Done

## Translations 🔠

By default English is the language of reference so you can check if there are entries in english that are missing in the language you may want to contribute to 

### How to

  1. Fork this repository
  2. Clone it
  3. Work on your translations (located in frontend/src/lib/i18n):
     - Create a JSON file of the language and register the language in the i18n.ts file (all of this if the language is not added already)
     - Add the entries (you can do manually or using [i18n Ally extension](https://marketplace.visualstudio.com/items?itemName=Lokalise.i18n-ally))
  5. When you have finished make a pull request to merge it with the main branch
  6. Wait for merge
  7. Done


## Thank you for reading this and also for your interest on contributing
