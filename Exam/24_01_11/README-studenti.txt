# Laboratorio di programmazione - Istruzioni per l'esame

**NOTA BENE**: le presenti istruzioni e tutti gli script forniti si basano su sistema operativo GNU/Linux (Debian in particolare) e presumono una certa abilità dello studente nell'uso della shell (bash in particolare).

Chi fosse abituato a usare sistemi operativi diversi dovrà adattarsi usando proprie conoscenze e abilità.



## Suggerimenti

Si consiglia una attenta lettura di questo file prima dell'esame.

Si consiglia di leggere attentamente tutti i testi degli esercizi prima di iniziare il lavoro e di usare gli esempi di esecuzione per verificare di aver compreso correttamente le specifiche. 

Si consiglia di fare una valutazione della complessità degli esercizi e di farsi quindi un ordine di lavoro.

Si invita a chiedere chiarimenti in caso di dubbi sulle specifiche o sugli esiti dei test.

La maggior parte degli editor fornisce la funzionalità di "wordwrapping" dei file.



## Espandere l'archivio con le specifiche

Lanciando il comando 'tar xvf <nome del file>.tar' viene espanso l'archivio a partire dalla directory in cui ci si trova (a patto che vi sia stato scaricato il file stesso).

Verrà creata una sottodirectory per ogni esercizio con tutto il materiale occorrente per l'esercizio stesso (specifiche, test, ecc.)

Si invitano gli studenti a lavorare nelle sottodirectory così create.



## Contenuto della directory "Tema"

Troverete una serie di sottodirectory con i testi degli esercizi.
In ogni directory, in testa al file '<nome esercizio>_test.go' si trova, sotto forma di commento, il testo dell'esercizio, di seguito si trovano i test da eseguire per fare una prima (auto)valutazione del proprio elaborato (vedi più avanti su come usare i test).



## Uso dei test Go (ove presenti)

Per svolgere il compito bisogna entrare in ogni sottodirectory e creare il relativo file 'go' (se, ad esempio, il test si chiama 'filtro_test.go', va creato il file 'filtro.go'; attenzione alle maiuscole e minuscole, in generale sono caratteri tutti minuscoli o in *camelcase*) secondo le specifiche.

Per "testare" il codice basta compilare innanzitutto il proprio elaborato con 'go build', poi si può lanciare il testing con 'go test -v'.
Ricordarsi di ricompilare dopo ogni modifica prima di lanciare di nuovo i test.

Il testing produce in output un esito per ogni test, in cui:

"=== RUN", seguito sulla stessa riga dal nome di un test, indica l'inizio dell'esito relativo a quel test

"--- PASS", seguito sulla stessa riga dal nome di un test, indica che il programma 'x.go' ha superato quel test

"--- FAIL", seguito sulla stessa riga dal nome di un test, indica che il programma 'x.go' NON ha superato quel test

Alla fine dell'output:

"PASS" indica che il programma 'x.go' ha superato tutti i test contenuti in 'x_test.go'

"FAIL" indica che il programma 'x.go' NON ha superato ALMENO UNO dei test contenuti in 'x_test.go'

In caso di FAIL si consiglia VIVAMENTE di esaminare l'output dei test per capire cosa va corretto nel proprio programma.

Si consiglia VIVAMENTE inoltre di esaminare anche il codice dei test oltre al testo dell'esercizio al fine di comprendere meglio ciò che viene chiesto per lo svolgimento.

**NOTA BENE**: i test effettivi eseguiti dai docenti in fase di correzione potrebbero essere in numero e tipo diversi da quelli presentati nel tema d'esame (cioè, un PASS finale non significa automaticamente che l'esercizio sia privo di errori).
Inoltre a determinare il voto concorreranno, oltre al numero di test passati, anche altri aspetti, quali la semplicità del codice, la leggibilità, l'uso della memoria, ecc.



## Consegna

Va effettuata caricando i SINGOLI file '<nome esercizio>.go' (quindi NON vanno consegnati gli eseguibili, i test e gli altri file) su https://upload.di.unimi.it utilizzando la propria login e scegliendo la SESSIONE CORRETTA (nel dubbio chiedere ai docenti).
Consegnare in una sessione sbagliata comporta la perdita della consegna e il conseguente annullamento dell'esame.

**NOTA BENE**: si può caricare più volte lo stesso file, verrà valutata soltanto l'ultima versione depositata.
Si consiglia di caricare i file anche 'in itinere' in modo da avere un backup in caso di problemi alla postazione.

**NOTA BENE**: la sessione (login) di upload scade dopo circa 15 minuti di inattività, quindi nel caso occorre semplicemente fornire di nuovo le proprie credenziali (evitare trovarsi in questa situazione a un minuto dall'ora di consegna).

***ATTENZIONE***: la sessione d'esame si CHIUDE AUTOMATICAMENTE all'orario "di consegna" che viene comunicato in aula dai docenti.
Fino a quell'ora si possono consegnare i propri elaborati, oltre no e NON SI FARANNO ECCEZIONI.
Al termine della consegna il sistema di upload NON ACCETTA più il caricamento dei file.
Si consiglia di non ridursi all'ultimo momento.



## Per ritirarsi

Occorre caricare un file **non vuoto** dal nome 'ritirato.txt' (entro l'orario di consegna), come testo di riempimento potete usare "ritirato".



## Chi ha terminato (sia per consegna che per ritiro)

E` pregato semplicemente di scollegarsi dalla sessione di upload, senza spegnere il PC.



## Valutazione

***ATTENZIONE***: condizione necessaria affinché tutto l'elaborato venga valutato è che sia presente ALMENO UN esercizio FILTRO (indicato dal docente) che compila e gira perfettamente (cioè produca l'output atteso).

In ogni caso verranno valutati (cioè ne verrà esaminato e giudicato il sorgente) solo gli esercizi che compilano senza errori, gli esercizi che non compilano non concorreranno alla formazione del voto.

Gli esercizi **i cui test** non dovessero *compilare* verranno fortemente penalizzati d'ufficio.

Alla formazione del voto concorreranno anche i seguenti aspetti:
- che il programma mostri un funzionamento corretto su casi di test significativi
- la semplicità della soluzione e della sua implementazione
- l'uso di nomi significativi per variabili, tipi e funzioni (leggibilità)
- che nel programma sia evitata la duplicazione di codice
- la struttura del programma e l'uso di funzioni (modularità)
- l'uso della memoria, che deve essere limitato alla sola memoria necessaria.

Gli esercizi hanno pesi diversi, per cui non c'è un "numero minimo di esercizi giusti per passare l'esame".
I pesi vengono assegnati a valle di una valutazione sul campo della difficoltà di svolgimento, quindi chiedere "quali sono i pesi?" durante l'esame non può ricevere risposta.



## Materiale utilizzabile

* documentazione online di Go (https://pkg.go.dev/std)
* libro di testo o manuale di Go (eventualmente pdf) (non eserciziari)
* penna
* carta (fornita dai docenti)
* finestra per l'upload
* editor
* terminale



## Materiale NON utilizzabile

Ogni altro "oggetto" non elencato nella sezione precedente.
Quindi, ad esempio, NON si possono tenere in vista telefoni, smartwatch, appunti, altri libri, quaderni, etc., che dovranno essere riposti nelle proprie borse o giacche e appesi agli attaccapanni. 



## Comportamenti sanzionabili

Comunicare con chiunque (docenti esclusi) a esame iniziato comporta espulsione e annullamento dell'esame, per TUTTI i soggetti coinvolti nella comunicazione.

Idem in caso di utilizzo di materiale non consentito.

Se si ha bisogno di delucidazioni sul tema d'esame o per altre domande rivolgersi ai docenti (è sufficiente alzare la mano).
