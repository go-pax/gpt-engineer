# GPT-Engineer

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

GPT-Engineer is a versatile software engineering tool built with Go. It offers a host of features designed to streamline your development workflow.

## Features

1. Code Generation: Automatically generate full codebases in any language for the project you describe.
2. Automated Testing: Tools for automatically testing code to catch bugs and other issues.
3. Performance Profiling: Features for measuring and improving the performance of your code.
4. Logging: All input and output with OpenAI is written out to files.
5. Prebuilt Executables: Executables available for most systems so no need to build from source code.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Download and Run

You can just download the executable from the Release section on GitHub.

1. Download the archive that matches your system here, https://github.com/go-pax/gpt-engineer/releases/tag/v0.1.0
2. Extract the contents of the archive into a folder where you will run `gpt-engineer`
3. Open a terminal/command prompt and `cd` to the folder with `gpt-engineer`
4. Generate an OpenAI API key at, https://platform.openai.com/account/api-keys and copy the key.
5. Set the `OPENAI_API_KEY` environment variable with the key in the terminal/command prompt,
   - __(macos/linux)__ `export OPENAI_API_KEY=sk_123456789`
   - __(windows)__ `set OPENAI_API_KEY=sk_123456789`
6. Run `gpt-engineer --help` to view the arguments you can set
   - __(macos/linux)__ `./gpt-engineer --help`
   - __(windows)__ `gpt-engineer.exe --help`
7. Run `gpt-engineer` and create your first project
   - __(macos/linux)__ `./gpt-engineer -prompt "hello world app in go" ./output_folder`
   - __(windows)__ `gpt-engineer.exe -prompt "hello world app in go" output_folder`
8. (optional) Choose to execute the source code by typing "yes" or "no" when prompted. This requires you have everything installed to execute the source code.

## Developer Setup

### Prerequisites

What things you need to install the software and how to install them:

```bash
go get -u github.com/go-pax/gpt-engineer
```

### Installation

A step-by-step series of examples that tell you how to get a development environment running:

```bash
go install github.com/go-pax/gpt-engineer
```

## Usage

Here are the steps to setup your GPT-Engineer environment. You can use both OpenAI or Azure OpenAI.

### Setup:

Get go packages,

```bash
go mod tidy
```

#### OpenAI
Run using [OpenAI](https://platform.openai.com) API key. Set the following environment variable:

```bash
export OPENAI_API_KEY=[your openai api key]
```
Attempts to use GPT4 if not then fallback to GPT-3.5.

#### Azure OpenAI
You will need an Azure account and an OAI resource, https://oai.azure.com/portal

```bash
export OPENAI_API_KEY=[your azure openai api key]
export OPENAI_API_BASE=https://[deployment name].openai.azure.com/
```

### Run:
- Create an empty folder. If inside the repo, you can run:
  - `cp -r projects/example/ projects/my-new-project`
- Edit `main_prompt` file in your new folder to specify what you want to build
- Run `go run . ./projects/my-new-project`, default language is English, if you want to use other language, and `-lang` argument like `go run . -lang=Chinese`
  - (Note, `go run . --help` lets you see all available options. For example `--steps use_feedback` lets you improve/fix code in a project)

**(optional) Database**
You should use the scheme of the database instead of the path to your projects. View the readme files in the `database` folder to understand how to use. example, instead of `./projects/my-new-project` use `file://./projects/my-new-project` to ensure the file database is used.

**(optional) Azure**
- use the argument `-model [your deployment name]`, when using Azure this must be set

### Results:

- Check the generated files in `projects/my-new-project`

## Fine Tuning:

You can specify the "identity" of the AI agent by editing the files in the `identity` folder.
Editing the identity, and evolving the main_prompt, is currently how you make the agent remember things between projects.

## Logging:

Each step in `steps.go` will have its communication history with GPT4 stored in the `logs` folder.


## Contributing

Please read [CONTRIBUTING.md](https://github.com/go-pax/gpt-engineer/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/go-pax/gpt-engineer/LICENSE.md) file for details.

## Thanks

Thanks to the projects;
- [AntonOsika/gpt-engineer](https://github.com/AntonOsika/gpt-engineer) that created GPT-Engeneer
- [geekr-dev/gpt-engineer](https://github.com/geekr-dev/gpt-engineer) this project is based off
