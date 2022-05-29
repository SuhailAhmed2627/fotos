import {
	Accordion,
	Avatar,
	Box,
	Button,
	Center,
	Navbar,
	Stack,
	Text,
} from "@mantine/core";
import { getUser } from "../../utils/helperFuntions";
import { MdEvent, MdGroup, MdImage, MdLogout } from "react-icons/md";
import SidebarLink from "../SidebarLink/SidebarLink";
import { useDispatch } from "react-redux";
import { loginSuccess } from "../../actions/user";
import { AnyAction, Dispatch } from "@reduxjs/toolkit";

const logout = (dispatch: Dispatch<AnyAction>) => {
	localStorage.clear();
	dispatch(loginSuccess(null));
	window.location.href = "/";
};

const SidebarElement = [
	{
		title: "Events",
		icon: <MdEvent />,
		link: "/event/events",
	},
];

const Sidebar = (props: SidebarProps) => {
	const user = getUser();
	const dispatch = useDispatch();
	if (!user) {
		return null;
	}
	if (user.firstLogin) {
		return null;
	}
	return (
		<Navbar className="text-gray-600 w-[70vw] md:w-[20vw]">
			<Navbar.Section className="p-3">
				<Text className="font-title text-h3 font-semibold">Dashboard</Text>
			</Navbar.Section>
			<Navbar.Section grow className="border-t-2 p-3">
				<Stack spacing={"xs"} className="w-full" align={"center"}>
					<SidebarLink
						icon={<MdImage />}
						title="Your Photos"
						link="/home"
					/>
					<Accordion className="w-full border-b-0">
						<Accordion.Item
							className="border-b-0"
							label="Recent Events"
							classNames={{
								control:
									"text-gray-600 bg-gray-100 hover:bg-gray-200 rounded-lg",
								content: "text-gray-600",
								contentInner: "pr-0",
							}}
						>
							<Stack spacing={"xs"}>
								{user &&
									user.events &&
									user.events.map((event) => (
										<SidebarLink
											key={event.eventId}
											link={`event/${event.eventId}`}
											icon={<MdEvent />}
											title={event.name}
										/>
									))}
							</Stack>
						</Accordion.Item>
					</Accordion>
					{SidebarElement.map((element, index) => (
						<SidebarLink
							key={index}
							link={element.link}
							icon={element.icon}
							title={element.title}
						/>
					))}{" "}
				</Stack>
			</Navbar.Section>
			<Navbar.Section className="border-t-2">
				<Box className="p-2 w-full justify-between flex flex-row">
					<Avatar size={50} radius={25} />
					<Stack justify={"center"} className="gap-0">
						<Text className="font-semibold" lineClamp={1}>
							{user ? user.name : "Not Logged In"}
						</Text>
						<Text lineClamp={1} className="text-small">
							{user ? user.username : "Not Logged In"}
						</Text>
					</Stack>
					<Center
						onClick={() => logout(dispatch)}
						className="justify-self-end cursor-pointer"
					>
						<MdLogout size={40} />
					</Center>
				</Box>
			</Navbar.Section>
		</Navbar>
	);
};

export default Sidebar;
