import { Center, Text } from "@mantine/core";
import { useNavigate } from "react-router-dom";
import { getUser } from "../../utils/helperFuntions";

const Home = (): JSX.Element => {
	const navigate = useNavigate();
	const user = getUser();
	if (!user) {
		navigate("/login");
	}
	return (
		<Center className="w-full h-full">
			<Text className="font-title text-h2">Hello, {user.name}</Text>
		</Center>
	);
};

export default Home;
