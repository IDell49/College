Ejecutar corriendo:

go run {PRG} < {INPUT}

Donde {PRG} es el fuente del programa solicitado y {INPUT} es alguno de los archivos de entrada

Por ejemplo:

 go run replace.go < rp-input-test/rp.input6.longline.txt

Para guardar la salida se puede ejecutar con el siguiente formato:


go run {PRG} < {INPUT} > {OUTPUT}

Por ejemplo:

 go run replace.go < rp-input-test/rp.input06.longline.txt > output.txt

Para medir tiempo en un ambiente unix-like se puede ejectuar algo como:

 time go run replace.go < rp-input-test/rp.input06.verylongline.txt > output.txt

NOTA: Archivos con nombre "verylongline" pueden demorar mucho tiempo en ser procesados

