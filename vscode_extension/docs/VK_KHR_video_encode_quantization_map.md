# VK\_KHR\_video\_encode\_quantization\_map(3) Manual Page

## Name

VK\_KHR\_video\_encode\_quantization\_map - device extension



## [](#_registered_extension_number)Registered Extension Number

554

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html)  
and  
     [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html)  
     or  
     [Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_video\_encode\_av1
- Interacts with VK\_KHR\_video\_encode\_h264
- Interacts with VK\_KHR\_video\_encode\_h265

## [](#_contact)Contact

- Ahmed Abdelkhalek [\[GitHub\]aabdelkh](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_video_encode_quantization_map%5D%20%40aabdelkh%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_video_encode_quantization_map%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_video\_encode\_quantization\_map](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_video_encode_quantization_map.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-09-23

**IP Status**

No known IP claims.

**Contributors**

- Ahmed Abdelkhalek, AMD
- Benjamin Cheng, AMD
- Srinath Kumarapuram, NVIDIA
- Tony Zlatinski, NVIDIA
- Ping Liu, Intel
- Daniel Rakos, RasterGrid
- Lynne Iribarren, Independent

## [](#_description)Description

This extension builds upon the `VK_KHR_video_encode_queue` extension by enabling fine grained control of codec-specific quantization parameters in video encode operations.

More specifically, it adds support for quantization maps:

- Quantization delta maps to directly control the relative value of quantization parameter values on a per-block basis for all rate control modes (including when rate control is disabled).
- Emphasis maps to indirectly control the relative quantization parameter values on a per-block basis when rate control is not disabled and the rate control mode is not configured to the implementation-defined default mode.

This extension is to be used in conjunction with other codec specific video encode extensions that specify the codec specific quantization parameters these maps control.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoEncodeQuantizationMapFeaturesKHR.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeQuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapCapabilitiesKHR.html)
- Extending [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html):
  
  - [VkVideoEncodeQuantizationMapInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapInfoKHR.html)
- Extending [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html):
  
  - [VkVideoFormatQuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatQuantizationMapPropertiesKHR.html)
- Extending [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html):
  
  - [VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR.html)

If [VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html) is supported:

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeAV1QuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1QuantizationMapCapabilitiesKHR.html)
- Extending [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html):
  
  - [VkVideoFormatAV1QuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatAV1QuantizationMapPropertiesKHR.html)

If [VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html) is supported:

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeH264QuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264QuantizationMapCapabilitiesKHR.html)

If [VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html) is supported:

- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeH265QuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265QuantizationMapCapabilitiesKHR.html)
- Extending [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html):
  
  - [VkVideoFormatH265QuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatH265QuantizationMapPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkVideoEncodeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagBitsKHR.html)
- [VkVideoSessionParametersCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagBitsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VIDEO_ENCODE_QUANTIZATION_MAP_EXTENSION_NAME`
- `VK_KHR_VIDEO_ENCODE_QUANTIZATION_MAP_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`
  - `VK_FORMAT_FEATURE_2_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_VIDEO_ENCODE_QUANTIZATION_MAP_KHR`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`
  - `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_QUANTIZATION_MAP_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUANTIZATION_MAP_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUANTIZATION_MAP_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUANTIZATION_MAP_SESSION_PARAMETERS_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_FORMAT_QUANTIZATION_MAP_PROPERTIES_KHR`
- Extending [VkVideoEncodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilityFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_CAPABILITY_EMPHASIS_MAP_BIT_KHR`
  - `VK_VIDEO_ENCODE_CAPABILITY_QUANTIZATION_DELTA_MAP_BIT_KHR`
- Extending [VkVideoEncodeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_WITH_EMPHASIS_MAP_BIT_KHR`
  - `VK_VIDEO_ENCODE_WITH_QUANTIZATION_DELTA_MAP_BIT_KHR`
- Extending [VkVideoSessionCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateFlagBitsKHR.html):
  
  - `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_EMPHASIS_MAP_BIT_KHR`
  - `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`
- Extending [VkVideoSessionParametersCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateFlagBitsKHR.html):
  
  - `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR`

If [VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_QUANTIZATION_MAP_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_FORMAT_AV1_QUANTIZATION_MAP_PROPERTIES_KHR`

If [VK\_KHR\_video\_encode\_h264](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h264.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_QUANTIZATION_MAP_CAPABILITIES_KHR`
- Extending [VkVideoEncodeH264CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilityFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_H264_CAPABILITY_MB_QP_DIFF_WRAPAROUND_BIT_KHR`

If [VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_QUANTIZATION_MAP_CAPABILITIES_KHR`
  - `VK_STRUCTURE_TYPE_VIDEO_FORMAT_H265_QUANTIZATION_MAP_PROPERTIES_KHR`
- Extending [VkVideoEncodeH265CapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilityFlagBitsKHR.html):
  
  - `VK_VIDEO_ENCODE_H265_CAPABILITY_CU_QP_DIFF_WRAPAROUND_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2024-08-21 (Daniel Rakos)
  
  - Internal revisions
- Revision 2, 2024-09-23 (Daniel Rakos)
  
  - Added interactions with VK\_KHR\_video\_encode\_av1

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_video_encode_quantization_map).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0