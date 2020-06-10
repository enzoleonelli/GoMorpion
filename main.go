/*package main

import "github.com/tadvi/winc"

func main() {
    mainWindow := winc.NewForm(nil)
    mainWindow.SetSize(400, 300)
    mainWindow.SetText("Hellow World Demo")
    btn := winc.NewPushButton(mainWindow)
    btn.SetText("Fermer")
    btn.SetPos(150, 200)
    btn.SetSize(100, 40)
    btn.OnClick().Bind(wndOnClose)
    mainWindow.Center()
    mainWindow.Show()
    mainWindow.OnClose().Bind(wndOnClose)
    winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
    winc.Exit()
}*/

package main

import (
	"strconv"
	"fmt"
	"github.com/tadvi/winc"
)

func main() {
	
	var PositionX = 160
	var PositionY = 160
	var NomJoueur = "X"
	var Text = "C'est au tour de la personne jouant : " + NomJoueur
	
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(800, 800)
	mainWindow.SetText("Morpion Go")

	txt := winc.NewLabel(mainWindow)
	txt.SetText(Text)
	txt.SetSize(1500, 100)

	tab := make(map[string]string)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			nom := strconv.Itoa(Case)
			newButton.SetSize(120, 120)
			newButton := winc.NewPushButton(mainWindow)
			newButton.SetPos(PositionY+y*120, PositionX+x*120)
			Case = Case + 1
			newButton.SetText(nom)

			newButton.OnClick().Bind(func(e *winc.Event) {
				if tab[nom] == "" {
					newButton.SetText(Joueur)
					tab[nom] = NomJoueur

					if NomJoueur == "X" {
						NomJoueur = "O"
					} else {
						NomJoueur = "X"
					}

					Text = "A vous de jouer, joueur ayant : " + NomJoueur
					txt.SetText(Text)
				}
			})
		}
	}
	
	//Création d'un tableau de gain pour le comparer au tableau du jeux en cours (Pas fonctionnel)
	func gagner() bool {

		
		Gain := [][tab]bool{
		{
			true, true, true,
			false, false, false,
			false, false, false},
	
		{
			false, false, true,
			false, false, true,
			false, false, true},
		{
			false, false, false,
			false, false, false,
			true, true, true},
		{
			true, false, false,
			true, false, false,
			true, false, false},
		{
			true, false, false,
			false, true, false,
			false, false, true},
		{
			false, false, true,
			false, true, false,
			true, false, false},
		{
			false, true, false,
			false, true, false,
			false, true, false}}
	
		var Gain [tab]bool
	
		identity := 0 
		for _, gagner := range Gain {
			for i := 0; i < len(Gain); i++ {
				if Gain[i] == true && Gain[i] == gagner[i] { 
					identity++
					if identity == 3 { 
						return true
					}
				}
			}
			identity = 0 
		}
		return false
	}

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}