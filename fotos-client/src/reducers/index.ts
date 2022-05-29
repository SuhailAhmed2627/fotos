import { combineReducers } from "@reduxjs/toolkit";
import user from "./user";
import image from "./image";

export default combineReducers({
	user: user,
	image: image,
});
