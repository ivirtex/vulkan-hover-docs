# VkVideoEncodeAV1SessionCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeAV1SessionCreateInfoKHR - Structure specifies AV1 encode session parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeAV1SessionCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_av1
typedef struct VkVideoEncodeAV1SessionCreateInfoKHR {
    VkStructureType     sType;
    const void*         pNext;
    VkBool32            useMaxLevel;
    StdVideoAV1Level    maxLevel;
} VkVideoEncodeAV1SessionCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `useMaxLevel` indicates whether the value of `maxLevel` should be used by the implementation. When it is set to `VK_FALSE`, the implementation ignores the value of `maxLevel` and uses the value of [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html)::`maxLevel`, as reported by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile.
- `maxLevel` is a `StdVideoAV1Level` value specifying the upper bound on the AV1 level for the video bitstreams produced by the created video session.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeAV1SessionCreateInfoKHR-sType-sType)VUID-VkVideoEncodeAV1SessionCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_AV1_SESSION_CREATE_INFO_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_av1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeAV1SessionCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0