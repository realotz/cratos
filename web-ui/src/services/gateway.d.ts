// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

declare namespace CratosApiV1Gateway {
	/** ListOption */
	type ListOption = {
		limit?:number
		Offset?:number
		sort?:string
		name?:string
		namespace?:string
	}
	/** GatewayList */
	type GatewayList = {
		list?:Array<CratosApiV1Gateway.Gateway>
		total?:number
	}
	/** Gateway */
	type Gateway = {
		apiVersion?:string
		kind?:string
		spec?:IstioNetworkingV1alpha3.Gateway
		metadata:K8sIoApimachineryPkgApisMetaV1.ObjectMeta
	}
	/** GetKind */
	type GetKind = {
		name?:string
		namespace?:string
		version?:string
	}
	/** DeleteKind */
	type DeleteKind = {
		name?:string
		namespace?:string
	}
	/** Request */
	type Request = {
	}
	/** Response */
	type Response = {
	}
}

declare namespace IstioNetworkingV1alpha3 {
	/** Gateway */
	type Gateway = {
		servers?:Array<IstioNetworkingV1alpha3.Server>
		selector?:Map<string,string>
	}
	/** Server */
	type Server = {
		port?:IstioNetworkingV1alpha3.Port
		bind?:string
		hosts?:Array<string>
		tls?:IstioNetworkingV1alpha3.ServerTLSSettings
		default_endpoint?:string
		name?:string
	}
	/** SelectorEntry */
	type SelectorEntry = {
		key?:string
		value?:string
	}
	/** Port */
	type Port = {
		number?:number
		protocol?:string
		name?:string
		target_port?:number
	}
	/** ServerTLSSettings */
	type ServerTLSSettings = {
		https_redirect?:boolean
		mode?:Array<any>
		server_certificate?:string
		private_key?:string
		ca_certificates?:string
		credential_name?:string
		subject_alt_names?:Array<string>
		verify_certificate_spki?:Array<string>
		verify_certificate_hash?:Array<string>
		min_protocol_version?:Array<any>
		max_protocol_version?:Array<any>
		cipher_suites?:Array<string>
	}
}

declare namespace K8sIoApimachineryPkgApisMetaV1 {
	/** FieldsV1 */
	type FieldsV1 = {
		Raw?:string
	}
	/** ObjectMeta */
	type ObjectMeta = {
		name?:string
		generateName?:string
		namespace?:string
		selfLink?:string
		uid?:string
		resourceVersion?:string
		generation?:number
		creationTimestamp?:K8sIoApimachineryPkgApisMetaV1.Time
		deletionTimestamp?:K8sIoApimachineryPkgApisMetaV1.Time
		deletionGracePeriodSeconds?:number
		labels?:Map<string,string>
		annotations?:Map<string,string>
		ownerReferences?:Array<K8sIoApimachineryPkgApisMetaV1.OwnerReference>
		finalizers?:Array<string>
		clusterName?:string
		managedFields?:Array<K8sIoApimachineryPkgApisMetaV1.ManagedFieldsEntry>
	}
	/** Time */
	type Time = {
		seconds?:number
		nanos?:number
	}
	/** LabelsEntry */
	type LabelsEntry = {
		key?:string
		value?:string
	}
	/** AnnotationsEntry */
	type AnnotationsEntry = {
		key?:string
		value?:string
	}
	/** OwnerReference */
	type OwnerReference = {
		apiVersion?:string
		kind?:string
		name?:string
		uid?:string
		controller?:boolean
		blockOwnerDeletion?:boolean
	}
	/** ManagedFieldsEntry */
	type ManagedFieldsEntry = {
		manager?:string
		operation?:string
		apiVersion?:string
		time?:K8sIoApimachineryPkgApisMetaV1.Time
		fieldsType?:string
		fieldsV1?:K8sIoApimachineryPkgApisMetaV1.FieldsV1
	}
}

