import imageCompression from 'browser-image-compression';

export class CompressImage  {
    public static async compress(image: File, sizeMB: number): Promise<File>{
        const options = {
            maxSizeMB: sizeMB,
        };

        try {
            const compressedFile = await imageCompression(image, options);
            return compressedFile;
        } catch (error) {
            throw error;
        }
    }

    public static async compressFromInput(input: HTMLInputElement, sizeMB: number): Promise<File[]>{
        let compressedFiles: File[] = [];
        for (const file of Array.from(input.files)) {
            try {
                const compressedFile = await this.compress(file, sizeMB);
                compressedFiles.push(compressedFile);
            } catch (error) {
                throw error;
            }
        }
        return compressedFiles;
    }
}