import { useEffect } from 'react';
import { register } from '../instrumentation';

const OpenTelemetryRegistration = () => {
	useEffect(() => {
    		register();
  	}, []);

  	return null; // This component doesn't render anything
};

export default OpenTelemetryRegistration;

