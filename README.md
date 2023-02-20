# todo-cli

<h3>Git</h3>
<p align="left">
<a href="" 
target="_blank"><img src="https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white" alt="git" /></a>&nbsp;
</p>


for continue on existing repo:
````
cd existing_repo
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
