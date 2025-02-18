# EN-US

## What is

The objective of this exercise is to implement a command-line interface (CLI) program that calculates the tax to be paid on profits or losses from stock market operations.

How should the program work?
Input: The program will receive lists, one per line, of stock market operations in JSON format through standard input (stdin).

Example input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":20.00, "quantity": 5000}]
[{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
{"operation":"sell", "unit-cost":10.00, "quantity": 5000}]`

Each line is an independent simulation; your program should not retain the state obtained from one line for the others.

The last line of the input will be an empty line.

Output: For each input line, the program must return a list containing the tax paid for each received operation. The elements of this list should be encoded in JSON format, and the output should be returned through standard output (stdout).

Example output: `[{"tax":0.00}, {"tax":10000.00}]
[{"tax":0.00}, {"tax":0.00}]`

The list returned by the program must have the same size as the list of operations processed in the input.

For example, if three operations (buy, buy, sell) are processed, the program's output should be a list with three values representing the tax paid on each operation.

## Technical Decisions

### Architecture

This project was designed with extensibility in mind, which is why all the I/O handling has been separated from the part that deals with the actual logic. Therefore, this code can be reused for other interfaces, such as a frontend, for example.

I/O it's basically stdin for input and stdout for output.

The logic was designed to be easy to understand and extend. If needed, it is simple to add new classes and functionalities without major changes to the structure.

### Tests

Soon

## Running

**WARNING**: To run this code directly on your machine, make sure you are using the correct Go version. The tests were conducted with version 1.24.0.

To run the program, use: `go run main.go < in/input2 > out/output2`

### Docker (Optional)

The Alpine variant of the Node Docker image was chosen to reduce the size, as Alpine is a lightweight Linux distribution.

1. Build the Docker image: `docker build -t capital-gains .`
2. Run the container: `docker run --rm -v $(pwd):/usr/src/app capital-gains sh -c '/usr/local/bin/app <replace-by-your-input-file> > <replace-by-your-output-file>'`

### Inputs

The project includes an `in` folder, where you can find all input examples for testing.

For example, to run case #8: `docker run --rm -v $(pwd):/usr/src/app nubank-test sh -c '/usr/local/bin/app < ./in/input6 > ./out/output6'`

### Outputs (Optional)

An `out` folder is also available where you can store output files. For instance, after running `docker run --rm -v $(pwd):/usr/src/app capital-gains sh -c '/usr/local/bin/app < in/input6 > out/output6'`, you can check the results with `cat out/output6`.

## Tests

The project includes unit tests, which can be executed with the following command: `SOON`

## Additional Information

`os.Stdin` and `os.Stdout` were chosen because they are native to Golang and efficiently handle streams. This allows processing large files without overwhelming memory.

# PT-BR

## O que é

O objetivo deste exercício é implementar um programa de linha de comando (CLI) que calcula o imposto a ser
pago sobre lucros ou prejuízos de operações no mercado financeiro de ações

Como o programa deve funcionar?

Entrada: O programa vai receber listas, uma por linha, de operações do mercado financeiro de ações em formato
json através da entrada padrão (stdin)

Exemplo entrada: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":20.00, "quantity": 5000}]
[{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
{"operation":"sell", "unit-cost":10.00, "quantity": 5000}]`

Cada linha é uma simulação independente, seu programa não deve manter o estado obtido em uma linha
para as outras.

A última linha da entrada será uma linha vazia.

Saída: para cada linha da entrada, o programa deve retornar uma lista contendo o imposto pago para cada operação
recebida. Os elementos desta lista devem estar codificados em formato json e a saída deve ser retornada
através da saída padrão (stdout).

Exemplo saída: `[{"tax":0.00}, {"tax":10000.00}]
[{"tax":0.00}, {"tax":0.00}]`

A lista retornada pelo programa deve ter o mesmo tamanho da lista de operações processadas na entrada.
Por exemplo, se foram processadas três operações (buy, buy, sell), o retorno do programa deve ser uma lista
com três valores que representam o imposto pago em cada operação.

## Decisões Técnicas

### Arquitetura

Este projeto foi feito visando extensibilidade, para isso foi separado toda a parte que lida com `I/O` da parte que lida com a lógica de fato. Portanto este código pode ser reutilizado para outras interfaces, como um frontend, por exemplo.

I/O é basicamente stdin para input e stdout para output.

A lógica foi pensada para ser fácil de entender e extensível. Se forem necessárias, é simples adicionar novas classes e funcionalidades sem grandes mudanças na estrutura.

### Testes

Em breve

## Execução

**AVISO:** Para executar este código diretamente na sua máquina, certifique-se de estar utilizando a versão correta do Go. Os testes foram realizados com a versão 1.24.0.

Para rodar o script, use: `go run main.go < in/input2 > out/output2`

### Docker (Opcional)

Foi escolhida a variante Alpine da imagem Docker do Node para reduzir o tamanho, pois o Alpine é uma distribuição Linux leve.

1. Build a imagem Docker: `docker build -t capital-gains .`
2. Execute o container: `docker run --rm -v $(pwd):/usr/src/app capital-gains sh -c '/usr/local/bin/app < <substitua-pelo-seu-arquivo-de-entrada> > <substitua-pelo-seu-arquivo-de-saída>'`

### Entradas

O projeto inclui uma pasta chamada `in`, onde você pode encontrar todos os exemplos de entrada para seus testes.

Por exemplo, para rodar o caso #8: `docker run --rm -v $(pwd):/usr/src/app capital-gains sh -c '/usr/local/bin/app < in/input8 > out/output8'`

### Saídas (Opcional)

Também há uma pasta chamada `out`, onde você pode armazenar seus arquivos de saída (se desejar). Por exemplo, após rodar `docker run --rm -v $(pwd):/usr/src/app capital-gains sh -c '/usr/local/bin/app < in/input6 > out/output6'`, você pode verificar o arquivo de saída com `cat out/output6`.

## Testes

Em breve

## Informações Adicionais

Foram escolhidos `os.Stdin` e `os.Stdout` porque são nativos do Golang e lidam eficientemente com streams. Isso permite o processamento de arquivos grandes sem sobrecarregar a memória.

