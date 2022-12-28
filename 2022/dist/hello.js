import fs from 'fs';
const input = fs.readFileSync(process.cwd() + "/src/day1/" + "puzzle1.txt", 'utf-8').toString();
// break the input into chunks
const calorieBlocks = input.split("\n\n");
console.log(calorieBlocks.length);
var max = 0;
for (let i = 0; i < calorieBlocks.length; i++) {
    console.log(input[i]);
    var snacks = calorieBlocks[i].split("\n");
    let sum = 0;
    snacks.forEach((snack) => {
        sum += parseInt(snack);
    });
    if (sum > max) {
        max = sum;
    }
}
console.log(max);
//# sourceMappingURL=hello.js.map