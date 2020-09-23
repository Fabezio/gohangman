# gohangman

Célèbre jeu du pendu entièrement rédigé en golang, **gohangman** est pour le moment réservé au terminal. 
  Les règles sont des plus simples: 
  - il s'agit de deviner un mot en énumérant les letres une par une
  - un échafaud est *dessiné dans le vide*
  - lorsqu'une lettre est proposée:
    - à chaque erreur, une partie de l'échafaud est ajoutée, 
    - sinon le mot se remplit peu à peu
  - la partie est terminée lorsque:
    - l'échaufaud est terminé, vous avez alors perdu, vous ne pouvez plus proposer de lettre
    - le mot est clair, vous avez gagné!
  - Vous pouvez relancer une partie quend vous voulez

## Comment ça marche
C'est tout simple: vous lancez la partie et vous proposez des lettres une par une au clavier. Evidemment les pénalités se dessineront sous forme d'échafaud dans la console. Notez que si les propositions sont supérieures à une lettre, il n'y a pas de pénalité. Attention: toute entrée est définitive, pas de rattrapage ou annulation. 
