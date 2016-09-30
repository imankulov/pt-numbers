# pt-numbers

HTTP service and a library to convert numeric values to text (Portuguese language)
up to 999 trillions (short scale, used in Brasil).

See [Lista de nomes dos números][] and [Escalas curta e longa][] for more
information.

[Escalas curta e longa]: https://pt.wikipedia.org/wiki/Escalas_curta_e_longa
[Lista de nomes dos números]: https://pt.wikipedia.org/wiki/Lista_de_nomes_dos_n%C3%BAmeros


## Usage

Build

    go get github.com/imankulov/pt-numbers

Run

    PORT=8080 pt-numbers

Test

    curl 'http://localhost:8080/?n=1995'
    um mil e novecentos e noventa e cinco
