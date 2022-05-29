import { Button } from "@mantine/core";
import { useNavigate } from "react-router-dom";
import { SidebarLinkProps } from "./types";

const SidebarLink = (props: SidebarLinkProps) => {
	const navigate = useNavigate();
	return (
		<Button
			onClick={() => navigate(props.link)}
			leftIcon={props.icon}
			classNames={{
				inner: "justify-start",
				label: "font-medium group-hover:font-semibold",
			}}
			className="w-full font-body group text-gray-600 bg-gray-100 hover:bg-gray-200 hover:font-semibold"
		>
			{props.title}
		</Button>
	);
};

export default SidebarLink;
