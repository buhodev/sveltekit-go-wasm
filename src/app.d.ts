// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}

	interface Window {
		processCsv: (csvContent: string) => number;
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
		Go: any;
	}
}

export {};
