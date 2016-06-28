package main
import "fmt"
func main() {
     elements := make(map[string]string)
     elements["H"] = "Hydrogen"
     elements["He"] = "Helium"
     elements["Li"] = "Lithium"
     elements["Be"] = "Beryllium"
     elements["B"] = "Boron"
     elements["C"] = "Carbon"
     elements["N"] = "Nitrogen"
     elements["O"] = "Oxygen"
     elements["F"] = "Fluorine"
     elements["Ne"] = "Neon"


     fmt.Println(elements["Li"])

     fmt.Println(elements["Un"])

     name, ok := elements["Un"]
     fmt.Println(name, ok)


    // If ok is true, then print the name and ok
     if name, ok := elements["Un"]; ok {
       fmt.Println(name, ok)
     }

  // Maps can also be stored as maps of maps
}
