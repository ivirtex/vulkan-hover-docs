# VkVideoDecodeVP9CapabilitiesKHR(3) Manual Page

## Name

VkVideoDecodeVP9CapabilitiesKHR - Structure describing VP9 decode capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities for an [VP9 decode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-vp9-profile), the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain **must** include a `VkVideoDecodeVP9CapabilitiesKHR` structure that will be filled with the profile-specific capabilities.

The `VkVideoDecodeVP9CapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_vp9
typedef struct VkVideoDecodeVP9CapabilitiesKHR {
    VkStructureType     sType;
    void*               pNext;
    StdVideoVP9Level    maxLevel;
} VkVideoDecodeVP9CapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxLevel` is a `StdVideoVP9Level` value specifying the maximum VP9 level supported by the profile, as defined in section A.1 of the [VP9 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#google-vp9).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeVP9CapabilitiesKHR-sType-sType)VUID-VkVideoDecodeVP9CapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_VP9_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_vp9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_vp9.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeVP9CapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0