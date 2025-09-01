# VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR - Structure specifying quantization map texel size for video session parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef struct VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkExtent2D         quantizationMapTexelSize;
} VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `quantizationMapTexelSize` specifies the [quantization map texel size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size) a video session parameters object created with `VK_VIDEO_SESSION_PARAMETERS_CREATE_QUANTIZATION_MAP_COMPATIBLE_BIT_KHR` is compatible with.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR-sType-sType)VUID-VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUANTIZATION_MAP_SESSION_PARAMETERS_CREATE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeQuantizationMapSessionParametersCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0