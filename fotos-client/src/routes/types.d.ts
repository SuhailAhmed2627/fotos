interface RouteType {
	title: string;
	path: `/${string}` | string;
	description: string;
	element: JSX.Element;
	children?: RouteType[];
}
