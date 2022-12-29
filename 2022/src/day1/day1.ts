import fs from 'fs';
import {MaxHeap} from '../util/heap.js'

const input = fs.readFileSync( process.cwd() + "/src/day1/"+ "puzzle1.txt", 'utf-8').toString();


// break the input into chunks

const calorieBlocks = input.split("\n\n");

console.log(calorieBlocks.length);

var max = [0,0,0]
let sums = [];

for (let i = 0; i < calorieBlocks.length; i++) {
		var snacks = calorieBlocks[i].split("\n");

		console.log(snacks);
		let sum = 0;
		snacks.forEach((snack) => {
			sum += parseInt(snack);
		});

		console.log("the sum of the calories is " + sum);
		sums.push(sum)

    for(let j = 0; j < max.length; j++){
			if(sum > max[j]){
				console.log("the sum of ", sum, " is greater than ", max[j], " at index ", j)
				max[j] = sum;
				break;
			}
		}

}


console.log("the max calories are " + max);

// const initialValue = 0;
// const sumWithInitial = max.reduce(
// 	(accumulator,currentValue) => accumulator + currentValue, initialValue );
let sortedSums = sums.sort((a,b) => b-a);
let top3 = sortedSums.slice(0,3);
let sum = top3.reduce((a,b) => a+b, 0);
console.log("the total calories carried by the top 3 elves are " + sum);