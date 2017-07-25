# TopanTerminalBrowser

Browse MeleteToPan generated theta.csv to explore Manhattan distances between passages of interest.

## Run the example

1. Clone the repo and open a terminal in the cloned folder.
2. On Mac or AMD-Linux run `./tbb`, on Windows run the win-exe, and if you have a 386-Linux system run the lin386-file (if you run it on a UNIX system you might have to chmod +x the executable before you can use it)
3. TTB reads in the theta.csv file that is in the same folder and presents the space generated by the topic model to you
4. Enter the identifier for the passage you are interested in (or hit enter to choose the suggested random passage)
5. Enter how many similar passages you want to generate.
6. Included in the sample is a LDAvis of the theta.csv used. You can open the `index.html` of [this folder] in your favourite browser to explore the topic model more thoroughly. Combine this approach with TBB to get the best experience.
7. Happy navigating!

## Example Terminal Output

<details><summary>When TBB starts up</summary>
```
Locals-MacBook-Pro:TopanTerminalBrowser koentges$ ./tbb 
Reading file.
All is read.
Browsing through the space of the following topic dimensions:
Topic 1 ἡμῶν_ὑμῖν_ἡμᾶς_ἡμῖν_ὑμεῖς_ὑμᾶς_ἃ
Topic 2 ἡμῖν_ἡμᾶς_γε_ἡμῶν_μόνον_οὐδ_ἔσται
Topic 3 πρέσβεις_λακεδαιμονίους_λακεδαιμονίων_ξυμμαχίαν_λακεδαιμόνιοι_σφίσι_σπονδὰς
Topic 4 νῆες_νεῶν_εἴκοσι_πέντε_πελοποννήσιοι_δύο_δέκα
Topic 5 εὖ_δι_ἐξ_αἰεὶ_ὄντες_καλῶς_πλεῖστα
Topic 6 δὴ_πρὸ_ἐγένετο_πολέμου_γε_τοῦδε_χρόνον
Topic 7 νεῶν_ναυσὶν_νῆες_χίον_δέκα_χίων_χῖοι
Topic 8 σφῶν_ὅπως_πρέσβεις_ἔς_ξυμμάχους_χρήματα_πέμπειν
Topic 9 εἴ_τις_ἄν_οὐδ_δέ_τρόπῳ_κατ
Topic 10 αὐτὸν_αὐτὸς_αὐτῷ_ὢν_ἀθηναίοις_πρότερον_τισσαφέρνην
Topic 11 ἡμῖν_ἢν_ἧσσον_ᾧ_οὐχ_ἡμῶν_πλέον
Topic 12 τοιάδε_τοιαῦτα_αὐτὸς_εἶπεν_εἴ_εἶπον_λέγειν
Topic 13 κέρας_πολὺ_συρακοσίων_συρακόσιοι_στράτευμα_δεξιὸν_ξύμμαχοι
Topic 14 ὑμᾶς_ἡμᾶς_ἡμεῖς_ὑμῖν_ὦ_ὑμῶν_ὑπ
Topic 15 ἢν_λακεδαιμόνιοι_λακεδαιμονίοις_λακεδαιμονίων_σπονδὰς_δέ_ξυμμάχοις
Topic 16 συρακοσίων_συρακόσιοι_ὅπως_ναυσὶ_αὖθις_ᾗ_πρότερον
Topic 17 βρασίδας_αὐτὸς_αὐτῷ_ἔχων_τότε_ἑαυτοῦ_θρᾴκης
Topic 18 ᾗ_ὅπως_τείχους_πόλεως_τεῖχος_εἴ_τείχει
Topic 19 σφίσι_τούς_οὐχ_σφῶν_ἐκείνων_νομίζοντες_δὴ
Topic 20 τις_παρὰ_σφᾶς_οὐδ_ἔσεσθαι_πρὶν_ἔργῳ
Topic 21 ὄντες_ἑλλήνων_ὅσοι_ἄλλων_βαρβάρων_ποτε_ἕλληνες
Topic 22 ἄλλοι_ὁπλῖται_πάντες_μετ_ξύμμαχοι_οἷς_ἦλθον
Topic 23 καθ_ἄλλους_τήν_εἴ_πάλιν_κινδύνου_πάντα
Topic 24 θαλάσσης_ναυτικὸν_γῆς_πρότερον_νεῶν_πόλεμον_γῆν
Topic 25 πόλις_ἐξ_μέχρι_πρότερον_θαλάσσης_ποταμοῦ_ἀπ
Topic 26 μήτε_ξυμμάχους_ὅπλα_σφᾶς_λακεδαιμόνιοι_ἑαυτῶν_αὐτούς
Topic 27 θέρους_λακεδαιμόνιοι_ἐπιγιγνομένου_λακεδαιμονίων_χειμῶνος_ξύμμαχοι_πάλιν
Topic 28 γῆν_οἴκου_ἀνεχώρησαν_γῆς_ἐδῄουν_χρόνον_ἀφίκοντο
Topic 29 πόλεως_πόλει_πάντα_ἃ_ἄλλα_ὅπερ_τις
Topic 30 βασιλέως_ἢν_ἔφη_τις_δέ_πόλεις_βασιλεὺς
Topic 31 αὐτὸν_αὐτῷ_βασιλέα_παρὰ_δὴ_δι_λόγοις
Topic 32 πλείους_οὐδὲν_ἄλλοι_μέρος_τινας_ὁπλιτῶν_ἐλάσσους
Topic 33 οὐχ_αἰεὶ_πλεῖστον_ᾗ_οὔσης_πολὺ_γνώμῃ
Topic 34 πόλις_τινα_σφίσιν_καί_ἕκαστος_σφίσι_κατ
Topic 35 ὁπλίτας_ξυμμάχων_ἱππέας_χιλίους_νεῶν_εἴκοσι_τριάκοντα
Topic 36 τείχη_θάλασσαν_γῆν_εἶχον_πόλεως_στρατηγοῦντος_πάλιν
Topic 37 πλέον_δι_πόλει_ἐξ_πόλεις_ὅπερ_λόγου
Topic 38 αὐτῷ_ἔδει_δημοσθένης_ναύπακτον_ἀκαρνᾶνες_χωρίων_ναυσὶν
Topic 39 ἄλλα_πολλὰ_ὅσα_τότε_ἃ_χρήματα_τούτων
Topic 40 νήσῳ_λακεδαιμόνιοι_ἄνδρας_παρὰ_πύλον_ἀθηναίοις_μυτιληναίων
Topic 41 δῆμον_σφῶν_οὐδὲν_πράγματα_πολλοὺς_τἆλλα_ἄρχειν
Topic 42 νῆες_τεσσαράκοντα_δύο_πέντε_ἔτυχον_ναυσὶ_τριάκοντα
Topic 43 συρακοσίων_σικελίας_σικελίαν_στρατιᾷ_σικελῶν_συρακοσίοις_συρακούσας
Topic 44 τροπαῖον_νεκροὺς_ἔστησαν_ἔλαβον_μάχῃ_τούς_ὑποσπόνδους
Topic 45 πολέμῳ_ἐτελεύτα_ἐγένετο_τῷδε_ἔτος_θέρος_αὕτη
Topic 46 πόλεως_σταδίους_τεῖχος_βοιωτοὶ_μάχην_νυκτὸς_μάχης
Topic 47 ἔπειτα_ὅμως_μέν_πόλεως_πρὶν_ἄλλο_πόλεμον
Topic 48 κορινθίων_κερκυραῖοι_ἐκέλευον_κορινθίοις_σφίσιν_κορίνθιοι_πόλει
Topic 49 τρόπῳ_ἔξω_δέ_τινες_ἀπόλλωνος_πολλὰ_ὃ
Topic 50 ὁπλίταις_αὐτῷ_ξυμμάχων_τρόπῳ_νικίας_τούτῳ_ἑαυτῶν
-------------------------------
Significant distance has been set to:  0.1
Happy navigating!
```
</details>
<details><summary>User input</summary>
```
Enter URN (e.g.urn:cts:greekLit:tlg0003.tlg001.perseus-grc2:2.2.4):
Enter number of similar passages(e.g. 3):3
```
</details>

<details><summary>Results</summary>

------------------------------------------
You queried:
URN: urn:cts:greekLit:tlg0003.tlg001.perseus-grc2:2.2.4
Text: θέμενοι δὲ ἐς τὴν ἀγορὰν τὰ ὅπλα τοῖς μὲν ἐπαγαγομένοις οὐκ ἐπείθοντο ὥστε εὐθὺς ἔργου ἔχεσθαι καὶ ἰέναι ἐπὶ τὰς οἰκίας τῶν ἐχθρῶν, γνώμην δ᾽ ἐποιοῦντο κηρύγμασί τε χρήσασθαι ἐπιτηδείοις καὶ ἐς ξύμβασιν μᾶλλον καὶ φιλίαν τὴν πόλιν ἀγαγεῖν ʽκαὶ ἀνεῖπεν ὁ κῆρυξ, εἴ τις βούλεται κατὰ τὰ πάτρια τῶν πάντων Βοιωτῶν ξυμμαχεῖν, τίθεσθαι παρ᾽ αὑτοὺς τὰ ὅπλἀ, νομίζοντες σφίσι ῥᾳδίως τούτῳ τῷ τρόπῳ προσχωρήσειν τὴν πόλιν.

Three most important topics:
Topic 49 τρόπῳ_ἔξω_δέ_τινες_ἀπόλλωνος_πολλὰ_ὃ : 0.457714285714286
Topic 19 σφίσι_τούς_οὐχ_σφῶν_ἐκείνων_νομίζοντες_δὴ : 0.172
Topic 47 ἔπειτα_ὅμως_μέν_πόλεως_πρὶν_ἄλλο_πόλεμον : 0.143428571428571
------------------------------------------
------------------------------------------
Result: 1
Distance: 0.7390476190476194
URN: urn:cts:greekLit:tlg0003.tlg001.perseus-grc2:4.97.3
Text: πᾶσι γὰρ εἶναι καθεστηκὸς ἰόντας ἐπὶ τὴν ἀλλήλων ἱερῶν τῶν ἐνόντων ἀπέχεσθαι, Ἀθηναίους δὲ Δήλιον τειχίσαντας ἐνοικεῖν, καὶ ὅσα ἄνθρωποι ἐν βεβήλῳ δρῶσι πάντα γίγνεσθαι αὐτόθι, ὕδωρ τε ὃ ἦν ἄψαυστον σφίσι πλὴν πρὸς τὰ ἱερὰ χέρνιβι χρῆσθαι, ἀνασπάσαντας ὑδρεύεσθαι:

Three most important topics:
Topic 49 τρόπῳ_ἔξω_δέ_τινες_ἀπόλλωνος_πολλὰ_ὃ : 0.524761904761905
Topic 47 ἔπειτα_ὅμως_μέν_πόλεως_πρὶν_ἄλλο_πόλεμον : 0.286666666666667
Topic 15 ἢν_λακεδαιμόνιοι_λακεδαιμονίοις_λακεδαιμονίων_σπονδὰς_δέ_ξυμμάχοις : 0.0485714285714286

Topics with significant distance:
Distance Topic 19 σφίσι_τούς_οὐχ_σφῶν_ἐκείνων_νομίζοντες_δὴ : 0.17104761904761903
Distance Topic 47 ἔπειτα_ὅμως_μέν_πόλεως_πρὶν_ἄλλο_πόλεμον : 0.14323809523809602
------------------------------------------
------------------------------------------
Result: 2
Distance: 0.8182857142857145
URN: urn:cts:greekLit:tlg0003.tlg001.perseus-grc2:8.95.2
Text: Ἀθηναῖοι δὲ κατὰ τάχος καὶ ἀξυγκροτήτοις πληρώμασιν ἀναγκασθέντες χρήσασθαι, οἷα πόλεώς τε στασιαζούσης καὶ περὶ τοῦ μεγίστου ἐν τάχει βουλόμενοι βοηθῆσαι ʽΕὔβοια γὰρ αὐτοῖς ἀποκεκλῃμένης τῆς Ἀττικῆς πάντα ἦν̓, πέμπουσι Θυμοχάρη στρατηγὸν καὶ ναῦς ἐς Ἐρετρίαν,

Three most important topics:
Topic 49 τρόπῳ_ἔξω_δέ_τινες_ἀπόλλωνος_πολλὰ_ὃ : 0.358571428571429
Topic 28 γῆν_οἴκου_ἀνεχώρησαν_γῆς_ἐδῄουν_χρόνον_ἀφίκοντο : 0.215714285714286
Topic 17 βρασίδας_αὐτὸς_αὐτῷ_ἔχων_τότε_ἑαυτοῦ_θρᾴκης : 0.144285714285714

Topics with significant distance:
Distance Topic 17 βρασίδας_αὐτὸς_αὐτῷ_ἔχων_τότε_ἑαυτοῦ_θρᾴκης : 0.1437142857142854
Distance Topic 28 γῆν_οἴκου_ἀνεχώρησαν_γῆς_ἐδῄουν_χρόνον_ἀφίκοντο : 0.2151428571428574
Distance Topic 47 ἔπειτα_ὅμως_μέν_πόλεως_πρὶν_ἄλλο_πόλεμον : 0.14199999999999957
------------------------------------------
------------------------------------------
Result: 3
Distance: 0.8274285714285718
URN: urn:cts:greekLit:tlg0003.tlg001.perseus-grc2:7.87.2
Text: πάντα τε ποιούντων αὐτῶν διὰ στενοχωρίαν ἐν τῷ αὐτῷ καὶ προσέτι τῶν νεκρῶν ὁμοῦ ἐπ᾽ ἀλλήλοις ξυννενημένων, οἳ ἔκ τε τῶν τραυμάτων καὶ διὰ τὴν μεταβολὴν καὶ τὸ τοιοῦτον ἀπέθνῃσκον, καὶ ὀσμαὶ ἦσαν οὐκ ἀνεκτοί, καὶ λιμῷ ἅμα καὶ δίψῃ ἐπιέζοντο ʽἐδίδοσαν γὰρ αὐτῶν ἑκάστῳ ἐπὶ ὀκτὼ μῆνας κοτύλην ὕδατος καὶ δύο κοτύλας σίτοὐ, ἄλλα τε ὅσα εἰκὸς ἐν τῷ τοιούτῳ χωρίῳ ἐμπεπτωκότας κακοπαθῆσαι, οὐδὲν ὅτι οὐκ ἐπεγένετο αὐτοῖς:

Three most important topics:
Topic 49 τρόπῳ_ἔξω_δέ_τινες_ἀπόλλωνος_πολλὰ_ὃ : 0.600666666666667
Topic 35 ὁπλίτας_ξυμμάχων_ἱππέας_χιλίους_νεῶν_εἴκοσι_τριάκοντα : 0.134
Topic 10 αὐτὸν_αὐτὸς_αὐτῷ_ὢν_ἀθηναίοις_πρότερον_τισσαφέρνην : 0.100666666666667

Topics with significant distance:
Distance Topic 10 αὐτὸν_αὐτὸς_αὐτῷ_ὢν_ἀθηναίοις_πρότερον_τισσαφέρνην : 0.10009523809523843
Distance Topic 19 σφίσι_τούς_οὐχ_σφῶν_ἐκείνων_νομίζοντες_δὴ : 0.1713333333333333
Distance Topic 35 ὁπλίτας_ξυμμάχων_ἱππέας_χιλίους_νεῶν_εἴκοσι_τριάκοντα : 0.13342857142857142
Distance Topic 49 τρόπῳ_ἔξω_δέ_τινες_ἀπόλλωνος_πολλὰ_ὃ : 0.142952380952381
------------------------------------------
</details>

## Run your own

1. Generate a Topic Model of your text with [MeleteToPan](https://github.com/ThomasK81/ToPan)
2. Identify the correct model folder in the subfolder www/data. Copy the theta.csv of your model.
3. Paste it into the folder where TBB is placed.
4. Run TBB and explore your data. 
3a-4a. Alternatively, you can always copy&paste the TBB executable 
5. Happy navigating!
