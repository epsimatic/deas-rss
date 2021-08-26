# DEAS RSS Reader

Is a simple online reader for RSS feeds

### Operations:

- **Installation and running:**    
  Follow Buffalo manual below

- **Adding RSS feed:**   
  ```bash
  buffalo task rss:fetch {URL} [URL]...
  ```
  example:   
  ```bash
  buffalo task rss:fetch https://www.feedforall.com/sample.xml
  ```

- **Updating RSS feed:**
  ```bash
  buffalo task rss:fetch {URL} [URL]...
  ```
  
### TODOs:

- Add RSS feed UI
- Delete oldest feeds on update
- Mass update
- Periodic mass update
- Better article index UI
- Cleanup article actions
- Refactor to use actual `feeds` model and table
- Fix and write tests
- See TODO / FIXME in code

---

# Buffalo manual

## Database Setup

It looks like you chose to set up your application using a database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start it for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

	$ buffalo pop create -a

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a home page.

**Congratulations!** You now have your Buffalo application up and running.

[Powered by Buffalo](http://gobuffalo.io)
