const path = require('path');
const fs = require('fs');
const { spawn } = require('child_process');

async function generateSchemas() {
  const inputDir = path.resolve(__dirname, '../../messages');
  const outputDir = path.resolve(__dirname, '../../messages/schemas');

  if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir, { recursive: true });
  }

  const files = fs.readdirSync(inputDir).filter((file: string) => file.endsWith('.ts'));

  for (const file of files) {
    const inputFile = path.join(inputDir, file);
    const outputFile = path.join(outputDir, file.replace('.ts', '.schemas.ts'));

    const relativeInputFile = path.relative(process.cwd(), inputFile);
    const relativeOutputFile = path.relative(process.cwd(), outputFile);

    const args: string[] = ['ts-to-zod', relativeInputFile, relativeOutputFile];

    console.log(`Running command: npx ${args.join(' ')}`);

    const childProcess = spawn('npx', args, { stdio: 'inherit', shell: true });

    childProcess.on('close', (code: number) => {
      if (code === 0) {
        console.log(`Schemas generated for ${file}`);
      } else {
        console.error(`Error generating schemas for ${file}, code: ${code}`);
      }
    });

    childProcess.on('error', (err: { message: any; }) => {
      console.error(`Error starting process for ${file}: ${err.message}`);
    });
  }
}

generateSchemas();
export {}