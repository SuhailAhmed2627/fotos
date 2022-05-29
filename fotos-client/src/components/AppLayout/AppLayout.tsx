import { AppShell, Navbar, Header } from "@mantine/core";
import { getUser } from "../../utils/helperFuntions";
import NavBar from "../NavBar/NavBar";
import Sidebar from "../Sidebar/Sidebar";

import { AppLayoutProps } from "./types";

const AppLayout = (props: AppLayoutProps) => {
	return (
		<AppShell
			padding={0}
			navbar={<Sidebar opened={true} />}
			header={<NavBar />}
			classNames={{ body: "h-[100vh]" }}
		>
			{props.children}
		</AppShell>
	);
};

export default AppLayout;
