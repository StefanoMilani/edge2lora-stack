// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var ProcessUplinkMessageRequestFieldPathsNested = []string{
	"end_device_version_ids",
	"end_device_version_ids.brand_id",
	"end_device_version_ids.firmware_version",
	"end_device_version_ids.hardware_version",
	"end_device_version_ids.model_id",
	"ids",
	"ids.application_ids",
	"ids.application_ids.application_id",
	"ids.dev_addr",
	"ids.dev_eui",
	"ids.device_id",
	"ids.join_eui",
	"message",
	"message.decoded_payload",
	"message.f_cnt",
	"message.f_port",
	"message.frm_payload",
	"message.received_at",
	"message.rx_metadata",
	"message.session_key_id",
	"message.settings",
	"message.settings.coding_rate",
	"message.settings.data_rate",
	"message.settings.data_rate.modulation",
	"message.settings.data_rate.modulation.fsk",
	"message.settings.data_rate.modulation.fsk.bit_rate",
	"message.settings.data_rate.modulation.lora",
	"message.settings.data_rate.modulation.lora.bandwidth",
	"message.settings.data_rate.modulation.lora.spreading_factor",
	"message.settings.data_rate_index",
	"message.settings.downlink",
	"message.settings.downlink.antenna_index",
	"message.settings.downlink.invert_polarization",
	"message.settings.downlink.tx_power",
	"message.settings.enable_crc",
	"message.settings.frequency",
	"message.settings.time",
	"message.settings.timestamp",
	"parameter",
}

var ProcessUplinkMessageRequestFieldPathsTopLevel = []string{
	"end_device_version_ids",
	"ids",
	"message",
	"parameter",
}
var ProcessDownlinkMessageRequestFieldPathsNested = []string{
	"end_device_version_ids",
	"end_device_version_ids.brand_id",
	"end_device_version_ids.firmware_version",
	"end_device_version_ids.hardware_version",
	"end_device_version_ids.model_id",
	"ids",
	"ids.application_ids",
	"ids.application_ids.application_id",
	"ids.dev_addr",
	"ids.dev_eui",
	"ids.device_id",
	"ids.join_eui",
	"message",
	"message.class_b_c",
	"message.class_b_c.absolute_time",
	"message.class_b_c.gateways",
	"message.confirmed",
	"message.correlation_ids",
	"message.decoded_payload",
	"message.f_cnt",
	"message.f_port",
	"message.frm_payload",
	"message.priority",
	"message.session_key_id",
	"parameter",
}

var ProcessDownlinkMessageRequestFieldPathsTopLevel = []string{
	"end_device_version_ids",
	"ids",
	"message",
	"parameter",
}
