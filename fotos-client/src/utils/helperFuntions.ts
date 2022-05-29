import { useDispatch, useSelector } from "react-redux";
import { loginSuccess } from "../actions/user";
import { ImageType, User } from "../types";

const getSize = (window: Window): "small" | "medium" | "large" => {
	if (window.innerWidth <= 640) {
		return "small";
	}
	if (window.innerWidth <= 1024) {
		return "medium";
	}
	return "large";
};

export const getUser = (): User => useSelector((state: any) => state.user);

export const getImages = (eventId: number) => {
	const allImages: Map<number, ImageType[]> = useSelector(
		(state: any) => state.image
	);
	if (allImages && allImages instanceof Map && allImages.has(eventId)) {
		return allImages.get(eventId);
	}
	return [];
};

const removeAscii = (string: string) => {
	return string.replace(/[^\x00-\x7F]/g, "");
};

export const validate = (
	name: string,
	username: string,
	email: string,
	password: string,
	repeat: string
) => {
	let errors = [];
	if (!name) {
		errors.push("Name field is empty");
	}
	if (!username) {
		errors.push("Username field is empty");
	}
	if (username.length < 5 || username.length > 25) {
		errors.push("Username field should be between 5 and 25 characters long");
	}
	if (!email) {
		errors.push("Email field is empty");
	}
	if (!password) {
		errors.push("Password field is empty");
	}
	if (!repeat && password) {
		errors.push("Confirm password field is empty");
	}
	if (repeat !== password && password) {
		errors.push("Passwords do not match");
	}
	let re = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$/;
	if (!re.test(email) && email && email !== removeAscii(email)) {
		errors.push("Incorrect email address");
	}
	re = /^(?=.*\d)(?!.*\s).{6,32}$/;
	if (!re.test(password) && password) {
		errors.push(
			"Password should be between 6 and 32 characters long and should contain at least one number"
		);
	}
	if (username !== removeAscii(username)) {
		errors.push("Incorrect Username");
	}
	if (errors.length > 0) {
		return errors;
	}
	return true;
};

export const dataFetch = async (
	url: string,
	user: User,
	headers: RequestInit["headers"] = {},
	method: string = "POST",
	body?: any
) => {
	headers = {
		...headers,
		Authorization: `Bearer ${user.userToken}`,
	};
	const response = await fetch(url, {
		method,
		headers,
		body: body,
	});
	if (response.status === 401) {
		localStorage.clear();
		window.location.href = "/login";
	}
	return response;
};

export { getSize };
