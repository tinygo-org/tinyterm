0.5.0
---
- **core**
    - add SetRotation() to tinyterm Displayer interface, to make it a little easier to adjust in actual use
- **modules**
    - update to latest drivers and tinyfont


0.4.0
---
- **core**
    - add displays subpackage (#13)
    - add support for Badger2040 (#12)
- **license**
    - Update license to 2024
- **modules**
    - update to latest drivers and tinyfont
- **docs**
    - some small improvments to README
- **examples**
    - httpclient: update to use new netdev interface


0.3.0
---
- **all**
    - add .gitignore
- **build**
    - add Makefile and GH Actions workflow
    - add 'make test' alias to Makefile
    - switch to ghcr.io for docker container
- **docs**
    - add GH action badge to README
    - add godoc explanations and general header about what this package does
    - explain what the different examples are and how to run them
- **examples**
    - correct go fmt of pybadge example
    - unify with initdisplay and add gopher-badge
    - examples/initdisplay/pybadge: fix rotation
- **license**
    - add same license file as the rest of TinyGo
- **modules**
    - update to latest drivers
- **core**
    - handle displays without hardware scrolling (#4)
