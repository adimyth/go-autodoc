# <p align="center">autodoc - A simple document generator</p>

`autodoc` is a simple document generator. It parses docstrings for any function & generates markdown document.

---

## ❓ What is it?

`autodoc` parses the docstrings for all functions in a given source code & generates markdown representation for them.

> **Note**
> We are following the Numpy DocString Format.

To understand, how to write function docstrings following the **_numpy docstring format_** - checkout [this excellent guide](https://developer.lsst.io/python/numpydoc.html)

## 🚀 Demo

Clonning the repo & run the `main` file. It will loop through all the files in `src` directory & generate a markdown file for each source file & store it in the `docs` directory.

In the repo, I have added 2 source files. Running the `main` script, created 2 new files in `docs` directory.

```bash
.
├── docs
│   ├── fibonacci.md
│   └── square_root.md
└── src
    ├── fibonacci.go
    └── square_root.go
```

Compare & check the docs file

## 🤔 How?

The `autodoc` exports a `DocGenerator` method that takes in a filepath & generates a markdown of the same name in docs file. It does this by -

1. Reading the source file
2. Generates the Abstract Syntax Tree for the source code
3. Traverses the tree to find `function` nodes
4. Reads the docstring assosciated with it & the function name
5. Formats it into a markdown representation
6. Writes it to a file inside `docs`

The `main` script recursively loops over all the files in `src` directory & applies all the previous step to each file
