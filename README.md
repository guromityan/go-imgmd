# imgmd
Convert image to Markdown.

## Build

Follow the steps below to build.

```shell
git clone https://github.com/guromityan/go-imgmd.git
cd go-imgmd
make
```


## Usage

Prepare a directory containing images to include in Markdown.

```
tree samples

samples
├── america
│   ├── north-america
│   │   ├── Canada.png
│   │   ├── United-States-of-America.png
│   │   └── test.txt
│   └── south-america
│       ├── Argentina.png
│       └── Brazil.png
├── asia
│   ├── China.png
│   └── Japan.png
└── europe
    ├── Italy.png
    └── Sweden.png
```

The following command will generate a Markdown file containing the images under the samples / directory as content.

```shell
imgmd samples/
```

As a result, samples.md was generated

The origin of samples.md is the name of the directory specified as the target.

```shell
ls

samples    samples.md
```

Use the --ouput option to rename the output file.

```shell
imgmd --output document.md samples/
```

The header level of the generated file is taken from the directory name under the directory specified as the target.
