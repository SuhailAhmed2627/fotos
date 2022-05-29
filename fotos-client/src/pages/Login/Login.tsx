import {
	Center,
	Text,
	Stack,
	PasswordInput,
	InputWrapper,
	Input,
	Button,
	SimpleGrid,
} from "@mantine/core";
import { useState } from "react";
import {
	MdVisibility,
	MdVisibilityOff,
	MdAlternateEmail,
	MdPerson,
	MdDangerous,
	MdLock,
} from "react-icons/md";
import { getUser, validate } from "../../utils/helperFuntions";
import { cleanNotifications, showNotification } from "@mantine/notifications";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { useDispatch } from "react-redux";
import { AnyAction, Dispatch } from "@reduxjs/toolkit";
import { loginSuccess } from "../../actions/user";

const handleLogin = async (
	email: string,
	password: string,
	dispatch: Dispatch<AnyAction>
) => {
	const reponse = await fetch("/api/user/login", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ email, password }),
	});
	const data = await reponse.json();
	if (reponse.status === 200) {
		dispatch(loginSuccess(data));
		if (data.firstLogin) {
			window.location.href = "/first-login";
		} else {
			window.location.href = "/home";
		}
	}
	showNotification({
		message: data.message,
		title: "Error",
		color: "red",
		autoClose: 2000,
	});
};

const handleSignup = async (
	username: string,
	name: string,
	email: string,
	password: string,
	repeat: string,
	setIsLogin: React.Dispatch<React.SetStateAction<boolean>>,
	setPassword: React.Dispatch<React.SetStateAction<string>>
) => {
	cleanNotifications();
	const result: true | string[] = validate(
		name,
		username,
		email,
		password,
		repeat
	);
	if (result !== true) {
		result.forEach((error) => {
			showNotification({
				message: error,
				title: "Error",
				color: "red",
				icon: <MdDangerous />,
				autoClose: 2000,
			});
		});
		return;
	}

	const response = await fetch("/api/user/signup", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			username,
			name,
			email,
			password,
		}),
	});
	const data = await response.json();

	showNotification({
		message: data.message,
		title: response.status === 200 ? "Signed Up" : "Error",
		color: response.status === 200 ? "green" : "red",
		autoClose: 2000,
	});
	if (response.status === 200) {
		setIsLogin(true);
		setPassword("");
	}
	return;
};

const Login = () => {
	const user = getUser();
	const navigate = useNavigate();
	const dispatch = useDispatch();
	const [isLogin, setIsLogin] = useState(true);
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [rePassword, setRePassword] = useState("");
	const [email, setEmail] = useState("");
	const [name, setName] = useState("");
	if (user && user.firstLogin) {
		navigate("/first-login");
	}
	if (user) {
		navigate("/home");
	}
	return (
		<Center className="items-start md:items-center h-full mt-[2%] w-full bg-gray-100 md:bg-transparent">
			<SimpleGrid className="grid flex-col md:flex w-[90%] justify-start md:justify-center min-h-fit h-[70%] md:w-[500px] bg-gray-100 p-5 md:p-10 rounded-3xl">
				<Text className="font-title text-h3 font-semibold">
					{isLogin ? "Login" : "Sign Up"}
				</Text>
				{!isLogin && (
					<>
						<InputWrapper label={"Name"}>
							<Input
								onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
									setName(e.target.value)
								}
								icon={<MdPerson />}
								placeholder={"Enter Name"}
							/>
						</InputWrapper>
						<InputWrapper label={"Username"}>
							<Input
								onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
									setUsername(e.target.value)
								}
								icon={<MdPerson />}
								placeholder={"Enter Username"}
							/>
						</InputWrapper>
					</>
				)}
				<InputWrapper label={"Email"}>
					<Input
						onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
							setEmail(e.target.value)
						}
						icon={<MdAlternateEmail />}
						placeholder={"Enter Email"}
					/>
				</InputWrapper>
				<div>
					<PasswordInput
						label="Password"
						placeholder="Enter Password"
						description={
							!isLogin
								? "Password should be between 6 and 32 characters long and should contain at least one number"
								: ""
						}
						icon={<MdLock />}
						defaultValue={password}
						visibilityToggleIcon={({ reveal, size }) =>
							reveal ? (
								<MdVisibility size={size} />
							) : (
								<MdVisibilityOff size={size} />
							)
						}
						onChange={(e) => setPassword(e.target.value)}
					/>
					{!isLogin && (
						<PasswordInput
							placeholder="Re-Enter Password"
							defaultValue={rePassword}
							icon={<MdLock />}
							visibilityToggleIcon={({ reveal, size }) =>
								reveal ? (
									<MdVisibility size={size} />
								) : (
									<MdVisibilityOff size={size} />
								)
							}
							onChange={(e) => setRePassword(e.target.value)}
						/>
					)}
				</div>
				<Button
					onClick={() => setIsLogin(!isLogin)}
					className="w-[fit-content]"
					compact
					variant="subtle"
				>
					<Text className="font-body text-small">
						{isLogin ? "New user?" : "Already Registered?"} Click here to{" "}
						{isLogin ? "Sign Up" : "Login"}
					</Text>
				</Button>
				<Button
					onClick={() =>
						isLogin
							? handleLogin(email, password, dispatch)
							: handleSignup(
									username,
									name,
									email,
									password,
									rePassword,
									setIsLogin,
									setPassword
							  )
					}
					className="bg-secondary"
					variant="filled"
				>
					{isLogin ? "Login" : "Sign Up"}
				</Button>
			</SimpleGrid>
		</Center>
	);
};

export default Login;
