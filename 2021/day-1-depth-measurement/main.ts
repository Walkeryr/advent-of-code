import { promises as fs } from "fs";
import { join } from "path";
import { measureSimpleIncreases, measureSlidingWindowIncreases } from "./depthMeasurement";

async function loadInput(inputPath: string): Promise<number[]> {
  const rawInput = await fs.readFile(inputPath, "utf8");
  const input = rawInput.split("\n").map((item) => Number(item));

  return input;
}

async function main() {
  const INPUT_PATH = join(__dirname, "input.txt");
  let depthsData: number[];

  try {
    depthsData = await loadInput(INPUT_PATH);
  } catch (error) {
    console.log("Error reading input file");
    process.exit(1);
  }

  const simpleIncreases = measureSimpleIncreases(depthsData);
  console.log("Number of depth increases", simpleIncreases);

  const slidingWindowIncreases = measureSlidingWindowIncreases(depthsData);
  console.log("Number of depth increases (sliding window)", slidingWindowIncreases);
}

main();
