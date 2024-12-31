import { PrimeReactProvider } from "primereact/api";
import { Layout } from "./components/Layout";
import { RepoPathProvider } from "./context";
import "primereact/resources/themes/lara-light-cyan/theme.css";
import "uno.css";
import "./App.css";

const App = () => {
	return (
		<PrimeReactProvider>
			<RepoPathProvider>
				<Layout />
			</RepoPathProvider>
		</PrimeReactProvider>
	);
};

export default App;
