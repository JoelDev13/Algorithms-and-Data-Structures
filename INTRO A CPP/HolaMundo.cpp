#include <iostream>

int main(){

    int primerNumero;
    int segundoNumero;

    /* output
    std::cout <<"Hola Mundo" << std::endl;*/

    // input
    std::cout <<"Ingrese el primer numero" << std::endl;
    std::cin >> primerNumero;

    std::cout <<"Ingrese el segundo numero" << std::endl;
    std::cin >> segundoNumero;

    int resultado = primerNumero + segundoNumero;
    std::cout << "La suma " << primerNumero << " + " << segundoNumero << " Es " << resultado << std::endl;

    return 0;


}