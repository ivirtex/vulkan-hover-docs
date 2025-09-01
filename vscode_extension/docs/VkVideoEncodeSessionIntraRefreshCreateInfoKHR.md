# VkVideoEncodeSessionIntraRefreshCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeSessionIntraRefreshCreateInfoKHR - Video encode session intra refresh parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeSessionIntraRefreshCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_intra_refresh
typedef struct VkVideoEncodeSessionIntraRefreshCreateInfoKHR {
    VkStructureType                             sType;
    const void*                                 pNext;
    VkVideoEncodeIntraRefreshModeFlagBitsKHR    intraRefreshMode;
} VkVideoEncodeSessionIntraRefreshCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `intraRefreshMode` is a [VkVideoEncodeIntraRefreshModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagBitsKHR.html) specifying the used [intra refresh mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-intra-refresh-modes).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeSessionIntraRefreshCreateInfoKHR-sType-sType)VUID-VkVideoEncodeSessionIntraRefreshCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_INTRA_REFRESH_CREATE_INFO_KHR`
- [](#VUID-VkVideoEncodeSessionIntraRefreshCreateInfoKHR-intraRefreshMode-parameter)VUID-VkVideoEncodeSessionIntraRefreshCreateInfoKHR-intraRefreshMode-parameter  
  If `intraRefreshMode` is not `0`, `intraRefreshMode` **must** be a valid [VkVideoEncodeIntraRefreshModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagBitsKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_intra\_refresh](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_intra_refresh.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeIntraRefreshModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeIntraRefreshModeFlagBitsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeSessionIntraRefreshCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0