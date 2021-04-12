```js
import gulp from 'gulp'
import fs from 'fs'
import git from 'gulp-git'
import del from 'del'
const COUNT_PER = 2
let files = []
function copy() {
  if (files.length > 0) {
    const src = files.slice(0, COUNT_PER).map(e => './remain/' + e)
    return gulp.src(src).pipe(gulp.dest('./content/'));
  }
  return Promise.resolve(0)
}
function remove() {
  const src = files.slice(0, COUNT_PER).map(e => './remain/' + e)
  return del(src)
}
function gitAdd() {
  return gulp.src('./content/*').pipe(git.add());
}
function gitCommit() {
  return gulp.src('./').pipe(git.commit(undefined, {
    args: '-m "fix"', disableMessageRequirement: true
  }));
}
function gitPush(cb) {
  return function push() {
    return new Promise((resolve, reject) => {
      git.push('origin', 'master', {args: " -f"}, function (err) {
        if (err){
          reject(err)
          gitPush(cb)()
        } else {
          resolve(0)
          cb && cb()
        }
      })
    })
  }
}

export function oneProcess(cb = null) {
  files = fs.readdirSync('./remain/').filter(e => e.endsWith('rar'))
  if (files.length > 0) {
    return gulp.series(copy,remove, gitAdd, gitCommit, gitPush(cb))()
  } else {
    return Promise.resolve(0)
  }
}

export default async function batchProcess() {
  async function step() {
    await oneProcess(() => {
      if (files.length > 0) {
        step()
      }
    })
  }
  await step()
}

```
