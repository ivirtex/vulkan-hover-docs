# VkVideoReferenceIntraRefreshInfoKHR(3) Manual Page

## Name

VkVideoReferenceIntraRefreshInfoKHR - Structure specifying video reference picture intra refresh parameters



## [](#_c_specification)C Specification

The `VkVideoReferenceIntraRefreshInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_intra_refresh
typedef struct VkVideoReferenceIntraRefreshInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           dirtyIntraRefreshRegions;
} VkVideoReferenceIntraRefreshInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dirtyIntraRefreshRegions` is the number of [dirty intra refresh regions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-dirty-intra-refresh-regions) in the reference picture.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoReferenceIntraRefreshInfoKHR-sType-sType)VUID-VkVideoReferenceIntraRefreshInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_REFERENCE_INTRA_REFRESH_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_intra\_refresh](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_intra_refresh.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoReferenceIntraRefreshInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0