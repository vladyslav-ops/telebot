package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v4"
)

// ── Рецепти ────────────────────────────────────────────────────────

type Recipe struct {
	Title       string
	Emoji       string
	Time        string
	Ingredients string
	Steps       string
}

var recipes = map[string]Recipe{
	"pizza": {
		Title: "Швидка піца на лаваші",
		Emoji: "🍕",
		Time:  "15 хв",
		Ingredients: `• Тонкий лаваш — 2 шт
• Томатний соус або кетчуп — 3 ст.л.
• Моцарела — 150 г
• Ковбаса або шинка — 100 г
• Оливки — жменька
• Орегано — за смаком`,
		Steps: `1. Розігрій духовку до 200°C
2. Поклади лаваш на деко
3. Змасти соусом, виклади начинку
4. Посип моцарелою та орегано
5. Запікай 10-12 хвилин до золотистої скоринки
6. Наріж та подавай гарячою! 🔥`,
	},
	"nuggets": {
		Title: "Курячі нагетси",
		Emoji: "🍗",
		Time:  "25 хв",
		Ingredients: `• Куряче філе — 500 г
• Яйця — 2 шт
• Панірувальні сухарі — 1 склянка
• Борошно — 0.5 склянки
• Сіль, перець, паприка — за смаком
• Олія для смаження`,
		Steps: `1. Наріж філе на шматочки 3×3 см
2. Підготуй 3 миски: борошно, збиті яйця, сухарі з паприкою
3. Обваляй кожен шматочок: борошно → яйце → сухарі
4. Смаж у розігрітій олії по 3-4 хв з кожного боку
5. Виклади на серветку, щоб зайвий жир стік
6. Подавай з улюбленим соусом! 🎉`,
	},
	"guacamole": {
		Title: "Гуакамоле з начос",
		Emoji: "🥑",
		Time:  "10 хв",
		Ingredients: `• Авокадо (стиглий) — 2 шт
• Лайм — 1 шт
• Помідор — 1 шт
• Цибуля червона — 0.5 шт
• Коріандр (кінза) — жменька
• Сіль, перець чилі — за смаком
• Чіпси начос — 1 пачка`,
		Steps: `1. Розріж авокадо, вийми кісточку, виклади м'якоть у миску
2. Розімни виделкою (не до пюре — залиш шматочки!)
3. Видави сік лайма
4. Дрібно наріж помідор, цибулю та кінзу
5. Змішай все, додай сіль та перець чилі
6. Подавай з начос та насолоджуйся! 🌮`,
	},
	"fries": {
		Title: "Картопля фрі в духовці",
		Emoji: "🍟",
		Time:  "30 хв",
		Ingredients: `• Картопля — 4-5 шт (середніх)
• Оливкова олія — 2 ст.л.
• Паприка — 1 ч.л.
• Часниковий порошок — 0.5 ч.л.
• Сіль — за смаком
• Пармезан (за бажанням) — 30 г`,
		Steps: `1. Розігрій духовку до 220°C
2. Наріж картоплю соломкою товщиною з мізинець
3. Замочи на 10 хв у холодній воді, обсуши
4. Змішай з олією та спеціями
5. Виклади в один шар на деко з пергаментом
6. Запікай 25 хв, перевертаючи через 15 хв
7. Посип пармезаном та подавай! 🧀`,
	},
	"wings": {
		Title: "Курячі крильця BBQ",
		Emoji: "🍗",
		Time:  "35 хв",
		Ingredients: `• Курячі крильця — 1 кг
• Соус BBQ — 4 ст.л.
• Мед — 2 ст.л.
• Соєвий соус — 2 ст.л.
• Часник — 3 зубчики
• Імбир тертий — 1 ч.л.
• Сіль, перець — за смаком`,
		Steps: `1. Розігрій духовку до 200°C
2. Крильця обсуши, посоли та поперчи
3. Змішай соус BBQ, мед, соєвий соус, тертий часник та імбир
4. Обваляй крильця в соусі
5. Виклади на деко з пергаментом
6. Запікай 30 хв, змащуючи соусом кожні 10 хв
7. Подавай гарячими з селерою! 🔥`,
	},
	"mix": {
		Title: "Мікс на компанію",
		Emoji: "🎉",
		Time:  "20 хв",
		Ingredients: `• Мікс горішків — 200 г
• Чіпси — 1 пачка
• Міні-ковбаски або кабаноси — 200 г
• Сир (нарізка) — 150 г
• Оливки та корнішони — по 100 г
• Хумус — 200 г
• Морквяні палички та селера — по жменьці
• Крекери — 1 пачка`,
		Steps: `1. Підігрій міні-ковбаски на сковорідці 5 хв
2. Наріж сир кубиками або скрути у трубочки
3. Хумус виклади у піалу по центру тарілки
4. Розклади все на великій дошці/тарілці секціями:
   — горішки, чіпси, крекери
   — ковбаски, сир
   — оливки, корнішони
   — овочі (морква, селера)
5. Фото для Інсти обов'язково! 📸
6. Насолоджуйся з друзями! 🥳`,
	},
}

// ── Меню (inline-клавіатура) ───────────────────────────────────────

func buildMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}

	btnPizza := menu.Data("🍕 Піца", "pizza")
	btnNuggets := menu.Data("🍗 Нагетси", "nuggets")
	btnGuacamole := menu.Data("🥑 Гуакамоле", "guacamole")
	btnFries := menu.Data("🍟 Картопля фрі", "fries")
	btnWings := menu.Data("🍗 Крильця BBQ", "wings")
	btnMix := menu.Data("🎉 Мікс на компанію", "mix")

	menu.Inline(
		menu.Row(btnPizza, btnNuggets),
		menu.Row(btnGuacamole, btnFries),
		menu.Row(btnWings, btnMix),
	)

	return menu
}

func formatRecipe(r Recipe) string {
	return fmt.Sprintf(
		`%s *%s*
⏱ Час приготування: %s

📋 *Інгредієнти:*
%s

👨‍🍳 *Приготування:*
%s`,
		r.Emoji, r.Title, r.Time, r.Ingredients, r.Steps,
	)
}

// ── Cobra-команда start ────────────────────────────────────────────

var telebotCmd = &cobra.Command{
	Use:   "start",
	Short: "Запустити SerialSnackBot",
	Long:  "Запускає Telegram-бота для пропозицій перекусів до серіалів.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("SerialSnackBot %s started\n", appVersion)

		token := os.Getenv("TELE_TOKEN")
		if token == "" {
			log.Fatal("TELE_TOKEN env variable is not set. Please set it to your Telegram bot token.")
		}

		bot, err := tele.NewBot(tele.Settings{
			Token:  token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Failed to create bot: %v", err)
		}

		menu := buildMenu()

		bot.Handle("/start", func(c tele.Context) error {
			welcomeMsg := `🍿 *Привіт! Я SerialSnackBot!*

Збираєшся дивитися серіал? Я допоможу обрати ідеальний перекус! 🎬

Обери що приготувати:`
			return c.Send(welcomeMsg, menu, tele.ModeMarkdown)
		})

		bot.Handle("/menu", func(c tele.Context) error {
			return c.Send("🍿 Обери перекус:", menu, tele.ModeMarkdown)
		})

		bot.Handle("/help", func(c tele.Context) error {
			helpMsg := `📖 *Як користуватися SerialSnackBot:*

/start — почати та побачити меню перекусів
/menu — показати меню знову
/help — ця довідка

Натисни кнопку з перекусом — і отримай швидкий рецепт! 🍕`
			return c.Send(helpMsg, tele.ModeMarkdown)
		})

		for key, recipe := range recipes {
			r := recipe // захоплюємо змінну для замикання
			bot.Handle(&tele.Btn{Unique: key}, func(c tele.Context) error {
				backMenu := &tele.ReplyMarkup{}
				btnBack := backMenu.Data("⬅️ Назад до меню", "back_to_menu")
				backMenu.Inline(backMenu.Row(btnBack))

				return c.Send(formatRecipe(r), backMenu, tele.ModeMarkdown)
			})
		}

		bot.Handle(&tele.Btn{Unique: "back_to_menu"}, func(c tele.Context) error {
			return c.Send("🍿 Обери перекус:", menu, tele.ModeMarkdown)
		})

		bot.Handle(tele.OnText, func(c tele.Context) error {
			msg := fmt.Sprintf(
				"🤔 Не зрозумів команду: *%s*\n\nНатисни /start щоб побачити меню перекусів! 🍿",
				c.Text(),
			)
			return c.Send(msg, tele.ModeMarkdown)
		})

		fmt.Println("Bot is running. Press Ctrl+C to stop.")
		bot.Start()
	},
}

func init() {
	rootCmd.AddCommand(telebotCmd)
}
