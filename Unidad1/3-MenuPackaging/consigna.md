Menu-packaging test
===================

La empresa "Labora" requiere el desarrollo de una interfaz de usuario completa que permita ejecutar diversas tareas necesarias para su eficiente funcionamiento. Para ello, la interfaz debe validar a los usuarios con roles de Administrador y Supervisor.

El Administrador tiene privilegios para crear y eliminar “laborers”, así como para generar archivos de texto con mensajes personalizados. Cada vez que se crea un "laborer", se almacenará en un array la palabra "laborer creado", junto con un contador que indica cuántas veces se ejecutó ese comando. De manera similar, cuando se elimina un "laborer", se registrará la frase "laborer eliminado" con un contador correspondiente.

Las credenciales para el Administrador son:

Usuario: admin
Contraseña: root

Las credenciales para el Supervisor son:

Usuario: seeker223
Contraseña: seekr

El Supervisor tiene la opción en su menú de "crear registro de administrador". Al seleccionar esta opción, se creará un archivo de texto que contendrá las palabras "laborer creado" y "laborer eliminado", junto con sus respectivos contadores.

El programa debe permitir el cambio de usuarios sin necesidad de interrumpir su ejecución. Asimismo, se debe proporcionar una opción para detener el programa antes de iniciar sesión.
