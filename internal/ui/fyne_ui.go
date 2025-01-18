package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/AndrivA89/fyne-todo-demo/internal/constants"
	"github.com/AndrivA89/fyne-todo-demo/internal/domain"
)

type FyneUI struct {
	app       fyne.App
	viewModel *TaskViewModel
	window    fyne.Window
	todoList  *fyne.Container
	cfg       config
}

type config struct {
	appliedFilter string
}

func NewFyneUI(app fyne.App, viewModel *TaskViewModel) *FyneUI {
	todoContainer := container.NewVBox()

	ui := &FyneUI{
		app:       app,
		viewModel: viewModel,
		window:    app.NewWindow(constants.WindowTitle),
		todoList:  todoContainer,
		cfg: config{
			appliedFilter: constants.FilterAll,
		},
	}

	input := widget.NewEntry()
	input.SetPlaceHolder(constants.AddTaskPlaceholder)

	addButton := widget.NewButton(constants.AddButtonLabel, func() {
		viewModel.AddTask(input.Text)
		ui.refreshTasks()
		input.SetText("")
	})

	themeButton := widget.NewButton(constants.ToggleThemeButtonLabel, func() {
		ui.toggleTheme()
	})

	filterDropdown := widget.NewSelect([]string{
		constants.FilterAll,
		constants.FilterActive,
		constants.FilterCompleted,
	}, func(selected string) {
		ui.cfg.appliedFilter = selected
		ui.refreshTasks()
	})
	filterDropdown.SetSelected(ui.cfg.appliedFilter)

	topContent := container.NewVBox(todoContainer, input, addButton, filterDropdown)
	mainContainer := container.NewBorder(
		topContent,
		themeButton,
		nil,
		nil,
	)

	ui.window.SetContent(mainContainer)
	ui.window.Resize(fyne.NewSize(600, 400))

	ui.refreshTasks()

	return ui
}

func (ui *FyneUI) refreshTasks() {
	filteredTasks := ui.viewModel.FilterTasks(ui.cfg.appliedFilter)

	ui.todoList.Objects = nil
	for _, task := range filteredTasks {
		ui.todoList.Add(ui.createTaskRow(task))
	}
	ui.todoList.Refresh()
}

func (ui *FyneUI) createTaskRow(task domain.Task) fyne.CanvasObject {
	label := widget.NewLabel(task.Text)

	check := widget.NewCheck(constants.EmptyLabel, func(checked bool) {
		ui.viewModel.ToggleTask(task.ID, checked)
	})
	check.SetChecked(task.Completed)

	editButton := widget.NewButtonWithIcon(constants.EmptyLabel, theme.SettingsIcon(), func() {
		entry := widget.NewEntry()
		entry.SetText(task.Text)

		form := dialog.NewForm(
			constants.EditTaskLabel,
			constants.SaveButtonLabel,
			constants.CancelButtonLabel,
			[]*widget.FormItem{
				widget.NewFormItem(constants.EditNewTextLabel, entry),
			},
			func(confirmed bool) {
				if newText := strings.TrimSpace(entry.Text); newText != "" && confirmed {
					ui.viewModel.EditTask(task.ID, newText)
					ui.refreshTasks()
				}
			},
			ui.window,
		)
		form.Resize(fyne.NewSize(400, 200))
		form.Show()
	})

	deleteButton := widget.NewButtonWithIcon(constants.EmptyLabel, theme.DeleteIcon(), func() {
		ui.viewModel.DeleteTask(task.ID)
		ui.refreshTasks()
	})

	moveUpButton := widget.NewButtonWithIcon(constants.EmptyLabel, theme.MoveUpIcon(), func() {
		ui.viewModel.MoveTask(task.ID, constants.UpButton)
		ui.refreshTasks()
	})

	moveDownButton := widget.NewButtonWithIcon(constants.EmptyLabel, theme.MoveDownIcon(), func() {
		ui.viewModel.MoveTask(task.ID, constants.DownButton)
		ui.refreshTasks()
	})

	buttons := container.NewHBox(moveUpButton, moveDownButton, editButton, deleteButton)
	row := container.NewBorder(nil, nil, check, buttons, label)

	return row
}

func (ui *FyneUI) updateTaskRow(task domain.Task) {
	for i, obj := range ui.todoList.Objects {
		if row, ok := obj.(*fyne.Container); ok {
			if label, ok := row.Objects[0].(*widget.Label); ok && label.Text == task.Text {
				ui.todoList.Objects[i] = ui.createTaskRow(task)
				ui.todoList.Refresh()
				return
			}
		}
	}
}

func (ui *FyneUI) toggleTheme() {
	currentTheme := ui.app.Preferences().StringWithFallback(constants.ThemeKey, constants.DarkThemeKey)
	if currentTheme == constants.LightThemeKey {
		ui.app.Preferences().SetString(constants.ThemeKey, constants.DarkThemeKey)
		ui.app.Settings().SetTheme(theme.DarkTheme())
	} else {
		ui.app.Preferences().SetString(constants.ThemeKey, constants.LightThemeKey)
		ui.app.Settings().SetTheme(theme.LightTheme())
	}
}

func (ui *FyneUI) Run() {
	ui.window.ShowAndRun()
}
