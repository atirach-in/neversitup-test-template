function Permutations(val) {
    if (val.length === 1) {
        return [val];
    };
  
    let permutations = new Set();
  
    for (let i = 0; i < val.length; i++) {
      let str = val[i];
      let slice = val.slice(0, i) + val.slice(i + 1);
      let again = Permutations(slice);
  
      for (let m of again) {
        permutations.add(str + m);
      }
    }

    return Array.from(permutations);
  }

let result = Permutations("abc");
console.log("Result Permutations: ", result);

function FindOdd(arr) {
    try{
        if(arr.length == 0) return 0;
        
        let key = {};

        for (let num of arr) {
        key[num] = (key[num] || 0) + 1;
        console.log("num: ", num);
        console.log("key: ", key);
        }
    
        for (let num in key) {
            if (key[num] % 2 !== 0) {
                return num;
            }
        }
    }catch(e){
        console.log("Err FindOdd: ", e);
        return "Has err: " + e;
    }
    
}

let resultOdd = FindOdd([1,1,2]);
console.log("Result FindOdd: ", resultOdd);

function CountSmile(arr){
    try{
        let filter = [')', 'D'];
        let count = 0;
        arr.forEach(element => {
            for(let f of filter) {
                let find = element.indexOf(f);
                if (find != -1) count++;
            }
        });

        return count;
    }catch(e){
        console.log("Err CountSmile:", e);
        return "Has err: " + e;
    }
}

let resultCountSmile = CountSmile([':)))))', ';(', ';}', ':-DDDDDD']);
console.log("Result CountSmile: ", resultCountSmile);