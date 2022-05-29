import routes from "./routes";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { NotFoundPage } from "../pages";
import { AppLayout } from "../components";

const Router = () => {
	return (
		<BrowserRouter>
			<AppLayout>
				<Routes>
					{routes.map((route) => {
						if (!route.children) {
							return (
								<Route
									key={route.path}
									path={route.path}
									element={route.element}
								/>
							);
						}
						return (
							<Route
								key={route.path}
								path={route.path}
								element={route.element}
							>
								{route.children.map((childRoute, index) => {
									return (
										<Route
											index={index === 0}
											key={childRoute.path}
											path={childRoute.path}
											element={childRoute.element}
										/>
									);
								})}
							</Route>
						);
					})}
					<Route path="*" element={<NotFoundPage />} />
				</Routes>
			</AppLayout>
		</BrowserRouter>
	);
};

export default Router;
