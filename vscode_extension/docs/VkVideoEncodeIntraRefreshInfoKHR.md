# VkVideoEncodeIntraRefreshInfoKHR(3) Manual Page

## Name

VkVideoEncodeIntraRefreshInfoKHR - Structure specifying video encode intra refresh parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeIntraRefreshInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_intra_refresh
typedef struct VkVideoEncodeIntraRefreshInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           intraRefreshCycleDuration;
    uint32_t           intraRefreshIndex;
} VkVideoEncodeIntraRefreshInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `intraRefreshCycleDuration` is the used [intra refresh cycle duration](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh-cycle-duration).
- `intraRefreshIndex` is the [intra refresh index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh-index) of the encoded picture.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeIntraRefreshInfoKHR-sType-sType)VUID-VkVideoEncodeIntraRefreshInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INTRA_REFRESH_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_intra\_refresh](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_intra_refresh.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeIntraRefreshInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0