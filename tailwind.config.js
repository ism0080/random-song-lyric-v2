module.exports = {
	content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
	theme: { extend: {}, },
	plugins: [
		require('@tailwindcss/forms')
	]
}
