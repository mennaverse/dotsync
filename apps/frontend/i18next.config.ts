export default {
  locales: [
    "en",
    "pt-br"
  ],
  extract: {
    input: "src/**/*.{js,jsx,ts,tsx}",
    output: "src/locales/{{language}}.json"
  }
}