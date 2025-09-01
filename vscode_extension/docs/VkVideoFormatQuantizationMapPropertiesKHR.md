# VkVideoFormatQuantizationMapPropertiesKHR(3) Manual Page

## Name

VkVideoFormatQuantizationMapPropertiesKHR - Structure describing quantization map properties



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html), the [VkVideoFormatQuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatQuantizationMapPropertiesKHR.html) structure **can** be included in the `pNext` chain of the [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structure to retrieve video format properties specific to video encode quantization maps.

The `VkVideoFormatQuantizationMapPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef struct VkVideoFormatQuantizationMapPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         quantizationMapTexelSize;
} VkVideoFormatQuantizationMapPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `quantizationMapTexelSize` indicates the [quantization map texel size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) of the video format, i.e. the number of pixels corresponding to each quantization map texel.

## [](#_description)Description

The values returned in this structure are only defined if the allowed image usage flags returned in [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`imageUsageFlags` for this video format include `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`.

Implementations **may** support multiple quantization map texel sizes for a particular video format which is indicated by [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html) returning multiple entries with different `quantizationMapTexelSize` values.

Valid Usage (Implicit)

- [](#VUID-VkVideoFormatQuantizationMapPropertiesKHR-sType-sType)VUID-VkVideoFormatQuantizationMapPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_FORMAT_QUANTIZATION_MAP_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoFormatQuantizationMapPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0