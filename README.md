Ceci est un émulateur Chip-8 écrit en Go.


Comment fonctionne l'émulateur ? : 

1. Lancez chip-8.exe
2. Mettez le nom de la ROM
   Touches : 1, 2, 3, 4, q/Q, w/W, e/R, a/A, s/S, d/D, f/F, z/Z, x/X, c/C,v/V

Fonctionnalités ✓ : 

- Affichage
- Keypad 
- Timer 

Rom ✓ :

Tests ;
- Chip8-logo   // affiche le logo de la suite de test Chip8
- 2-ibm-logo   // affiche le logo IBM
- 3-corax+     // Vérifie les opcodes
- 4-flags      // Vérfie les flags (opcodes 0x8XXX)
- 5-quirks (à part le clipping)  // Vérifie les détails correspondant à la version de Chip8
- 6-keypad     // Vérifie le clavier (opcode 0xEXXX & 0xFX0A)

Jeux :
- Breakout
- Invaders
- Particle
- Tetris

Auteurs :
- Fabie ARTHUR
- Alexandre JIN
- Mathieu DENIEUL LE DIRAISON
- Gabriel GARCIA


08/10/2023
