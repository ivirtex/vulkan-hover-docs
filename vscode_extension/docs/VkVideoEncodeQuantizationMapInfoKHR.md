# VkVideoEncodeQuantizationMapInfoKHR(3) Manual Page

## Name

VkVideoEncodeQuantizationMapInfoKHR - Structure specifying quantization map information to use for video encode operations



## [](#_c_specification)C Specification

The `VkVideoEncodeQuantizationMapInfoKHR` structure **can** be included in the `pNext` chain of the [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html) structure passed to the [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEncodeVideoKHR.html) command to specify the quantization map used by the issued video encode operations.

The `VkVideoEncodeQuantizationMapInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_quantization_map
typedef struct VkVideoEncodeQuantizationMapInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        quantizationMap;
    VkExtent2D         quantizationMapExtent;
} VkVideoEncodeQuantizationMapInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `quantizationMap` specifies the image view to use as the quantization map.
- `quantizationMapExtent` specifies the extent of the image subregion of `quantizationMap` to use as the quantization map starting at offset (0,0).

## [](#_description)Description

Valid Usage

- [](#VUID-VkVideoEncodeQuantizationMapInfoKHR-quantizationMapExtent-10352)VUID-VkVideoEncodeQuantizationMapInfoKHR-quantizationMapExtent-10352  
  `quantizationMapExtent.width` **must** be less than or equal to the width of `quantizationMap`
- [](#VUID-VkVideoEncodeQuantizationMapInfoKHR-quantizationMapExtent-10353)VUID-VkVideoEncodeQuantizationMapInfoKHR-quantizationMapExtent-10353  
  `quantizationMapExtent.height` **must** be less than or equal to the height of `quantizationMap`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeQuantizationMapInfoKHR-sType-sType)VUID-VkVideoEncodeQuantizationMapInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUANTIZATION_MAP_INFO_KHR`
- [](#VUID-VkVideoEncodeQuantizationMapInfoKHR-quantizationMap-parameter)VUID-VkVideoEncodeQuantizationMapInfoKHR-quantizationMap-parameter  
  If `quantizationMap` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `quantizationMap` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_quantization\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_quantization_map.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeQuantizationMapInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0