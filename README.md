# Crabi-Code-Challenge (CCC)

Ejercicio de servicio web de PLD

## CHANGELOG
### v0.1 
FEATURES:
1) Endpoint CreateUser. No soporta el servicio externo de PLD
2) Endpoint GetUserInformation. No hay auth, solo username por parametro.

TO-DO
1) consumir PLD externo
2) Endpoint login - respuesta de token
3) usar el token de respuesta como autenticacion de GetUserInformation
4) Refactor - organizar en paquetes/modularizar.

### v0.2
FEATURES:
1) Endpoint CreateUser. El servicio externo es un stub.
2) Endpoint Login que retorna accessToken para uso de la plataforma.
3) Endpoint getMe, con Authorization via token, devuelve el usuario que hace el llamado.

TO-DO
1) Consumir el servicio PLD especifico

NICE-TO-HAVE
1) Mejorar el logging de la plataforma.
2) Mejorar el objeto Token
    - que tenga una expiracion temporal desde que se crea por seguridad
    - upgradear a JWT
3) Proveer usuarios iniciales 

### v1.0 - MVP

FEATURES:
1) Endpoint CreateUser con proveedor externo para PLD.
2) Endpoint Login que retorna accessToken para uso de la plataforma.
3) Endpoint getMe, con Authorization via token, devuelve el usuario que hace el llamado.


[![Unit Tests](https://github.com/fedemiodo/Crabi-code-challenge/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/fedemiodo/Crabi-code-challenge/actions/workflows/unit-tests.yml/badge.svg)
[![Coverage Status](https://codecov.io/github/fedemiodo/Crabi-code-challenge/coverage.svg?branch=release-candidate)](https://codecov.io/gh/fedemiodo/Crabi-code-challenge/branch/release-candidate)
[![Group loading check](https://github.com/fedemiodo/Crabi-code-challenge/actions/workflows/loading-groups.yml/badge.svg)](https://github.com/fedemiodo/Crabi-code-challenge/actions/workflows/loading-groups.yml)
[![Markdown Lint](https://github.com/fedemiodo/Crabi-code-challenge/actions/workflows/markdown-lint.yml/badge.svg)](https://github.com/fedemiodo/Crabi-code-challenge/actions/workflows/markdown-lint.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-chi/render.svg)](https://pkg.go.dev/github.com/go-chi/render)

[![GitHub release](https://img.shields.io/github/release/fedemiodo/Crabi-code-challenge.svg)](https://github.com/fedemiodo/Crabi-code-challenge/releases/latest)

Quick links

- [**Explore the docs**](docs/)
- [Report a defect](https://github.com/fedemiodo/Crabi-code-challenge/issues/new?labels=Type%3A+Defect)
- [Request a feature](https://github.com/fedemiodo/Crabi-code-challenge/issues/new?labels=Type%3A+Feature)



## License

- The code is licensed under [MIT](LICENSE).
- The documentation is licensed under [CC BY-SA 4.0](http://creativecommons.org/licenses/by-sa/4.0/).


## Installation

