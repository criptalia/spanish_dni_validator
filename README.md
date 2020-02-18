# spanish_dni_validator

A Golang module to validate Spanish DNI (Documento Nacional de Identidad) numbers.
It's a porting of the PHP library at https://github.com/ulabox/nif-validator.

It can also validate:
- NIF (Número de Indentifación Fiscal)
- CIF (Certificado de Identificación Fiscal)
- NIE (Número de Identidad de Extranjero) 

## Installation

```
go get github.com/criptalia/spanish_dni_validator
```

## Usage


```go
package main

import (
    "github.com/criptalia/spanish_dni_validator"
    "fmt"
)

func main() {
    //CIF
    if(spanish_dni_validator.IsValid("B65410011")) {
        fmt.Println("B65410011 is valid")
    } else {
        fmt.Println("B65410011 is NOT valid!")
    }

    //DNI
    if(spanish_dni_validator.IsValid("93471790C")) {
        fmt.Println("93471790C is valid")
    } else {
        fmt.Println("93471790C is NOT valid!")
    }

    //NIE
    if(spanish_dni_validator.IsValid("X5102754C")) {
        fmt.Println("X5102754C is valid")
    } else {
        fmt.Println("X5102754C is NOT valid!")
    }
}
```

You can also use an import alias to make calls look prettier:
```go
package main

import (
    dni "github.com/criptalia/spanish_dni_validator"
    "fmt"
)

func main() {
    if(dni.IsValid("B65410011")) {
        fmt.Println("B65410011 is valid")
    } else {
        fmt.Println("B65410011 is NOT valid!")
    }
}
```

## Personal and entity NIFs
You can also separetely validate personal and entity NIFs:

```go
package main

import (
    dni "github.com/criptalia/spanish_dni_validator"
    "fmt"
)

func main() {
    //CIF
    if(dni.IsValidEntity("B65410011")) {
        fmt.Println("B65410011 is valid")
    } else {
        fmt.Println("B65410011 is NOT valid!")
    }

    //DNI
    if(dni.IsValidPersonal("93471790C")) {
        fmt.Println("93471790C is valid")
    } else {
        fmt.Println("93471790C is NOT valid!")
    }

    //NIE
    if(dni.IsValidPersonal("X5102754C")) {
        fmt.Println("93471790C is valid")
    } else {
        fmt.Println("93471790C is NOT valid!")
    }
}
```

## Individual NIF types
You can also separately validate individual NIF types:
```go
package main

import (
    dni "github.com/criptalia/spanish_dni_validator"
    "fmt"
)

func main() {
    //CIF
    if(dni.IsValidCif("B65410011")) {
        fmt.Println("B65410011 is valid")
    } else {
        fmt.Println("B65410011 is NOT valid!")
    }

    //DNI
    if(dni.IsValidDni("93471790C")) {
        fmt.Println("93471790C is valid")
    } else {
        fmt.Println("93471790C is NOT valid!")
    }

    //NIE
    if(dni.IsValidNie("X5102754C")) {
        fmt.Println("93471790C is valid")
    } else {
        fmt.Println("93471790C is NOT valid!")
    }
}
```
