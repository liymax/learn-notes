```js
const fs = require('fs')
const path = require('path')
const { parse } = require('@babel/parser')
const traverse = require('@babel/traverse').default
const template = require('@babel/template').default
const generate = require('@babel/generator').default
const t = require('@babel/types')
const pretty = require('ast-pretty-print')
const babelPlugins = [
  'asyncGenerators',
  'classProperties',
  ['decorators', { decoratorsBeforeExport: true }],
  'doExpressions',
  'dynamicImport',
  'exportDefaultFrom',
  'exportNamespaceFrom',
  'objectRestSpread',
  'optionalCatchBinding',
  'throwExpressions'
]
const node = t.variableDeclaration('const', [
  t.variableDeclarator(t.identifier('url'), t.stringLiteral('/p/k'))
])
const targetPath = path.resolve(__dirname, '../src', 'App.vue')
const originCode = fs.readFileSync(targetPath, { encoding: 'utf8' })
const mt = originCode.match(/<script.*>([\s\S]*)<\/script>/)
const ast = parse(mt[1], { sourceType: 'module', plugins: babelPlugins })
traverse(ast, {
  enter(path) {
    const { type, name } = path.node
    // console.log('type:', type, name)
    if (type === 'CallExpression') {
      const { arguments: args, callee } = path.node
      // logAst(arguments, callee)
      if (args[0].name === '_url') {
        path.insertAfter(node)
        path.remove()
      }
    }
    if (path.isIdentifier({ name: '_url' })) {
      path.replaceWith(t.identifier('url'))
    }
  }
})
// console.log(generator(ast).code)
const newCode = generate(ast, {
  comments: false,
  quotes: "double"
}, mt[1]).code
rewrite(newCode, originCode)

function logAst(ast) {
  console.log(pretty(ast))
}
function rewrite(newCode, originCode) {
  const newScript = `<script>\n${newCode}\n</script>`
  const newVue = originCode.replace(/<script.*>[\s\S]*<\/script>/, newScript)
  console.log(newVue)
  fs.writeFileSync(targetPath, newVue)
}
const files = fs.readdirSync(path.resolve(__dirname, '../src'))
console.log(files)
```
