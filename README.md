# acorn-load-balancer-plugin

The Acorn Load Balancer plugin allows you to add arbitrary annotations to load-balancers that Acorn creates. This is ideal for situations that you need to configure load-balancers via annotations, such as with the AWS `LoadBalancerController`.

## Usage

This plugin requires Acorn to be installed as it is installed an acorn. There are two ways to specify what annotations you would like your `LoadBalancer` services to have and you can do a mix of both.

### --path

This flag outlines a path to a yaml file containing `key: value` pairs that will be used as annotations.

For example this file:

```yaml
# ./annotations.yaml
foo: bar
```

Will be used by this command:
```bash
acorn run ghcr.io/tylerslaton/acorn-load-balancer-plugin:main --path ./annotations.yaml
```

### --annotations

If you would not like to use a file and would instead like to specify them directly you can do so with this flag. It is done in the form of `key=value,foo=bar` to specify as many as you may need.

```bash
acorn run ghcr.io/tylerslaton/acorn-load-balancer-plugin:main --annotations key=value,foo=bar
```

## Development

The best way to run this plugin locally is through development mode but you will need access to a cluster with Acorn installed.

```bash
acorn run --name controller -i .
```

## License
Copyright (c) 2023 [Acorn Labs, Inc.](http://acorn.io)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.