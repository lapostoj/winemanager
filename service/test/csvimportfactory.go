package test

// ACsvImportFile returns a string matching a CSV file for import of wines
func ACsvImportFile() string {
	return `Nom,Appellation,Cru,Millésime,Région,Couleur,Type,Stock,Producteur,Origine,Format,
Riesling - Vendanges Tardives,Riesling,,1990,Alsace,Blanc,Liquoreux,1,,,75,
Château Chalon,Château-Chalon,,1992,Jura,Blanc,,1,,,62,
Vin Jaune d'Arbois,Arbois,,1993,Jura,Blanc,Liquoreux,1,,,62,`
}
