import fs from 'fs';

const input = fs.readFileSync( process.cwd() + "/src/day2/"+ "puzzle2.txt", 'utf-8').toString();


// take the first 3 lines of the input
const regex = /(?:.*\n?){1,3}/g;
const matches = input.match(regex);
for (let i = 0; i < matches.length; i++) {

	console.log(matches[i]);
	let round = matches[i].split('\n').filter(x => x.trim())
	if (!validateMatch(round)) {
		continue;
	}
	round1Result = round1Score(round);
	// console.log(round1Result);
}


function round1Score(a:string[]): [number,number] {
	let score = [0,0];
	for(let i = 0; i < a.length; i++){
		if(a[i] == 'rock'){
			score[0] += 1;
		}else if(a[i] == 'paper'){
			score[1] += 1;
		}
	}
	return [0,0];

}

function validateMatch(round: string[]): Boolean{
	// round = ["B Z", "C D", "A B"]
	if (round.length != 3) {
		console.log("Invalid match: ", round);
		return false;
	}
	return true;
}