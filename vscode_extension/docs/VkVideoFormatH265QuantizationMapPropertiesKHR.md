# VkVideoFormatH265QuantizationMapPropertiesKHR(3) Manual Page

## Name

VkVideoFormatH265QuantizationMapPropertiesKHR - Structure describing H.265 quantization map properties



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html), the `VkVideoFormatH265QuantizationMapPropertiesKHR` structure **can** be included in the `pNext` chain of the [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structure to retrieve video format properties specific to video encode quantization maps used with an [H.265 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-profile).

The `VkVideoFormatH265QuantizationMapPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265 with VK_KHR_video_encode_quantization_map
typedef struct VkVideoFormatH265QuantizationMapPropertiesKHR {
    VkStructureType                     sType;
    void*                               pNext;
    VkVideoEncodeH265CtbSizeFlagsKHR    compatibleCtbSizes;
} VkVideoFormatH265QuantizationMapPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `compatibleCtbSizes` is a bitmask of [VkVideoEncodeH265CtbSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CtbSizeFlagBitsKHR.html) indicating the CTB sizes that quantization maps using this video format are compatible with.
  
  Note
  
  The value of `compatibleCtbSizes` does not limit the use of the specific quantization map format, but does limit the implementation in being able to encode pictures with CTB sizes not included in `compatibleCtbSizes` but otherwise supported by the used video profile, as indicated by [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`ctbSizes`. In particular, using smaller [quantization map texel sizes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) may prevent implementations from encoding with larger CTB sizes which may have a negative impact on the efficiency of the encoder.

## [](#_description)Description

The values returned in this structure are only defined if the allowed image usage flags returned in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`imageUsageFlags` for this video format include `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`.

Valid Usage (Implicit)

- [](#VUID-VkVideoFormatH265QuantizationMapPropertiesKHR-sType-sType)VUID-VkVideoFormatH265QuantizationMapPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_FORMAT_H265_QUANTIZATION_MAP_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoFormatH265QuantizationMapPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0