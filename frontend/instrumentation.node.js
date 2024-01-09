import { NodeSDK } from '@opentelemetry/sdk-node'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-grpc'
import { Resource } from '@opentelemetry/resources'
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions'
import { SimpleSpanProcessor, TraceIdRatioBasedSampler } from '@opentelemetry/sdk-trace-node'
import { ChannelCredentials } from '@grpc/grpc-js';
import * as fs from 'fs';

const keyPath = '/home/jerome/Documents/projects/asonrythme/app/testObservability/confs/client.key'
const certPath = '/home/jerome/Documents/projects/asonrythme/app/testObservability/confs/client.crt'	
const caPath = '/home/jerome/Documents/projects/asonrythme/app/testObservability/confs/rootCA.crt'

function sslCreds(){
	const clientCert = fs.readFileSync(certPath)
	const clientKey = fs.readFileSync(keyPath)
	const rootCert = fs.readFileSync(caPath)

	return ChannelCredentials.createSsl(rootCert, clientKey, clientCert)	
}

const sdk = new NodeSDK({
	resource: new Resource({
    		[SemanticResourceAttributes.SERVICE_NAME]: 'next-app',
  	}),
  	spanProcessor: new SimpleSpanProcessor(
		new OTLPTraceExporter({
			url: 'localhost:4317', // Replace with your OTLP collector URL
      			headers: {
        			// Add any necessary headers here
      			},
			credentials: sslCreds(),
		}),
	),
	sampler: new TraceIdRatioBasedSampler(0.5),
})

sdk.start()

console.log('OpenTelemetry instrumentation setup completed');
