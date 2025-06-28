# VkVideoFormatAV1QuantizationMapPropertiesKHR(3) Manual Page

## Name

VkVideoFormatAV1QuantizationMapPropertiesKHR - Structure describing AV1 quantization map properties



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html), the `VkVideoFormatAV1QuantizationMapPropertiesKHR` structure **can** be included in the `pNext` chain of the [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structure to retrieve video format properties specific to video encode quantization maps used with an [AV1 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-profile).

The `VkVideoFormatAV1QuantizationMapPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1 with VK_KHR_video_encode_quantization_map
typedef struct VkVideoFormatAV1QuantizationMapPropertiesKHR {
    VkStructureType                           sType;
    void*                                     pNext;
    VkVideoEncodeAV1SuperblockSizeFlagsKHR    compatibleSuperblockSizes;
} VkVideoFormatAV1QuantizationMapPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `compatibleSuperblockSizes` is a bitmask of [VkVideoEncodeAV1SuperblockSizeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SuperblockSizeFlagBitsKHR.html) indicating the AV1 superblock sizes that quantization maps using this video format are compatible with.
  
  Note
  
  The value of `compatibleSuperblockSizes` does not limit the use of the specific quantization map format, but does limit the implementation in being able to encode pictures with superblock sizes not included in `compatibleSuperblockSizes` but otherwise supported by the used video profile, as indicated by [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`superblockSizes`. In particular, using smaller [quantization map texel sizes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) may prevent implementations from encoding with larger superblock sizes which may have a negative impact on the efficiency of the encoder.

## [](#_description)Description

The values returned in this structure are only defined if the allowed image usage flags returned in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`imageUsageFlags` for this video format include `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`.

Valid Usage (Implicit)

- [](#VUID-VkVideoFormatAV1QuantizationMapPropertiesKHR-sType-sType)VUID-VkVideoFormatAV1QuantizationMapPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_FORMAT_AV1_QUANTIZATION_MAP_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeAV1SuperblockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1SuperblockSizeFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoFormatAV1QuantizationMapPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0