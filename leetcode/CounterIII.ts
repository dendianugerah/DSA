type ReturnObj = {
  increment: () => number;
  decrement: () => number;
  reset: () => number;
};

function createCounter(init: number): ReturnObj {
  let count = init;
  return {
    increment: () => ++count,
    decrement: () => --count,
    reset: () => (count = init),
  };
}

const counter = createCounter(10);
