package android

type Product_variables struct {
	Additional_gralloc_10_usage_bits struct {
		Cppflags []string
	}
	Bootloader_message_offset struct {
		Cflags []string
	}
	Has_legacy_camera_hal1 struct {
		Cflags []string
		Overrides []string
		Shared_libs []string
	}
	Target_process_sdk_version_override struct {
		Cppflags []string
	}
	Target_shim_libs struct {
		Cppflags []string
	}
	Target_uses_color_metadata struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}
	Uses_qti_camera_device struct {
		Cppflags    []string
		Shared_libs []string
	}
}

type ProductVariables struct {
	Additional_gralloc_10_usage_bits      *string `json:",omitempty"`
	Bootloader_message_offset             *int    `json:",omitempty"`
	Has_legacy_camera_hal1                *bool   `json:",omitempty"`
	QTIAudioPath                          *string `json:",omitempty"`
	QTIDisplayPath                        *string `json:",omitempty"`
	QTIMediaPath                          *string `json:",omitempty"`
	Specific_camera_parameter_library     *string `json:",omitempty"`
	Target_process_sdk_version_override   *string `json:",omitempty"`
	Target_shim_libs                      *string `json:",omitempty"`
	Target_uses_color_metadata            *bool   `json:",omitempty"`
	Target_use_sdclang                    *bool   `json:",omitempty"`
	Uses_generic_camera_parameter_library *bool   `json:",omitempty"`
	Uses_qcom_bsp_legacy                  *bool   `json:",omitempty"`
	Uses_qti_camera_device                *bool   `json:",omitempty"`
}
