package test

// app := &cli.App{
// 	Name:  "tasker",
// 	Usage: "A simple CLI program to manage your tasks",
// 	Action: func(c *cli.Context) error {
// 		tasks, err := getPending()
// 		if err != nil {
// 			if err == mongo.ErrNoDocuments {
// 				fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
// 				return nil
// 			}

// 			return err
// 		}

// 		printTasks(tasks)
// 		return nil
// 	},
// 	Commands: []*cli.Command{
// 		{
// 			Name:    "add",
// 			Aliases: []string{"a"},
// 			Usage:   "add a task to the list",
// 			Action: func(c *cli.Context) error {
// 				str := c.Args().First()
// 				if str == "" {
// 					return errors.New("Cannot add an empty task")
// 				}

// 				task := &Task{
// 					ID:        primitive.NewObjectID(),
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					Text:      str,
// 					Completed: true,
// 				}

// 				return createTask(task)
// 			},
// 		},
// 		{
// 			Name:    "all",
// 			Aliases: []string{"l"},
// 			Usage:   "list all tasks",
// 			Action: func(c *cli.Context) error {
// 				tasks, err := getAll()
// 				if err != nil {
// 					if err == mongo.ErrNoDocuments {
// 						fmt.Print("Nothing to see here.\nRun `add 'task'` to add a task")
// 						return nil
// 					}

// 					return err
// 				}

// 				printTasks(tasks)
// 				return nil
// 			},
// 		},
// 		{
// 			Name:    "done",
// 			Aliases: []string{"d"},
// 			Usage:   "complete a task on the list",
// 			Action: func(c *cli.Context) error {
// 				text := c.Args().First()
// 				return completeTask(text)
// 			},
// 		},
// 		{
// 			Name:    "finished",
// 			Aliases: []string{"f"},
// 			Usage:   "list completed tasks",
// 			Action: func(c *cli.Context) error {
// 				tasks, err := getFinished()
// 				if err != nil {
// 					if err == mongo.ErrNoDocuments {
// 						fmt.Print("Nothing to see here.\nRun `done 'task'` to complete a task")
// 						return nil
// 					}

// 					return err
// 				}

// 				printTasks(tasks)
// 				return nil
// 			},
// 		},
// 		{
// 			Name:  "rm",
// 			Usage: "deletes a task on the list",
// 			Action: func(c *cli.Context) error {
// 				text := c.Args().First()
// 				err := deleteTask(text)
// 				if err != nil {
// 					return err
// 				}

// 				return nil
// 			},
// 		},
// 	},
// }

// err := app.Run(os.Args)
// if err != nil {
// 	log.Fatal(err)
// }
