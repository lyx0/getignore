package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// pLang holds the parsed flag string.
	var pLang string

	flag.StringVar(&pLang, "lang", "go", "Programming language")
	flag.Parse()

	// Check if the parsed flag is in our languages hashmap.
	for name, url := range languages {
		if string(name) == strings.ToLower(pLang) {
			if err := getFile(url); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Downloaded the .gitignore for %s from %s\n", name, url)
		}
	}

}

// getFile tries to download the body of the raw gitignore file.
func getFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
		return err
	}

	defer resp.Body.Close()

	// The body contains the content of the gitignore file.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR:", err)
		return err
	}

	// Print out file contents as a check.
	fmt.Println("File Contents:")
	fmt.Println(string(body))

	// Write the body contents to a file named ".gitignore".
	writeToFile(body)
	return nil
}

// writeToFile creates a .gitignore file and then writes
// the passed in contents to it.
func writeToFile(contents []byte) error {
	err := os.WriteFile(".gitignore", contents, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// Languages is a hashmap containing the lowercased name of the
// gitignore file as the key, and the url as the value.
// For example craftcms is the name of the CraftCMS.gitignore file.
var languages = map[string]string{
	"al":                    "https://raw.githubusercontent.com/github/gitignore/main/AL.gitignore",
	"actionscript":          "https://raw.githubusercontent.com/github/gitignore/main/Actionscript.gitignore",
	"ada":                   "https://raw.githubusercontent.com/github/gitignore/main/Ada.gitignore",
	"agda":                  "https://raw.githubusercontent.com/github/gitignore/main/Agda.gitignore",
	"android":               "https://raw.githubusercontent.com/github/gitignore/main/Android.gitignore",
	"appengine":             "https://raw.githubusercontent.com/github/gitignore/main/AppEngine.gitignore",
	"appceleratortitanium":  "https://raw.githubusercontent.com/github/gitignore/main/AppceleratorTitanium.gitignore",
	"archlinuxpackages":     "https://raw.githubusercontent.com/github/gitignore/main/ArchLinuxPackages.gitignore",
	"autotools":             "https://raw.githubusercontent.com/github/gitignore/main/Autotools.gitignore",
	"c++":                   "https://raw.githubusercontent.com/github/gitignore/main/C++..gitignore",
	"c":                     "https://raw.githubusercontent.com/github/gitignore/main/C.gitignore",
	"cfwheels":              "https://raw.githubusercontent.com/github/gitignore/main/CFWheels.gitignore",
	"cmake":                 "https://raw.githubusercontent.com/github/gitignore/main/CMake.gitignore",
	"cuda":                  "https://raw.githubusercontent.com/github/gitignore/main/CUDA.gitignore",
	"cakephp":               "https://raw.githubusercontent.com/github/gitignore/main/CakePHP.gitignore",
	"chefcookbook":          "https://raw.githubusercontent.com/github/gitignore/main/ChefCookbook.gitignore",
	"clojure":               "https://raw.githubusercontent.com/github/gitignore/main/Clojure.gitignore",
	"codeigniter":           "https://raw.githubusercontent.com/github/gitignore/main/CodeIgniter.gitignore",
	"commonlisp":            "https://raw.githubusercontent.com/github/gitignore/main/CommonLisp.gitignore",
	"composer":              "https://raw.githubusercontent.com/github/gitignore/main/Composer.gitignore",
	"concrete5":             "https://raw.githubusercontent.com/github/gitignore/main/Concrete5.gitignore",
	"coq":                   "https://raw.githubusercontent.com/github/gitignore/main/Coq.gitignore",
	"craftcms":              "https://raw.githubusercontent.com/github/gitignore/main/CraftCMS.gitignore",
	"d":                     "https://raw.githubusercontent.com/github/gitignore/main/D.gitignore",
	"dm":                    "https://raw.githubusercontent.com/github/gitignore/main/DM.gitignore",
	"dart":                  "https://raw.githubusercontent.com/github/gitignore/main/Dart.gitignore",
	"delphi":                "https://raw.githubusercontent.com/github/gitignore/main/Delphi.gitignore",
	"drupal":                "https://raw.githubusercontent.com/github/gitignore/main/Drupal.gitignore",
	"episerver":             "https://raw.githubusercontent.com/github/gitignore/main/EPiServer.gitignore",
	"eagle":                 "https://raw.githubusercontent.com/github/gitignore/main/Eagle.gitignore",
	"elisp":                 "https://raw.githubusercontent.com/github/gitignore/main/Elisp.gitignore",
	"elixir":                "https://raw.githubusercontent.com/github/gitignore/main/Elixir.gitignore",
	"elm":                   "https://raw.githubusercontent.com/github/gitignore/main/Elm.gitignore",
	"erlang":                "https://raw.githubusercontent.com/github/gitignore/main/Erlang.gitignore",
	"expressionengine":      "https://raw.githubusercontent.com/github/gitignore/main/ExpressionEngine.gitignore",
	"extjs":                 "https://raw.githubusercontent.com/github/gitignore/main/ExtJs.gitignore",
	"fancy":                 "https://raw.githubusercontent.com/github/gitignore/main/Fancy.gitignore",
	"finale":                "https://raw.githubusercontent.com/github/gitignore/main/Finale.gitignore",
	"flaxengine":            "https://raw.githubusercontent.com/github/gitignore/main/FlaxEngine.gitignore",
	"forcedotcom":           "https://raw.githubusercontent.com/github/gitignore/main/ForceDotCom.gitignore",
	"fortran":               "https://raw.githubusercontent.com/github/gitignore/main/Fortran.gitignore",
	"fuelphp":               "https://raw.githubusercontent.com/github/gitignore/main/FuelPHP.gitignore",
	"gwt":                   "https://raw.githubusercontent.com/github/gitignore/main/GWT.gitignore",
	"gcov":                  "https://raw.githubusercontent.com/github/gitignore/main/Gcov.gitignore",
	"gitbook":               "https://raw.githubusercontent.com/github/gitignore/main/GitBook.gitignore",
	"go":                    "https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore",
	"godot":                 "https://raw.githubusercontent.com/github/gitignore/main/Godot.gitignore",
	"gradle":                "https://raw.githubusercontent.com/github/gitignore/main/Gradle.gitignore",
	"grails":                "https://raw.githubusercontent.com/github/gitignore/main/Grails.gitignore",
	"haskell":               "https://raw.githubusercontent.com/github/gitignore/main/Haskell.gitignore",
	"igorpro":               "https://raw.githubusercontent.com/github/gitignore/main/IGORPro.gitignore",
	"idris":                 "https://raw.githubusercontent.com/github/gitignore/main/Idris.gitignore",
	"jboss":                 "https://raw.githubusercontent.com/github/gitignore/main/JBoss.gitignore",
	"jenkins_home":          "https://raw.githubusercontent.com/github/gitignore/main/JENKINS_HOME.gitignore",
	"java":                  "https://raw.githubusercontent.com/github/gitignore/main/Java.gitignore",
	"jekyll":                "https://raw.githubusercontent.com/github/gitignore/main/Jekyll.gitignore",
	"joomla":                "https://raw.githubusercontent.com/github/gitignore/main/Joomla.gitignore",
	"julia":                 "https://raw.githubusercontent.com/github/gitignore/main/Julia.gitignore",
	"kicad":                 "https://raw.githubusercontent.com/github/gitignore/main/KiCad.gitignore",
	"kohana":                "https://raw.githubusercontent.com/github/gitignore/main/Kohana.gitignore",
	"kotlin":                "https://raw.githubusercontent.com/github/gitignore/main/Kotlin.gitignore",
	"labview":               "https://raw.githubusercontent.com/github/gitignore/main/LabVIEW.gitignore",
	"laravel":               "https://raw.githubusercontent.com/github/gitignore/main/Laravel.gitignore",
	"leiningen":             "https://raw.githubusercontent.com/github/gitignore/main/Leiningen.gitignore",
	"lemonstand":            "https://raw.githubusercontent.com/github/gitignore/main/LemonStand.gitignore",
	"lilypond":              "https://raw.githubusercontent.com/github/gitignore/main/Lilypond.gitignore",
	"lithium":               "https://raw.githubusercontent.com/github/gitignore/main/Lithium.gitignore",
	"lua":                   "https://raw.githubusercontent.com/github/gitignore/main/Lua.gitignore",
	"magento":               "https://raw.githubusercontent.com/github/gitignore/main/Magento.gitignore",
	"maven":                 "https://raw.githubusercontent.com/github/gitignore/main/Maven.gitignore",
	"mercury":               "https://raw.githubusercontent.com/github/gitignore/main/Mercury.gitignore",
	"metaprogrammingsystem": "https://raw.githubusercontent.com/github/gitignore/main/MetaProgrammingSystem.gitignore",
	"nanoc":                 "https://raw.githubusercontent.com/github/gitignore/main/Nanoc.gitignore",
	"nim":                   "https://raw.githubusercontent.com/github/gitignore/main/Nim.gitignore",
	"node":                  "https://raw.githubusercontent.com/github/gitignore/main/Node.gitignore",
	"ocaml":                 "https://raw.githubusercontent.com/github/gitignore/main/OCaml.gitignore",
	"objective-c":           "https://raw.githubusercontent.com/github/gitignore/main/Objective-C.gitignore",
	"opa":                   "https://raw.githubusercontent.com/github/gitignore/main/Opa.gitignore",
	"opencart":              "https://raw.githubusercontent.com/github/gitignore/main/OpenCart.gitignore",
	"oracleforms":           "https://raw.githubusercontent.com/github/gitignore/main/OracleForms.gitignore",
	"packer":                "https://raw.githubusercontent.com/github/gitignore/main/Packer.gitignore",
	"perl":                  "https://raw.githubusercontent.com/github/gitignore/main/Perl.gitignore",
	"phalcon":               "https://raw.githubusercontent.com/github/gitignore/main/Phalcon.gitignore",
	"playframework":         "https://raw.githubusercontent.com/github/gitignore/main/PlayFramework.gitignore",
	"plone":                 "https://raw.githubusercontent.com/github/gitignore/main/Plone.gitignore",
	"prestashop":            "https://raw.githubusercontent.com/github/gitignore/main/Prestashop.gitignore",
	"processing":            "https://raw.githubusercontent.com/github/gitignore/main/Processing.gitignore",
	"purescript":            "https://raw.githubusercontent.com/github/gitignore/main/PureScript.gitignore",
	"python":                "https://raw.githubusercontent.com/github/gitignore/main/Python.gitignore",
	"qooxdoo":               "https://raw.githubusercontent.com/github/gitignore/main/Qooxdoo.gitignore",
	"qt":                    "https://raw.githubusercontent.com/github/gitignore/main/Qt.gitignore",
	"r":                     "https://raw.githubusercontent.com/github/gitignore/main/R.gitignore",
	"ros":                   "https://raw.githubusercontent.com/github/gitignore/main/ROS.gitignore",
	"racket":                "https://raw.githubusercontent.com/github/gitignore/main/Racket.gitignore",
	"rails":                 "https://raw.githubusercontent.com/github/gitignore/main/Rails.gitignore",
	"raku":                  "https://raw.githubusercontent.com/github/gitignore/main/Raku.gitignore",
	"rhodesrhomobile":       "https://raw.githubusercontent.com/github/gitignore/main/RhodesRhomobile.gitignore",
	"ruby":                  "https://raw.githubusercontent.com/github/gitignore/main/Ruby.gitignore",
	"rust":                  "https://raw.githubusercontent.com/github/gitignore/main/Rust.gitignore",
	"scons":                 "https://raw.githubusercontent.com/github/gitignore/main/SCons.gitignore",
	"sass":                  "https://raw.githubusercontent.com/github/gitignore/main/Sass.gitignore",
	"scala":                 "https://raw.githubusercontent.com/github/gitignore/main/Scala.gitignore",
	"scheme":                "https://raw.githubusercontent.com/github/gitignore/main/Scheme.gitignore",
	"scrivener":             "https://raw.githubusercontent.com/github/gitignore/main/Scrivener.gitignore",
	"sdcc":                  "https://raw.githubusercontent.com/github/gitignore/main/Sdcc.gitignore",
	"seamgen":               "https://raw.githubusercontent.com/github/gitignore/main/SeamGen.gitignore",
	"sketchup":              "https://raw.githubusercontent.com/github/gitignore/main/SketchUp.gitignore",
	"smalltalk":             "https://raw.githubusercontent.com/github/gitignore/main/Smalltalk.gitignore",
	"stella":                "https://raw.githubusercontent.com/github/gitignore/main/Stella.gitignore",
	"sugarcrm":              "https://raw.githubusercontent.com/github/gitignore/main/SugarCRM.gitignore",
	"swift":                 "https://raw.githubusercontent.com/github/gitignore/main/Swift.gitignore",
	"symfony":               "https://raw.githubusercontent.com/github/gitignore/main/Symfony.gitignore",
	"symphonycms":           "https://raw.githubusercontent.com/github/gitignore/main/SymphonyCMS.gitignore",
	"tex":                   "https://raw.githubusercontent.com/github/gitignore/main/TEX.gitignore",
	"terraform":             "https://raw.githubusercontent.com/github/gitignore/main/Terraform.gitignore",
	"textpattern":           "https://raw.githubusercontent.com/github/gitignore/main/Textpattern.gitignore",
	"turbogears2":           "https://raw.githubusercontent.com/github/gitignore/main/TurboGears2.gitignore",
	"twincat3":              "https://raw.githubusercontent.com/github/gitignore/main/TwinCAT3.gitignore",
	"typo3":                 "https://raw.githubusercontent.com/github/gitignore/main/Typo3.gitignore",
	"unity":                 "https://raw.githubusercontent.com/github/gitignore/main/Unity.gitignore",
	"unrealengine":          "https://raw.githubusercontent.com/github/gitignore/main/UnrealEngine.gitignore",
	"vvvv":                  "https://raw.githubusercontent.com/github/gitignore/main/VVVV.gitignore",
	"visualstudio":          "https://raw.githubusercontent.com/github/gitignore/main/VisualStudio.gitignore",
	"waf":                   "https://raw.githubusercontent.com/github/gitignore/main/Waf.gitignore",
	"wordpress":             "https://raw.githubusercontent.com/github/gitignore/main/WordPress.gitignore",
	"xojo":                  "https://raw.githubusercontent.com/github/gitignore/main/Xojo.gitignore",
	"yeoman":                "https://raw.githubusercontent.com/github/gitignore/main/Yeoman.gitignore",
	"yii":                   "https://raw.githubusercontent.com/github/gitignore/main/Yii.gitignore",
	"zendframework":         "https://raw.githubusercontent.com/github/gitignore/main/ZendFramework.gitignore",
	"zephir":                "https://raw.githubusercontent.com/github/gitignore/main/Zephir.gitignore",
}
