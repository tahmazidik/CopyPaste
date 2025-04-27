package editor_test

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/tahmazidik/Copy_Paste/internal/editor"
)

// TestEditorScenario1 - тестовый сценарий для редактора
// t *testing.T - объект для управления тестами
func TestEditorScenario1(t *testing.T) {
	// Создаем временную директорию
	//Она автоматически удалится после завершения теста
	//Используется для изоляции тестов
	tmpDir := t.TempDir()

	// Подготовка тестовых файлов
	inputContent := []byte("My\nprogram\nis\nawful\nbad\npoor\nwrong\nawesome\n")
	cmdContent := []byte("Down\nDown\nDown\nShift\nDown\nDown\nDown\nDown\nCtrl+X")
	expectedContent := []byte("My\nprogram\nis\nawesome\n")

	//Создание файлов во врменной директории
	// Записываем тестовые данные во временные файлы
	inputPath := filepath.Join(tmpDir, "input.txt") //filepath.Join - собирает путь к файлу
	cmdPath := filepath.Join(tmpDir, "commands.txt")
	outputPath := filepath.Join(tmpDir, "output.txt")

	// 0644 - права доступа к файлу
	if err := os.WriteFile(inputPath, inputContent, 0644); err != nil { // os.WriteFile - записывает данные в файл
		t.Fatalf("Failed to create input file: %v", err)
	}
	if err := os.WriteFile(cmdPath, cmdContent, 0644); err != nil {
		t.Fatalf("Failed to create commands file: %v", err)
	}

	// Создаем новый экземпляр редактора и запускаем редактор
	e := editor.NewEditor()
	if err := e.LoadFile(inputPath); err != nil { //Загражает файл в редактор
		t.Fatalf("LoadFile failed: %v", err)
	}

	// Обрабатываем команды
	processTestCommands(cmdPath, e, t)

	// Сохраняем результат
	if err := e.SaveFile(outputPath); err != nil {
		t.Fatalf("SaveFile failed: %v", err)
	}

	// Сравниваем с ожидаемым результатом
	compareFiles(t, outputPath, expectedContent)
}

func processTestCommands(path string, e *editor.Editor, t *testing.T) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Error opening commands file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		e.ProcessCommand(scanner.Text())
	}
}

func compareFiles(t *testing.T, actualPath string, expected []byte) {
	actualContent, err := os.ReadFile(actualPath)
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	// Убираем последнюю пустую строку для сравнения
	expectedStr := string(bytes.TrimSpace(expected))
	actualStr := string(bytes.TrimSpace(actualContent))

	if actualStr != expectedStr {
		t.Errorf("\nExpected:\n%s\n\nActual:\n%s", expectedStr, actualStr)
	}
}
