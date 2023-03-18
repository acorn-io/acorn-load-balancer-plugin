# acorn-load-balancer-plugin

The Acorn Load Balancer plugin allows you to add arbitrary annotations to load-balancers that Acorn creates. This is ideal for situations that you need to configure load balancers via annotations, such as with the AWS LoadBalancerController.

### Build

```bash
make build
```

### Development

The best way to run the plugin is through acorn. Run 

```bash
acorn run --name controller -i .
```

### Production

```bash
acorn run ghcr.io/tylerslaton/acorn-load-balancer-plugin:main
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