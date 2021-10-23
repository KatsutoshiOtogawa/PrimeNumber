// 指定した自然数以下の素数を配列で返す。
function createPrimeNumberArray(naturalNumber: number):number[] | undefined {
    // 自然数かどうか判定
    if ( naturalNumber < 1 || !Number.isInteger(naturalNumber) ) {
        console.error("this number is not natural Number")
        return;
    }
    // 指定した数字を元に1...naturalNumberという配列を作成し、
    // 
    const resultArray = [...Array(naturalNumber + 1).keys()].filter((value: number, index: number, array: number[]) => {
        // 0, 1は当然素数でないので省く。
        if (value === 0 || value === 1) {
            return false;
        }
        // 最初に与えられた配列から、合成数かどうか判定する。
        const judge = array.some((value2: number) => {
            // 0, 1は素数判定に使えないので省く。
            // 素数かもしない数字より大きいのは割り切れないのですべてfalse
            if (value2 === 0 || value2 === 1 || value <= value2) {
                return false;
            }
            // modが0の場合は合成数なのでtrueが帰る。
            return value % value2 === 0
        })
        // 求めるべきは素数なので真偽値を反転。
        return !judge;
    })
    return resultArray;
}

// 指定した自然数以下の素数を配列で返す。
// 効率的な方法で求めた場合
function createPrimeNumberArray2(naturalNumber: number):number[] | undefined {
    // rootを使う効率的なやり方。
    // 自然数かどうか判定
    if ( naturalNumber < 1 || !Number.isInteger(naturalNumber) ) {
        console.error("this number is not natural Number")
        return;
    }
    // 指定した数字を元に1...naturalNumberという配列を作成し、
    // 
    const resultArray = [...Array(naturalNumber + 1).keys()].filter((value: number, index: number, array: number[]) => {
        // 0,1は当然素数でないので省く。
        if (value ===0 || value === 1) {
            return false;
        }
        // 効率的な方法で求めた場合
        const root = Math.sqrt(value);
        // 最初に与えられた配列から、合成数かどうか判定する。
        const judge = array.some((value2: number) => {
            // 0, 1は素数判定に使えないので省く。
            // 素数かもしない数字nの平方根までの数字で合成数、素数の判定ができるので
            if (value2 === 0 || value2 === 1 || root < value2) {
                return false;
            }
            // modが0の場合は合成数なのでtrueが帰る。
            return value % value2 === 0
        })
        // 求めるべきは素数なので真偽値を反転。
        return !judge;
    })
    return resultArray;
}

// 指定した自然数以下の素数を配列で返す。
// エラトステネスのふるいを使った場合。
function eratosthenePrimeNumberArray(naturalNumber: number):number[] | undefined {
    // 自然数かどうか判定
    if ( naturalNumber < 1 || !Number.isInteger(naturalNumber) ) {
        console.error("this number is not natural Number")
        return;
    }
    // インデックス0があるので、1ずらす。
    const len: number = naturalNumber + 1;
    // trueなら素数で、falseなら合成数または0,1。
    // 素数かもしれない値は取りあえすtrueとしておく。
    const judge: boolean = true;
    const list = new Array<boolean>(len).fill(judge);
    // 0, 1は素数でないので、falseを入れておく。
    list[0] = false;
    list[1] = false;
    // マークが終わったら、リストにある次の素数かもしれない数字に移る。
    for(let p=2; Math.pow(p,2) < list.length; p = list.findIndex((value: boolean, index: number) => value && p < index)) {        
        for(let q = Math.pow(p,2);q < list.length;q += p){
            // p の２乗+pは合成数。
            list[q] = false;
        }
    }
    const resultArray = list.reduce(function(accumulator: number[], currentValue: boolean, index: number) {
        if (currentValue === true) {
            accumulator.push(index);
        }
        return accumulator;
    }, [])
    return resultArray;
}

const num1 = 10
const array1 = createPrimeNumberArray(num1);
const array15 = createPrimeNumberArray2(num1);
const array17 = eratosthenePrimeNumberArray(num1);

console.log(`num1: ${num1} array1:` + array1);
console.log(`num1: ${num1} array15:` + array15);
console.log(`num2: ${num1} array17:` + array17);

const num2 = 100
const array2 = createPrimeNumberArray(num2);
const array25 = createPrimeNumberArray2(num2);
const array27 = eratosthenePrimeNumberArray(num2);

console.log(`num2: ${num2} array2: ` + array2);
console.log(`num2: ${num2} array25:` + array25);
console.log(`num2: ${num2} array27:` + array27);

// const num3 = 2147483647
// const array3 = eratosthenePrimeNumberArray(num3);

// console.log(`num3: ${num3} array3: ` + array3);