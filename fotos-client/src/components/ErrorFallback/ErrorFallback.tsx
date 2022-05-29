import React from "react";

const ErrorFallback = (props: ErrorFallbackPropType) => {
	console.error(props.error);
	return <div>Some Error Occured</div>;
};

export default ErrorFallback;
