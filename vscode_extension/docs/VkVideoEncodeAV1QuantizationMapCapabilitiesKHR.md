# VkVideoEncodeAV1QuantizationMapCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeAV1QuantizationMapCapabilitiesKHR - Structure describing AV1 encode quantization map capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities of an [AV1 encode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-profile), the `VkVideoEncodeAV1QuantizationMapCapabilitiesKHR` structure **can** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve additional video encode quantization map capabilities specific to AV1 encode profiles.

The `VkVideoEncodeAV1QuantizationMapCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1 with VK_KHR_video_encode_quantization_map
typedef struct VkVideoEncodeAV1QuantizationMapCapabilitiesKHR {
    VkStructureType    sType;
    void*              pNext;
    int32_t            minQIndexDelta;
    int32_t            maxQIndexDelta;
} VkVideoEncodeAV1QuantizationMapCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `minQIndexDelta` indicates the minimum quantizer index delta value supported for [AV1 quantizer index delta maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-qindex-delta-map).
- `maxQIndexDelta` indicates the maximum quantizer index delta value supported for [AV1 quantizer index delta maps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-qindex-delta-map).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1QuantizationMapCapabilitiesKHR-sType-sType)VUID-VkVideoEncodeAV1QuantizationMapCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_QUANTIZATION_MAP_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1QuantizationMapCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0