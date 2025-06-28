# VkVideoEncodeQuantizationMapCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeQuantizationMapCapabilitiesKHR - Structure describing video encode quantization map capabilities for a video profile



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) with `pVideoProfile->videoCodecOperation` specifying an encode operation, the [VkVideoEncodeQuantizationMapCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapCapabilitiesKHR.html) structure **can** be included in the `pNext` chain of the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure to retrieve capabilities specific to video encode quantization maps.

The `VkVideoEncodeQuantizationMapCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef struct VkVideoEncodeQuantizationMapCapabilitiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         maxQuantizationMapExtent;
} VkVideoEncodeQuantizationMapCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxQuantizationMapExtent` indicates the maximum supported width and height of quantization maps.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeQuantizationMapCapabilitiesKHR-sType-sType)VUID-VkVideoEncodeQuantizationMapCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUANTIZATION_MAP_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeQuantizationMapCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0