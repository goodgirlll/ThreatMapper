/* tslint:disable */
/* eslint-disable */
/**
 * Deepfence ThreatMapper
 * Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: community@deepfence.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ModelComputeMetrics } from './ModelComputeMetrics';
import {
    ModelComputeMetricsFromJSON,
    ModelComputeMetricsFromJSONTyped,
    ModelComputeMetricsToJSON,
} from './ModelComputeMetrics';
import type { ModelContainerImage } from './ModelContainerImage';
import {
    ModelContainerImageFromJSON,
    ModelContainerImageFromJSONTyped,
    ModelContainerImageToJSON,
} from './ModelContainerImage';
import type { ModelProcess } from './ModelProcess';
import {
    ModelProcessFromJSON,
    ModelProcessFromJSONTyped,
    ModelProcessToJSON,
} from './ModelProcess';

/**
 * 
 * @export
 * @interface ModelContainer
 */
export interface ModelContainer {
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    cloud_compliance_latest_scan_id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    cloud_compliance_scan_status: string;
    /**
     * 
     * @type {number}
     * @memberof ModelContainer
     */
    cloud_compliances_count: number;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    compliance_latest_scan_id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    compliance_scan_status: string;
    /**
     * 
     * @type {number}
     * @memberof ModelContainer
     */
    compliances_count: number;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    docker_container_name: string;
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof ModelContainer
     */
    docker_labels: { [key: string]: any; };
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    host_name: string;
    /**
     * 
     * @type {ModelContainerImage}
     * @memberof ModelContainer
     */
    image: ModelContainerImage;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    malware_latest_scan_id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    malware_scan_status: string;
    /**
     * 
     * @type {number}
     * @memberof ModelContainer
     */
    malwares_count: number;
    /**
     * 
     * @type {{ [key: string]: any; }}
     * @memberof ModelContainer
     */
    metadata: { [key: string]: any; };
    /**
     * 
     * @type {ModelComputeMetrics}
     * @memberof ModelContainer
     */
    metrics: ModelComputeMetrics;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    node_id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    node_name: string;
    /**
     * 
     * @type {Array<ModelProcess>}
     * @memberof ModelContainer
     */
    processes: Array<ModelProcess> | null;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    secret_latest_scan: string;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    secret_scan_status: string;
    /**
     * 
     * @type {number}
     * @memberof ModelContainer
     */
    secrets_count: number;
    /**
     * 
     * @type {number}
     * @memberof ModelContainer
     */
    vulnerabilities_count: number;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    vulnerability_latest_scan_id: string;
    /**
     * 
     * @type {string}
     * @memberof ModelContainer
     */
    vulnerability_scan_status: string;
}

/**
 * Check if a given object implements the ModelContainer interface.
 */
export function instanceOfModelContainer(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "cloud_compliance_latest_scan_id" in value;
    isInstance = isInstance && "cloud_compliance_scan_status" in value;
    isInstance = isInstance && "cloud_compliances_count" in value;
    isInstance = isInstance && "compliance_latest_scan_id" in value;
    isInstance = isInstance && "compliance_scan_status" in value;
    isInstance = isInstance && "compliances_count" in value;
    isInstance = isInstance && "docker_container_name" in value;
    isInstance = isInstance && "docker_labels" in value;
    isInstance = isInstance && "host_name" in value;
    isInstance = isInstance && "image" in value;
    isInstance = isInstance && "malware_latest_scan_id" in value;
    isInstance = isInstance && "malware_scan_status" in value;
    isInstance = isInstance && "malwares_count" in value;
    isInstance = isInstance && "metadata" in value;
    isInstance = isInstance && "metrics" in value;
    isInstance = isInstance && "node_id" in value;
    isInstance = isInstance && "node_name" in value;
    isInstance = isInstance && "processes" in value;
    isInstance = isInstance && "secret_latest_scan" in value;
    isInstance = isInstance && "secret_scan_status" in value;
    isInstance = isInstance && "secrets_count" in value;
    isInstance = isInstance && "vulnerabilities_count" in value;
    isInstance = isInstance && "vulnerability_latest_scan_id" in value;
    isInstance = isInstance && "vulnerability_scan_status" in value;

    return isInstance;
}

export function ModelContainerFromJSON(json: any): ModelContainer {
    return ModelContainerFromJSONTyped(json, false);
}

export function ModelContainerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelContainer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'cloud_compliance_latest_scan_id': json['cloud_compliance_latest scan_id'],
        'cloud_compliance_scan_status': json['cloud_compliance_scan_status'],
        'cloud_compliances_count': json['cloud_compliances_count'],
        'compliance_latest_scan_id': json['compliance_latest_scan_id'],
        'compliance_scan_status': json['compliance_scan_status'],
        'compliances_count': json['compliances_count'],
        'docker_container_name': json['docker_container_name'],
        'docker_labels': json['docker_labels'],
        'host_name': json['host_name'],
        'image': ModelContainerImageFromJSON(json['image']),
        'malware_latest_scan_id': json['malware_latest_scan_id'],
        'malware_scan_status': json['malware_scan_status'],
        'malwares_count': json['malwares_count'],
        'metadata': json['metadata'],
        'metrics': ModelComputeMetricsFromJSON(json['metrics']),
        'node_id': json['node_id'],
        'node_name': json['node_name'],
        'processes': (json['processes'] === null ? null : (json['processes'] as Array<any>).map(ModelProcessFromJSON)),
        'secret_latest_scan': json['secret_latest_scan'],
        'secret_scan_status': json['secret_scan_status'],
        'secrets_count': json['secrets_count'],
        'vulnerabilities_count': json['vulnerabilities_count'],
        'vulnerability_latest_scan_id': json['vulnerability_latest_scan_id'],
        'vulnerability_scan_status': json['vulnerability_scan_status'],
    };
}

export function ModelContainerToJSON(value?: ModelContainer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'cloud_compliance_latest scan_id': value.cloud_compliance_latest_scan_id,
        'cloud_compliance_scan_status': value.cloud_compliance_scan_status,
        'cloud_compliances_count': value.cloud_compliances_count,
        'compliance_latest_scan_id': value.compliance_latest_scan_id,
        'compliance_scan_status': value.compliance_scan_status,
        'compliances_count': value.compliances_count,
        'docker_container_name': value.docker_container_name,
        'docker_labels': value.docker_labels,
        'host_name': value.host_name,
        'image': ModelContainerImageToJSON(value.image),
        'malware_latest_scan_id': value.malware_latest_scan_id,
        'malware_scan_status': value.malware_scan_status,
        'malwares_count': value.malwares_count,
        'metadata': value.metadata,
        'metrics': ModelComputeMetricsToJSON(value.metrics),
        'node_id': value.node_id,
        'node_name': value.node_name,
        'processes': (value.processes === null ? null : (value.processes as Array<any>).map(ModelProcessToJSON)),
        'secret_latest_scan': value.secret_latest_scan,
        'secret_scan_status': value.secret_scan_status,
        'secrets_count': value.secrets_count,
        'vulnerabilities_count': value.vulnerabilities_count,
        'vulnerability_latest_scan_id': value.vulnerability_latest_scan_id,
        'vulnerability_scan_status': value.vulnerability_scan_status,
    };
}

