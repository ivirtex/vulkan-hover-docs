# VkVideoEncodeH265SessionCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265SessionCreateInfoKHR - Structure specifies H.265 encode session parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeH265SessionCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265SessionCreateInfoKHR {
    VkStructureType         sType;
    const void*             pNext;
    VkBool32                useMaxLevelIdc;
    StdVideoH265LevelIdc    maxLevelIdc;
} VkVideoEncodeH265SessionCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `useMaxLevelIdc` indicates whether the value of `maxLevelIdc` should be used by the implementation. When it is `VK_FALSE`, the implementation ignores the value of `maxLevelIdc` and uses the value of [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxLevelIdc`, as reported by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile.
- `maxLevelIdc` is a `StdVideoH265LevelIdc` value specifying the upper bound on the H.265 level for the video bitstreams produced by the created video session, where enum constant `STD_VIDEO_H265_LEVEL_IDC_<major>_<minor>` identifies H.265 level `<major>.<minor>` as defined in section A.4 of the [ITU-T H.265 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#itu-t-h265).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeH265SessionCreateInfoKHR-sType-sType)VUID-VkVideoEncodeH265SessionCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_CREATE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_h265](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_h265.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeH265SessionCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0