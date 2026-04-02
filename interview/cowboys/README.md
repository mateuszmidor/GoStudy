# Zadanie „Cowboy” – symulator gry z współbieżnością

- Wyobraź sobie scenę z westernu – saloon pełen pijanych, wściekłych kowbojów.
- Mamy zestaw kowbojów.
- Każdy kowboj ma unikalne imię, punkty zdrowia oraz punkty obrażeń.
- Wszyscy kowboje zaczynają jednocześnie (równolegle), wybierają losowy cel i strzelają.
- Od punktów zdrowia celu odejmujemy punkty obrażeń strzelca.
- Jeśli punkty zdrowia celu spadną do zera lub poniżej, kowboj ginie.
- Kowboje nie strzelają do siebie ani do martwych kowbojów.
- Po oddaniu strzału strzelec „usypia” na 1 sekundę.
- Na końcu „może pozostać tylko jeden”.
- Ostatni żywy kowboj zostaje zwycięzcą.