/** @type {import('tailwindcss').Config} */
const config = {
  content: [
      "./src/**/*.{js,ts,jsx,tsx}",
      "./node_modules/@material-tailwind/react/components/**/*.{js,ts,jsx,tsx}",
      "./node_modules/@material-tailwind/react/theme/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
      extend: {},
  },
  plugins: [],
};

export default config;
