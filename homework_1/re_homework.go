package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unicode/utf8"
)

// draw a main box
func drawBox(_screen tcell.Screen, _x1, _y1, _x2, _y2 int, _style tcell.Style) {
	_screen.SetContent(_x1, _y1, '┌', nil, _style)
	_screen.SetContent(_x2, _y1, '┐', nil, _style)
	_screen.SetContent(_x1, _y2, '└', nil, _style)
	_screen.SetContent(_x2, _y2, '┘', nil, _style)

	//drawing a horizontal lines
	for _x := _x1 + 1; _x < _x2; _x++ {
		_screen.SetContent(_x, _y1, '─', nil, _style)
		_screen.SetContent(_x, _y2, '─', nil, _style)
	}

	//do the same but for vertical lines
	for _y := _y1 + 1; _y < _y2; _y++ {
		_screen.SetContent(_x1, _y, '│', nil, _style)
		_screen.SetContent(_x2, _y, '│', nil, _style)
	}
}

// so as I understood from the tcell documentation, functino SetContent() accept 5 arguements the first is x-axis, that is we actually say "print each symbol in range of given text in x-axis"
// for example  "print hello world) when i is equal to it will print p so p = 1, then r, r = 2, and so on
// the second arguement do the same but for y-axis, and initially the x-axis letters types first and only then y-axis like step over step
// the third arguement is our c that is a symbol 
// the fourth is special symbols like start or the circle,  i don't need it so i set it nil
// and style it's a style we set later

func drawText(_screen tcell.Screen, _x, _y int, _text string, _style tcell.Style) {
	for _i, _c := range _text {
		_screen.SetContent(_x+_i, _y, _c, nil, _style)
	}
}

// function to write an error in file
func write_log(_inputText string) {
	_date := time.Now()
	_format_date := _date.Format("2006-01-02 15:04:05")
	// Adding a error message, bu the way i don't think that this is a useful, because, whenever this functions calls it just write the same text in log
	// so, this is function is just FOR FUN, no more no less
	_logfile_message := []byte(fmt.Sprintf("Error to Create a screen in function <-main->\nline 48\nInput: %s\n--- %s\n", _inputText, _format_date))

	// determinate a /home dir in user system
	// even though i really don't know how it will work on windows system, because I'm doing this for UNIX systems
	// so probably this part of code is not a multi-platform
	// I think that later when i finish the CLI interface and main code, I will add the switch for this functin
	// a.k.a if = 0 then exucate this function otherwise if = 1, ignore it
	_homeDir, _err := os.UserHomeDir()
	if _err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to dedicate the home directory, check your OS, if you on Windows. please set the variable flag as var flag bool = false %v\n", _err)
		return 
	}
	_logDir := filepath.Join(_homeDir, "_go", "homework_1")
	_write_message_to_logfile := filepath.Join(_logDir, "logfile.log")

	if _err := os.MkdirAll(_logDir, 0755); _err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to create  the directory, check your OS, if you on Windows. please set the variable flag as var flag bool = false : %v\n", _err)
		return
	}

	if _err := os.WriteFile(_write_message_to_logfile, _logfile_message, 0644); _err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to write a message in log file , check your OS, if you on Windows. Please, set the variable flag as var flag bool = false  %v\n", _err)
	}
}

// creating consts for calculationg
const (
	USD_RATE      = 38.5  // from UA to US exchange 
	EUR_RATE      = 42.1  // from UAH to EUR exchange
	INFLATION     = 6.0   // Infaltion
	MAX_INPUT_LEN = 15    // Max input, we are so rich to enter such a big numbers
)

// So this is basically this is  our Input, where 0 is initial 1 is rate 2 is years and 3 is contibutions
type _Inputstate struct {
	_ActiveField int    // 0: Initial, 1: Rate, 2: Years, 3: Contribution
	_Texts       [4]string // adding a froud fields for initial, rate, years, contribituion 
}


// I KNOW, THIS IS WRONG, AND YEARS SHOULDN'T BE WITH A FLOAT64 
// BUT CALCULATING YEARS IN INT TYPE WITH OTHERS FLOAT64 TYPES VARIABLE
// NOT ONLY WILL IT CAUSE AN ERROR, BUT YOU'LL HAVE TO USE SOME STATIC CAST ANALOG

type _Results struct {
	_SimpleInterest         float64
	_FinalAmountSimple      float64
	_FinalAmountCompound    float64
	_FinalAmountWithContrib float64
	_TotalProfit            float64
	_RealRate               float64
	_GrowthRatio            float64
	_AnnualReturn           float64
	_CompoundVsSimpleDiff   float64
	_InflationImpact        float64
}

func main() {
	// set xterm-250color for normal work
	if os.Getenv("TERM") == "" {
		os.Setenv("TERM", "xterm-256color")
	}

	// initializing a screen
	// only god know why doing this is waaaay harder then in c++ ncurses library
	_screen, _err := tcell.NewScreen()
	if _err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to Create the screen: %v\n", _err)
		// I ain't don't know why you will this error, I think it will appear only if you don't install a  gdamore/tcell/v2
		// library or if you using xterm ;)
		write_log("Error to Create the Screen \n Please, check your terminal compatibility \n Or your terminal resolution")
		os.Exit(1)
	}

	if _err := _screen.Init(); _err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to Initalize the screen: %v\n", _err)
		// the same thing, only if you do not install gdamore/tcell/v2 library
		write_log("Error to Initalize the Screen \n Please, check your terminal compatibility \n Or your terminal resolution")
		os.Exit(1)
	}

	// set it because of when you decide to shoot the programm off, your terminal will transform in garbage
	// I believe that I have the example of this when I was working with ncurses but I haven't
	// but basically if my mind do not betraying me, such command like "ls" will output directories and files in chaotic way
	// and the only working method is just restart your terminal
	defer _screen.Fini()

	// Define styles
	_defaultStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	_inputStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorBlack)
	_Title := tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack)

	// Initialize input state
	_state := _Inputstate{_ActiveField: 0}
	_showResults := false
	var _results _Results

	// Main loop for drawing a box and subbox
	for {
		// clearing screen to remove all garbage 
		_screen.Clear()

		// drawing a Title
		drawText(_screen, 5, 1, "Calculator which make you rich or broke :) ", _Title)

		// Our primary labels for inputing 
		_labels := []string{
			"Initial Investment (UAH):",
			"Annual Rate (%):",
			"Years:",
			"Monthly Contribution (UAH):",
		}

		for _i, _label := range _labels {
		
			_y := 3 + _i*2

			//drawing a labels
			drawText(_screen, 5, _y, _label, _defaultStyle)
			// okay this call drawBox call looking weird, the first "box" is a little bit "uncorrect"
			// when iterating in _y and _y sets a number 5, the little line just go down
			// i don't know how to fix it
			// 
			drawBox(_screen, 30, _y, 50, _y+1, _inputStyle)
			_text := _state._Texts[_i]
			if _i == _state._ActiveField {
				_text += "_"
			}
			drawText(_screen, 32, _y+1, _text, _inputStyle)
		}

		// Draw results if calculated
		if _showResults {
			drawBox(_screen, 5, 11, 60, 30, _defaultStyle)
			drawText(_screen, 7, 12, "Results:", _Title)

			// Convert to USD and EUR
			_usdSimple := _results._FinalAmountSimple / USD_RATE
			_eurSimple := _results._FinalAmountSimple / EUR_RATE
			_usdCompound := _results._FinalAmountCompound / USD_RATE
			_eurCompound := _results._FinalAmountCompound / EUR_RATE
			_usdContrib := _results._FinalAmountWithContrib / USD_RATE
			_eurContrib := _results._FinalAmountWithContrib / EUR_RATE

			// Okay so this is part of code starting from this line and ending with line
			// Which is named by "THIS LINE"
			// That all generated grok
			// because of I really tired to writing all this line, and keep getting an error because of
			// LITTLE MISMATH I DID
			_resultLines := []string{
				fmt.Sprintf("Simple Interest (UAH): %.2f", _results._SimpleInterest),
				fmt.Sprintf("Final Amount Simple (UAH): %.2f", _results._FinalAmountSimple),
				fmt.Sprintf("  USD: %.2f", _usdSimple),
				fmt.Sprintf("  EUR: %.2f", _eurSimple),
				fmt.Sprintf("Final Amount Compound (UAH): %.2f", _results._FinalAmountCompound),
				fmt.Sprintf("  USD: %.2f", _usdCompound),
				fmt.Sprintf("  EUR: %.2f", _eurCompound),
				fmt.Sprintf("Final Amount with Contributions (UAH): %.2f", _results._FinalAmountWithContrib),
				fmt.Sprintf("  USD: %.2f", _usdContrib),
				fmt.Sprintf("  EUR: %.2f", _eurContrib),
				fmt.Sprintf("Total Profit (UAH): %.2f", _results._TotalProfit),
				fmt.Sprintf("Real Annual Rate (%%): %.1f", _results._RealRate),
				fmt.Sprintf("Growth Ratio: %.2f", _results._GrowthRatio),
				fmt.Sprintf("Average Annual Return (%%): %.1f", _results._AnnualReturn),
				fmt.Sprintf("Compound vs Simple Difference (UAH): %.2f", _results._CompoundVsSimpleDiff),
				fmt.Sprintf("Inflation Impact (UAH): %.2f", _results._InflationImpact),
			}
// THIS LINE 
			//drawing a result box
			for _i, _line := range _resultLines {
				drawText(_screen, 7, 14+_i, _line, _defaultStyle)
			}
		}

		// Writing an insturckinos 
		_instr := "Enter: Next field, Ctrl+C/Esc: Exit, Ctrl+R: Calculate"
		var _instrY int
		if _showResults {
			_instrY = 31
		} else {
			_instrY = 11
		}

		// I need to cooment it  ?
		drawText(_screen, 5, _instrY, _instr, _defaultStyle)


		// updating the screen
		_screen.Show()

		// Hooking an key hitings on keyboard


		// so this is basically A WONDERFUL function 
		// because of it, I don't need to write timers or Infinity loop as I used to Did it in 
		// ncurses project
		_ev := _screen.PollEvent() 
		switch _ev := _ev.(type) {
		case *tcell.EventKey:
			switch _ev.Key() {
			case tcell.KeyCtrlC, tcell.KeyEscape:
				return
			case tcell.KeyEnter:
				// Move to next input field
				_state._ActiveField = (_state._ActiveField + 1) % 4
			case tcell.KeyCtrlR:
				// Calculate results
				if validInputs(_state) {
					_results = calculateResults(_state)
					_showResults = true
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:

				// THIS IS ALSO GROK GENERATED 
				if len(_state._Texts[_state._ActiveField]) > 0 {
					_, _size := utf8.DecodeLastRuneInString(_state._Texts[_state._ActiveField])
					_state._Texts[_state._ActiveField] = _state._Texts[_state._ActiveField][:len(_state._Texts[_state._ActiveField])-_size]
				}

			
			case tcell.KeyRune:
				// Add character if valid (digits, dot)
				if len(_state._Texts[_state._ActiveField]) < MAX_INPUT_LEN {
					_r := _ev.Rune()
					if (_r >= '0' && _r <= '9') || _r == '.' {
						_state._Texts[_state._ActiveField] += string(_r)

				// GENERATED TO THIS MOMENT
					}
				}
			}
		}
	}
}


// checking if all inputs that we did is normal
func validInputs(_state _Inputstate) bool {
	for _, _text := range _state._Texts {
		if _, _err := strconv.ParseFloat(_text, 64); _err != nil {
			return false
		}
	}
	return true
}

// calculateResults computes all investment metrics
func calculateResults(_state _Inputstate) _Results {
	var _results _Results

	// Parse inputs
	_initialInvestment, _ := strconv.ParseFloat(_state._Texts[0], 64)
	_annualRate, _ := strconv.ParseFloat(_state._Texts[1], 64)
	_years, _ := strconv.ParseFloat(_state._Texts[2], 64)
	_monthlyContribution, _ := strconv.ParseFloat(_state._Texts[3], 64)

	// simpleInterest = initialInvestment * (annualRate/100) * years
	_results._SimpleInterest = _initialInvestment * (_annualRate / 100) * _years

	// finalAmountSimple = initialInvestment + simpleInterest
	_results._FinalAmountSimple = _initialInvestment + _results._SimpleInterest

	// finalAmountCompound = initialInvestment * (1 + annualRate/100)^years
	_results._FinalAmountCompound = _initialInvestment * math.Pow(1+_annualRate/100, _years)

	// Final Amount with Monthly Contributions
	_monthlyRate := (_annualRate / 100) / 12
	_totalMonths := _years * 12

	// Contribution Effect: monthlyContribution * ((1 + monthlyRate)^totalMonths - 1) / monthlyRate 
	var _contributionEffect float64
	if _monthlyRate > 0 {
		_contributionEffect = _monthlyContribution * (math.Pow(1+_monthlyRate, _totalMonths)-1) / _monthlyRate
	}
	// finalAmountWithContrib = finalAmountCompound + contributionEffect
	_results._FinalAmountWithContrib = _results._FinalAmountCompound + _contributionEffect

	// Total Profit a.k.a finalAmountWithContrib - initialInvestment - (monthlyContribution * totalMonths)
	_results._TotalProfit = _results._FinalAmountWithContrib - _initialInvestment - (_monthlyContribution * _totalMonths)

	// realRate = ((1 + annualRate/100) / (1 + inflationRate/100) - 1) * 100
	_results._RealRate = ((1 + _annualRate/100) / (1 + INFLATION/100) - 1) * 100

	// Growth Ratio: finalAmountWithContrib / initialInvestment
	_results._GrowthRatio = _results._FinalAmountWithContrib / _initialInvestment


	// if average anarual is higher then 0 then execute, if not then skip it
	if _years > 0 {
		_results._AnnualReturn = (math.Pow(_results._GrowthRatio, 1/_years) - 1) * 100
	}

	// diffrence between SimpleInterest and Compoud
	_results._CompoundVsSimpleDiff = _results._FinalAmountCompound - _results._FinalAmountSimple


	// Inflation impact, god even here...
	if _years > 0 {
		_realValue := _results._FinalAmountWithContrib / math.Pow(1+INFLATION/100, _years)
		_results._InflationImpact = _results._FinalAmountWithContrib - _realValue
	}


	// All information just in ONE return..... 
	return _results
}
