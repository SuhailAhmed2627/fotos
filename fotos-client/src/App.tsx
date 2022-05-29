import Router from "./routes";
import "/public/assets/index.css";
import { store, persistor } from "./store";
import { Provider } from "react-redux";
import { MantineProvider } from "@mantine/core";
import { NotificationsProvider } from "@mantine/notifications";
import { PersistGate } from "redux-persist/integration/react";

const App = () => {
	return (
		<Provider store={store}>
			<PersistGate loading={null} persistor={persistor}>
				<MantineProvider
					theme={{
						fontFamily: "Inter, sans-serif",
					}}
				>
					<NotificationsProvider>
						<Router />
					</NotificationsProvider>
				</MantineProvider>
			</PersistGate>
		</Provider>
	);
};

export default App;
