<script lang="ts">
	import { onMount } from 'svelte';

	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	let wasmInstance: any;
	let total: number = $state(0);

	onMount(async () => {
		try {
			// eslint-disable-next-line no-undef
			const go = new Go();
			const result = await WebAssembly.instantiateStreaming(
				fetch('/wasm/main.wasm'),
				go.importObject
			);
			wasmInstance = result.instance;
			go.run(wasmInstance);
		} catch (err) {
			console.error('Failed to load WASM:', err);
		}
	});

	async function handleFileUpload(event: Event) {
		const input = event.target as HTMLInputElement;
		const file = input.files?.[0];

		if (!file) return;

		const text = await file.text();
		if (typeof window.processCsv === 'function') {
			try {
				total = window.processCsv(text);
			} catch (err) {
				console.error('Error processing CSV:', err);
			}
		}
	}
</script>

<svelte:head>
	<script src="/wasm/wasm_exec.js"></script>
</svelte:head>

<main class="container">
	<h1>CSV Price Calculator</h1>

	<div class="upload-section">
		<label for="csvFile">Upload CSV file:</label>
		<input type="file" id="csvFile" accept=".csv" onchange={handleFileUpload} />
	</div>

	<div class="result-section">
		<h2>Total: ${total.toFixed(2)}</h2>
	</div>
</main>

<style>
	.container {
		max-width: 800px;
		margin: 0 auto;
		padding: 2rem;
	}

	.upload-section {
		margin: 2rem 0;
	}

	.upload-section label {
		display: block;
		margin-bottom: 0.5rem;
	}

	.result-section {
		margin-top: 2rem;
	}

	h1 {
		color: #333;
		text-align: center;
	}

	h2 {
		color: #666;
	}
</style>
