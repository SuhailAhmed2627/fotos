import { createRoot } from "react-dom/client";
import { ErrorBoundary } from "react-error-boundary";
import { ErrorFallback } from "./components";
import App from "./App";

const root = createRoot(document.getElementById("root")!);
root.render(
	<ErrorBoundary FallbackComponent={ErrorFallback}>
		<App />
	</ErrorBoundary>
);
