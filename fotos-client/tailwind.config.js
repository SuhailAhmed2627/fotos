const defaultTheme = require("tailwindcss/defaultTheme");

const colors = {
	primary: {
		50: "#ffb6ff",
		100: " #ff9eff",
		200: " #ff87ff",
		300: " #ee6fff",
		400: " #d557ff",
		500: " #bc3eff",
		600: " #a31eff",
		700: " #8a00f4",
		800: " #7000dc",
		900: " #5600c5",
		DEFAULT: "#9600ff",
	},
	secondary: {
		50: "#ffbfff",
		100: " #ffa8ff",
		200: " #e691ff",
		300: " #cc7bff",
		400: " #b164ff",
		500: " #974eff",
		600: " #7b36ff",
		700: " #5d1aff",
		800: " #3900f5",
		900: " #0000de",
		DEFAULT: "#4900ff",
	},
	white: {
		DEFAULT: "#ffffff",
	},
	black: {
		DEFAULT: "#000000",
	},
	transparent: {
		DEFAULT: "#00000000",
	},
};

const fontFamily = {
	title: ["Poppins", "Helvetica", "Arial", "sans-serif"],
	display: ["MOMCAKE", "Helvetica", "Arial", "sans-serif"],
	body: ["Inter", "Helvetica", "Arial", "sans-serif"],
};

const fontSize = {
	display: "6rem",
	h1: "3rem",
	h2: "2.5rem",
	h3: "2rem",
	h4: "1.5rem",
	h5: "1.25rem",
	h6: "1rem",
	small: "0.75rem",
	base: "1rem",
	large: "1.5rem",
};

module.exports = {
	content: ["./src/**/*.{js,jsx,ts,tsx}"],
	important: "#root",
	theme: {
		extend: {
			fontFamily: {
				sans: ["Inter", ...defaultTheme.fontFamily.sans],
			},
			colors,
			transitionProperty: {
				height: "height",
			},
		},
		fontFamily,
		fontSize,
	},
	plugins: [],
};
