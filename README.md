# rest-agent
Use LLM to automatically generate and execute interface test cases.

## Install

```bash
make
mkdir -p /usr/local/bin
cp ./bin/rest-agent /usr/local/bin

```

## Quick Start

* Support models such as openai, groq, google vertex ai
* Run `rest-agent auth add` to configure the providers.
    * You can provide the password directly using the `--password` flag.
* Run `rest-agent generate` to generate test cases from openapi.[json|yaml] files.

## Examples

### provider

```bash
# add a new provider
rest-agent auth add -b openai -m openai/gpt-4o-mini -u https://openrouter.ai/api/v1 -l 8192 -p ${API_KEY}

# set the default provider
rest-agent auth default -p openai

# update the provider
rest-agent auth update -b openai -m openai/gpt-4o-mini

# remove a provider
rest-agent auth remove -b openai
```

### generate

> default outfile save in ~/Library/Application Support/rest-agent/274a1427519511bb62fac51c5f24a149_2024-07-28_21:49:38.yaml

```bash
# generate test cases from openapi.json and save to yaml
rest-agent generate -f pkg/ai/test_data/openapi.json -t yaml

# generate test cases from openapi.json and save to json
rest-agent generate -f pkg/ai/test_data/openapi.json -t json
```

## Roadmap

- [x] Add test case generate module

- [ ] Add api test case execution engine