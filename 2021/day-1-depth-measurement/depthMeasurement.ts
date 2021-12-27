export function measureSimpleIncreases(depths: number[]): number {
  let increases = 0;

  for (let i = 1; i < depths.length; i++) {
    if (depths[i] > depths[i - 1]) {
      increases += 1;
    }
  }

  return increases;
}

function getWindowSum(values: number[]) {
  return values.reduce((prev, curr) => prev + curr, 0);
}

export function measureSlidingWindowIncreases(
  depths: number[],
  windowSize = 3
): number {
  if (depths.length <= windowSize) {
    throw Error("Depths array couldn't be shorter than sliding window");
  }

  let increases = 0;

  for (let i = windowSize; i < depths.length; i++) {
    const prevWindowValue = getWindowSum(depths.slice(i - windowSize, i));
    const windowValue = getWindowSum(depths.slice(i - (windowSize - 1), i + 1));

    if (windowValue > prevWindowValue) {
      increases += 1;
    }
  }

  return increases;
}
