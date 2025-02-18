# EN-US

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

