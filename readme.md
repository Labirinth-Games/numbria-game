# Numbria - The King Orc

## Livro do jogo
Aqui segue o descritivo e roteiro do jogo onde será baseado a aventura RPG Guiada com um oraculo virtual para 
que venha conduzir o mundo.

## Aventura
Essa aventura se baseia em uma missão ao qual o mercenário precisa eliminar uma ameaça, então prontamente ele vai até a floresta escura para encontrar e eliminar seu alvo, mas ele não esperava que seu alvo fosse tão inteligente e que comandava outras criaturas. Durante sua missão pode acontecer muitas coisa inclusive nada entao vamos descobre em numbria - the king orc.

## Descrição do mundo

O mundo é magico do jeito que consegue ser, onde a magia não é abundante mas se souber onde buscar pode ser que 
consiga domina-la e usar de forma simples mas vantajosa dentro de um universo onde é muito raro o uso dela.

O jogo é medieval onde existe especializações de funçoes que conhecemos como classes de pessoas, cada habilidade
foi sido criada para poder defender a humanidade dos monstros que cercam as cidades.

As classes que possuimos são:
- Arqueiros
- guerreiro
- assasino
- medico

### Ambiente
O jogador se encontra numa floresta conhecida como a floresta negra, a floresta é composta por:

- clareira
- caverna
- rios
- floresta densa

## Items no mundo
No mundo temos diversos itens que podem ser coletados e usados para poder ajudar na sua aventura

## Estatisticas
Durante a gameplay precisamos definir alguns valores padrões referentes aos status, segue abaixo a lista de stats e seus valores maximos baseado em rolagem de dados.

|  Stats    |dice | description |
|--|--|--|
| Accuracy  | N  | Precisão é usado como base para saber se o atacante consegue acertar ou não |
| Strength  | d12 | Força aplicada que resultará em dano | 

## Açoes do jogador
o jogador consegue fazer algumas açoes como:
- perguntar o que está a sua frente
- Perguntar onde está
- Pode andar libremente pelo mapa, apenas informando a direção que deseja ir (norte, sul, leste, oeste)
- Pode tentar fugir de uma situação como ataque ou suspeita de emboscada
- Pode atacar um alvo que ele tenha conhecimento
- Pode observar, assim se tiver sorte ter mais informações oferecidas pelo sistema
- Entrar em lugares

## Monstros
Vamos ter alguns tipos de monstros iniciais incluindo nosso alvo e suas estatisticas

### Goblins
|Stats | value |
|-|-|
| HP        | 4 | 
| Accuracy | 2 |
| Strength  | d4 |

### Huki
Monstro exclusivo, pequeno menor que um goblin, mas com pele escura e olhos amarelos grandes, vesntindo folhas secas e pelos como de coelho saindo abaixo nas narinas apontando para cima.

|Stats | value |
|-|-|
| HP        | 5 | 
| Accuracy  | 1 |
| Strength  | d4 |

### Orc Menor

|Stats | value |
|-|-|
| HP        | 12 | 
| Accuracy  | 3  |
| Strength  | d6  |

### Orc Major

|Stats | value |
|-|-|
| HP        | 20 | 
| Accuracy  | 4  |
| Strength  | d6  |

### Orc Leader

|Stats | value |
|-|-|
| HP        | 35 | 
| Accuracy  | 6  |
| Strength  | d10 |



## Batalha
A arma usada vai influenciar na contagem de dano aflingida durante o combate, se o jogador tiver uma arma que da +1 em ataque esse valor será somado no dano.

Quando o jogador vai atacar o inimigo são rolado 2d12 uma para cada, se o atacante tiver mais pontos o ataque é bem sucedido,
mas se o atacado tiver o numero  maior igual o ataque tem falha, ou seja, ele consegue desviar. Mas o jogador pode adquirir pontos de precisão que influencia nessa rolagem, ele consegue rolagem + pontos assim ajudando no ataque bem sucedido.

### Caracteristicas usadas no combate
| Type | Description |
|-|-|
|Precisão (accuracy)| É rolado um dado de destreza(d8) caso o valor seja maior ou igual a destreza do alvo é considerado um acerto então pode rolar o dado de ataque.|
|Ataque (strength) | Quando a precisão tem sucesso é rolado o dado de força(strength) do atacante para saber quanto dano vai ser deferido, se o jogador tiver equipado com alguma arma, será acrecido o adicional da arma no dano deferido |
| Vida (HP) | pontos de vida do jogador|

## Eventos de gameplay
Durante o jogo podemos ter alguns eventos toda vez que o player se move pelo mundo tem a chance de um desses eventos acontecerem.

- aparição de monstros
        - Pode encontrar o monstro a sua frente
        - Pode encontrar um ninho de monstro assim podento ter a chande ou de pegar um material do monstro ou o monstro aparecer
- encontrar algum item ou arma especial
- encotrar um NPC perdido, que pode vender algo, ou dar informaçoes sobre o lugar ou onde pode ter algo legal


# WIKI
Essa wiki é sobre a programação da historia do jogo e itens conhecidos como books, é um modelo experimental de programar açoes e historia do jogo sem necessáriamente manipular codigo

## Estrutura
um arquivo .book comeneça com o tipo do book, podemos ver a tag `__TYPE__` seguido do separador `:` e o valor que ele contem, usado para poder definir qual o dominio do libro, por exemplo podemos ver `__TYPE__:Lore`.

após isso podemos ver separadores definidos por `-----` onde em cada bloco temos comandos referentes a um elemento/sessão que será executada dentro do game. Abaixo podemos ver mais sobre os comandos que podem existir

### Comandos no Lore
Os comandos são definidos por `#` e um nome uppercase onde será tratado de foema interna pelo sistema como mappeador de comando. dos comandos temos:

#### `#INDEX` 
Ele é usado para indexar a proxima parte da historia, assim quando um trecho finalizar e ocorrer qualquer tipo de açao, evento ou iteração a aplicação sabe para onde continuar a historia.

#### `#CONTENT`
Conteudo que será imprimido na tela do jogador, onde pode ser multilinhas e cada linha terá um tempo para ser imprimida

#### `#NEXT`
Esse comando é usado para poder dizer qual á a proxima parte da historia, então ele é a ligaçao da chave encadeada de partes da historias

#### `#NEXT_AUTOMATIC`
Esse comando é tem a mesma funçao do next só que é executado de forma automatica, quando termina todo o dialog de uma lore ele chama o proximo setado no next

#### `#SYSTEM`
Mensagem do sistema para o jogar, momento que o sistema pede para ele fazer alguma coisa ou tomar uma decisão.

### Comandos no Player
esses comandos sao do book relacionado ao jogador basicamente é o treinamento dos possiveis comandos que o jogador pode fazer a aplicação e como ela vai responder a esse cenário

#### `#ASK`
Aqui temos uma lista de possiveis comandos que o jogador pode digitar, ou seja, seria o treinamento quando o player perguntar X devo responder Y, onde o Y é o comando abaixo o `#ANSWER`.

#### `#ANSWER`
Lista de possiveis respostas para as perguntas feitas acima, é uma lista por que o sistema vai decidir de forma aleatoria qual resposta imprimir quando alguma ask ou parte dela for perguntada pelo jogador.

#### `#ACTION`
Comando importante, usado para definir qual o nome da funçao que será chamada quando o jogador perguntar algo que se enquadre com o treinamento do comando #ASK. É importante lembrar que ele irá chamar o metodo definido dentro da entidade descrita no type no inicio do arquivo, ou seja, como aqui é player então vai chamar o metodo dentro do pacote do player

#### `#PRIORITY`
Quando adicionado é responsavel por valor na busca após o treinamento ou seja, se tiver duas açoes com comandos bem semelhantes ele serve de desempate para poder imprimir uma ação

### Comandos no Betiario

#### `#OBSERVER_SUCESS`
Usado como resposta quando o player encontra com a criatura e tenta descobrir quem é e ao rolar a iniciativa
consegue ver, ou seja, a rolagem é favoravel ao player entao o dialog vai ser esse aqui

### Comandos no Eventos

#### `#FAIL`
Usado quando rola o teste pra saber se consegue descobrir, em caso de falha ele mostra essa mensagem

