import PageLayout from "../components/PageLayout/PageLayout";
import {
	Landing,
	Login,
	Home,
	Events,
	Event,
	CreateEvent,
	UserImageUpload,
	Image,
	JoinEvent,
} from "../pages";

const routes: RouteType[] = [
	{
		title: "Landing",
		path: "/",
		description: "Landing Page of Fotos",
		element: <Landing />,
	},
	{
		title: "Login",
		path: "/login",
		description: "Login/SignUp Page of Fotos App",
		element: <Login />,
	},
	{
		title: "First Time Image Upload",
		path: "/first-login",
		description: "First Time Image Upload Page of Fotos App",
		element: <UserImageUpload />,
	},
	{
		title: "Home",
		path: "/home",
		description: "Home Page of Fotos App",
		element: <Home />,
	},
	{
		title: "Image",
		path: "/image/:imageId",
		description: "Image Page of Fotos App",
		element: <Image />,
	},
	{
		title: "Event",
		path: "event",
		description: "Event Page of Fotos App",
		element: <PageLayout title="Events" />,
		children: [
			{
				title: "Events",
				path: "events",
				description: "Events Page of Fotos App",
				element: <Events />,
			},
			{
				title: "Event",
				path: ":eventId",
				description: "Event Page of Fotos App",
				element: <Event />,
			},
			{
				title: "Create Event",
				path: "create",
				description: "Create Event Page of Fotos App",
				element: <CreateEvent />,
			},
			{
				title: "Join Event",
				path: "join/:joinLink",
				description: "Join Event Page of Fotos App",
				element: <JoinEvent />,
			},
		],
	},
];

export default routes;
