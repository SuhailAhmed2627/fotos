import { Stack, Tabs, Text } from "@mantine/core";
import React, { useState } from "react";
import {
	NavigateFunction,
	Outlet,
	useLocation,
	useNavigate,
} from "react-router-dom";

const eventPage = {
	title: "Events",
	description: "View and Create Events",
	links: [
		{
			tab: "Your Events",
			path: "/event/events",
		},
		{
			tab: "Create Event",
			path: "/event/create",
		},
	],
};

const handleTabChange = (
	tabIndex: number,
	setActiveTab: React.Dispatch<React.SetStateAction<number>>,
	navigate: NavigateFunction
) => {
	if (tabIndex == 0) {
		navigate("/event/events");
	} else {
		navigate("/event/create");
	}
	setActiveTab(tabIndex);
};

const PageLayout = (props: PageLayoutProps) => {
	const location = useLocation();
	const navigate = useNavigate();
	const [activeTab, setActiveTab] = useState(
		location.pathname == "/event/events" ? 0 : 1
	);
	const page = props.title == "events" ? eventPage : eventPage;

	if (
		location.pathname == page.links[0].path ||
		location.pathname == page.links[1].path
	) {
		return (
			<Stack className="p-10 h-full overflow-hidden">
				<Text className="font-title h-[10%] font-bold text-h1 md:text-[5rem] text-gray-700 leading-none">
					{page.title}
				</Text>
				<Text className="font-title text-h5 text-gray-600 leading-loose">
					{page.description}
				</Text>
				<Tabs
					className="flex-grow flex flex-col"
					classNames={{
						body: "flex-grow",
					}}
					active={activeTab}
					onTabChange={(tabIndex) =>
						handleTabChange(tabIndex, setActiveTab, navigate)
					}
					grow
					position="center"
				>
					{page.links.map((link) => (
						<Tabs.Tab key={link.path} label={link.tab}>
							{location.pathname == link.path && <Outlet />}
						</Tabs.Tab>
					))}
				</Tabs>
			</Stack>
		);
	}

	return <Outlet />;
};

export default PageLayout;
