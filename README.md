# Lector OCR

#### Seguir los siguientes pasos para instalar el api

Para ejecutar el api, dentro de la carpeta del proyecto:
```sh
docker-compose up --build
```
Para parar el api, dentro de la carpeta del proyecto:
```sh
docker-compose down
```

#### Funcionalidad

Endpoint de prueba:
```sh
GET / 
```
Respuesta:
```sh
{"data":"ocr"}
```
Endpoint primera historia de usuario:
```sh
GET /number 
```
Respuesta:
```sh
{"number":"88888888?"}
```

#### Primera Historia de Usuario

Imagina que trabajas para un banco, que recientemente compró una máquina ingeniosa para ayudarlo a leer cartas y faxes enviados a las sucursales. La máquina escanea los documentos en papel y produce un archivo con una cantidad de entradas que se ven así:
```sh
_ _ _ _ _ _ _ 
| _| _||_||_ |_ ||_||_| 
||_ _| | _||_| ||_| _| 
```
Cada entrada se conforma de 4 líneas y cada línea tiene 27 caracteres. Las primeras 3 líneas de cada entrada contienen un número de cuenta escrito usando pipes y guiones bajos, y la cuarta línea está en blanco. Cada número de cuenta debe tener 9 dígitos, todos los cuales deben estar en el rango 0-9. Un archivo normal contiene alrededor de 500 entradas. 
La primera tarea es escribir un programa que pueda tomar este archivo y convertirlo en números de cuenta reales. 

#### Segunda Historia de Usuario 
Una vez hecho lo anterior, rápidamente te das cuenta de que la ingeniosa máquina no es infalible. A veces comete errores en su escaneo. En consecuencia, el siguiente paso es validar que los números que lee son, de hecho, números de cuenta válidos. Se sabe que un número de cuenta válido tiene una suma de verificación válida y esto se puede calcular de la siguiente manera: 
```sh
account number: 3 4 5 8 8 2 8 6 5 
position name: d9 d8 d7 d6 d5 d4 d3 d2 d1 
checksum calculation: 
(1*d1 + 2*d2 + 3*d3 + … + 9*d9) mod 11 = 0 
```
Se necesita escribir un programa que calcule la suma de verificación para un número de cuenta determinado e identifique si se trata de un número de cuenta válido.

#### Tercera Historia de Usuario

Tu jefe está ansioso por ver sus resultados. Él te pide que escribas un archivo de los hallazgos, con una línea para cada número de cuenta, en este formato: 

```sh
457508000 OK 
664371495 ERR 
86110??36 ILL 
```
Es decir, el archivo tiene un número de cuenta por fila. Si algunos caracteres son ilegibles, se reemplazan por un '?'. En el caso de que alguna suma de verificación sea incorrecta o existe algún número ilegible, este estado se indica en una segunda columna. 

#### Pistas 

La recomendación es encontrar una forma de escribir celdas 3x3 en 3 líneas en su código, para que formen dígitos identificables. Incluso si el código en realidad no los representa de esa manera internamente. Es preferible leer: 

```sh
" " + 
"|_|" + 
" |" 
```
en lugar de: 
```sh
" |_| |"
```
#### Casos de prueba sugeridos 

```sh
1a historia de usuario

 _ _ _ _ _ _ _ _ _ | || || || || || || || || | |_||_||_||_||_||_||_||_||_| 
=> 000000000 OK 
| | | | | | | | | | | | | | | | | | 
=> 111111111 ERR 
_ _ _ _ _ _ _ _ _ _| _| _| _| _| _| _| _| _| |_ |_ |_ |_ |_ |_ |_ |_ |_ 
=> 222222222 ERR 
_ _ _ _ _ _ _ _ _ _| _| _| _| _| _| _| _| _| _| _| _| _| _| _| _| _| _| 
=> 333333333 ERR 
|_||_||_||_||_||_||_||_||_| | | | | | | | | | 
=> 444444444 ERR 
_ _ _ _ _ _ _ _ _ |_ |_ |_ |_ |_ |_ |_ |_ |_ _| _| _| _| _| _| _| _| _| 
=> 555555555 ERR 
_ _ _ _ _ _ _ _ _ |_ |_ |_ |_ |_ |_ |_ |_ |_ |_||_||_||_||_||_||_||_||_| 
=> 666666666 ERR 
_ _ _ _ _ _ _ _ _ | | | | | | | | | | | | | | | | | | 
=> 777777777 ERR
_ _ _ _ _ _ _ _ _ |_||_||_||_||_||_||_||_||_| |_||_||_||_||_||_||_||_||_| 
=> 888888888 ERR 
_ _ _ _ _ _ _ _ _ |_||_||_||_||_||_||_||_||_| _| _| _| _| _| _| _| _| _| 
=> 999999999 ERR 
_ _ _ _ _ _ _ | _| _||_||_ |_ ||_||_| ||_ _| | _||_| ||_| _| 
=> 123456789 ERR 

3a historia de usuario:

 _ _ _ _ _ _ _ _ | || || || || || || ||_ | |_||_||_||_||_||_||_| _| | 
=> 000000051 OK 
_ _ _ _ _ _ _ |_||_|| || ||_ | | | _ | _||_||_||_| | | | _| 
=> 49006771? ILL 
_ _ _ _ _ _ _ _ | _| _||_| _ |_ ||_||_| ||_ _| | _||_| ||_| _ 
=> 1234?678? ILL
```

#### Ejemplo de Post, envía el contenido del archivo test_numbers.txt, en un json codificado en base64

```sh
curl -i -X  POST http://localhost:8080/numbers   -H "Accept: application/json" -H "Content-Type: application/json"   -d '{ 
    "numbers": 
    "IF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gCnwgfHwgfHwgfHwgfHwgfHwgfHwgfHwgfHwgfAp8X3x8X3x8X3x8X3x8X3x8X3x8X3x8X3x8X3wKCiAgICAgICAgICAgICAgICAgICAgICAgICAgIAogIHwgIHwgIHwgIHwgIHwgIHwgIHwgIHwgIHwKICB8ICB8ICB8ICB8ICB8ICB8ICB8ICB8ICB8CgogXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAKIF98IF98IF98IF98IF98IF98IF98IF98IF98CnxfIHxfIHxfIHxfIHxfIHxfIHxfIHxfIHxfIAoKIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gCiBffCBffCBffCBffCBffCBffCBffCBffCBffAogX3wgX3wgX3wgX3wgX3wgX3wgX3wgX3wgX3wKCiAgICAgICAgICAgICAgICAgICAgICAgICAgIAp8X3x8X3x8X3x8X3x8X3x8X3x8X3x8X3x8X3wKICB8ICB8ICB8ICB8ICB8ICB8ICB8ICB8ICB8CgogXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAKfF8gfF8gfF8gfF8gfF8gfF8gfF8gfF8gfF8gCiBffCBffCBffCBffCBffCBffCBffCBffCBffAoKIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gCnxfIHxfIHxfIHxfIHxfIHxfIHxfIHxfIHxfIAp8X3x8X3x8X3x8X3x8X3x8X3x8X3x8X3x8X3wKCiBfICBfICBfICBfICBfICBfICBfICBfICBfIAogIHwgIHwgIHwgIHwgIHwgIHwgIHwgIHwgIHwKICB8ICB8ICB8ICB8ICB8ICB8ICB8ICB8ICB8CgogXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAKfF98fF98fF98fF98fF98fF98fF98fF98fF98CnxffHxffHxffHxffHxffHxffHxffHxffHxffAoKIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gIF8gCnxffHxffHxffHxffHxffHxffHxffHxffHxffAogX3wgX3wgX3wgX3wgX3wgX3wgX3wgX3wgX3wKCiAgICBfICBfICAgICBfICBfICBfICBfICBfIAogIHwgX3wgX3x8X3x8XyB8XyAgIHx8X3x8X3wKICB8fF8gIF98ICB8IF98fF98ICB8fF98IF98CgogXyAgXyAgXyAgXyAgXyAgXyAgXyAgXyAgICAKfCB8fCB8fCB8fCB8fCB8fCB8fCB8fF8gICB8CnxffHxffHxffHxffHxffHxffHxffCBffCAgfAoKICAgIF8gIF8gIF8gIF8gIF8gIF8gICAgIF8gCnxffHxffHwgfHwgfHxfICAgfCAgfCAgfCAgfAogIHwgX3x8X3x8X3x8X3wgIHwgIHwgIHx8X3wKCiAgICBfICBfICAgICBfICBfICBfICBfICAgIAogIHwgX3wgX3x8X3wgXyB8XyAgIHx8X3x8X3wKICB8fF8gIF98ICB8IF8gfF98ICB8fF98IF98" 
}'
```

#### Acceso a la base de datos docker desde una terminal: 
```sh
docker exec -it ocr_db mysql -u santo -pasdf bunsan
```

#### Tabla log_db: 

| Field  | Type  | Null  | Key  | Default  |  Extra |
|---|---|---|---|---|---|
| id   | varchar(36)  | NO  | PRI  | NULL   |   |
| endpoint  | text  | YES  |   | NULL   |   |
| method  | text  | YES  |   | NULL  |    |
| json_input  | json  | YES  |   |  NULL  |   |
| json_output  | json  | YES  |   | NULL   |   |
| created_on  | datetime  | NO  |   | CURRENT_TIMESTAMP  | DEFAULT_GENERATED  |