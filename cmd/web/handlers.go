package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Обработчик главной странице.
// Меняем сигнатуры обработчика home, чтобы он определялся как метод
// структуры *application.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		app.notFound(w) // Использование помощника notFound()
		return
	}

	// Инициализируем срез содержащий пути к двум файлам.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Поскольку обработчик home теперь является методом структуры application
		// он может получить доступ к логгерам из структуры.
		// Используем их вместо стандартного логгера от Go.
		//app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.

	err = ts.Execute(w, nil)
	if err != nil {

		// Обновляем код для использования логгера-ошибок
		// из структуры application.
		//app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err) // Использование помощника serverError()
	}

}

// Меняем сигнатуру обработчика showSnippet, чтобы он был определен как метод
// структуры *application
// Обработчик для отображения содержимого заметки.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Отображение заметки..."))
	id, err := strconv.Atoi(r.URL.Query().Get("id")) //Извлекаем значение параметра id из URL и  конв в integer
	if err != nil || id < 1 {
		//http.NotFound(w, r)
		app.notFound(w) // Использование помощника notFound()
		return
	}
	fmt.Fprintf(w, "Отображение выбранной заметки с %d...", id)
}

// Обработчик для создания новой заметки.
// Меняем сигнатуру обработчика createSnippet, чтобы он определялся как метод
// структуры *application.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
	// что http.MethodPost является строкой и содержит текст "POST".
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		// Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением.
		//http.Error(w, "Метод запрещен!", 405)

		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
		app.clientError(w, http.StatusMethodNotAllowed) // Используем помощник clientError()
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}
