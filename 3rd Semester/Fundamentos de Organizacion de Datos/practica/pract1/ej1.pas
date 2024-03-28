{Realizar un algoritmo que cree un archivo de números enteros no ordenados y permita
incorporar datos al archivo. Los números son ingresados desde teclado. La carga finaliza
cuando se ingresa el número 30000, que no debe incorporarse al archivo. El nombre del
archivo debe ser proporcionado por el usuario desde teclado.}
program ej1;

type 
    numero = file of integer;

var 
    num:numero;
    nombre_fisico: string[10];
    n: integer;

begin
    write('ingrese el nombre del archivo: ');
    readln(nombre_fisico);
    assign(num, nombre_fisico);
    rewrite(num);

    writeln('ingrese un nro: ');
    readln(n);

    while(n<>30000) do begin
        writeln('ingrese un nro: ');
        readln(n);
        write(num,n);
    end;

    close(num);
end.
