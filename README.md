# todo-cli

<h3>Requirements</h3>

[Requirements for this project](docs/Requirements)

> ---
> - <h4>Under Developing:</h4>
> - [ ] Entities Category Properties: Title Color
> - [ ] Behaviors: Create a new Category List User Categories with progress status Edit a category
> - [ ] Task Title DueDate Category IsDone
> - [ ] Behaviors: Create a new Task List User Today Tasks List User Tasks By Date Change Task status (done/undone) Edit a task
> - [ ] User Properties: ID Email Password
> - [ ] Behaviors: Register a user Log in user
> - [ ] User Story (use-cases) User should be registered [] User should be able to log in to the application [] User can create a new category [] 
> - [ ] User can add a new task [] User can see the list of categories with progress status User can see the List of its tasks [] User can see the 
> - [ ] Today’s Tasks User can see the Tasks by date User can Done/Undone a task User can Edit a task User can Edit a category



<h3>Git</h3>
<p align="left">
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white" alt="git" /></a>&nbsp;
</p>


for continue on existing repo:
````
cd existing_repo
git init
git remote add origin `git address`
git branch -M main
git push -uf origin main

````


for production, there is stable **main** branch:

```` 
 git checkout main
 
````

for develop:

```` 
 git checkout -b developing
 
````
after all, back to main:
```` 
 git merge --no-ff developing
 
````
