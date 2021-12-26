import { promises as fs } from "fs";
import { join } from "path";

async function loadInput(inputPath: string): Promise<number[]> {
  const rawInput = await fs.readFile(inputPath, "utf8");
  const input = rawInput.split("\n").map((item) => Number(item));

  return input;
}

function measureDepthIncreases(depths: number[]): number {
  let depthIncreases = 0;

  for (let i = 1; i < depths.length; i++) {
    if (depths[i] > depths[i - 1]) {
      depthIncreases += 1;
    }
  }

  return depthIncreases;
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

  const depthIncreases = measureDepthIncreases(depthsData);

  console.log("Number of depth increases", depthIncreases);
}

main();
