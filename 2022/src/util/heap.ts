

// class MaxHeap is a zero-based max heap
export class MaxHeap {
	public heap: Array<Number> = new Array<Number>();
	constructor () {}

	private swapIndex = (a,b) => {
		[this.heap[a], this.heap[b]] = [this.heap[b], this.heap[a]];
	}
	public getMax (): Number {
		return this.heap[0];
	}

	public insert (value: Number): void {
		this.heap.push(value);

		if (this.heap.length > 1) {
			let currentIdx = this.heap.length - 1;

			while (currentIdx > 1 && this.heap[Math.floor(currentIdx / 2)] > this.heap[currentIdx]) {
				this.swapIndex(currentIdx, Math.floor(currentIdx / 2));
				currentIdx = Math.floor(currentIdx / 2);
			}
		}
	}

	public remove (): Number {
		this.heap[0] = this.heap[this.heap.length - 1];
		this.heap.splice(this.heap.length - 1, 1);


		let parent = 0;
		let leftChildIdx = 2 *parent+1;
		let rightChildIdx = 2 *parent+2;

		let left = leftChildIdx > this.heap.length  ? null: this.heap[leftChildIdx];
		let right = rightChildIdx > this.heap.length ? null: this.heap[rightChildIdx];

		do {

			if (!left && !right) {
				return;
			}

			// 1 child
			else if(!right) {
				if (this.heap[parent] > left) {
					this.swapIndex(parent, leftChildIdx);
					parent = leftChildIdx;
				}
				else {
					return;
				}
			}
			// 2 children
			else if (left && right) {
				if (left < right && this.heap[parent] > left) {
					this.swapIndex(parent, leftChildIdx);
					parent = leftChildIdx;
				}
				else if (left > right && this.heap[parent] > right) {
					this.swapIndex(parent, rightChildIdx);
					parent = rightChildIdx;
				}
			}

			console.log(JSON.stringify(this.heap));

			leftChildIdx = 2 *parent+1;
			rightChildIdx = 2 *parent+2;
			left = leftChildIdx > this.heap.length ? null: this.heap[leftChildIdx];
			right = rightChildIdx > this.heap.length ? null: this.heap[rightChildIdx];

		} while (left && right && this.heap[parent] > left && this.heap[parent] > right)
	}

}
