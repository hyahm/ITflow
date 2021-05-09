const a = [1, 2, 4, 5, 6]

let b = a.filter( item => {
    console.log(item)
    console.log(item!=2)
    item != 2
})


const l = a.length
let c = []
for (let v of a) {
    if (v != 2) {
        c.push(v)
    }
}

console.log(a)
console.log(b)
console.log(c)