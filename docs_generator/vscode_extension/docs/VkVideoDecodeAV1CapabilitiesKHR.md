# VkVideoDecodeAV1CapabilitiesKHR(3) Manual Page

## Name

VkVideoDecodeAV1CapabilitiesKHR - Structure describing AV1 decode capabilities



## [](#_c_specification)C Specification

When calling [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) to query the capabilities for an [AV1 decode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-av1-profile), the [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain **must** include a `VkVideoDecodeAV1CapabilitiesKHR` structure that will be filled with the profile-specific capabilities.

The `VkVideoDecodeAV1CapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1CapabilitiesKHR {
    VkStructureType     sType;
    void*               pNext;
    StdVideoAV1Level    maxLevel;
} VkVideoDecodeAV1CapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxLevel` is a `StdVideoAV1Level` value specifying the maximum AV1 level supported by the profile, as defined in section A.3 of the [AV1 Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#aomedia-av1).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoDecodeAV1CapabilitiesKHR-sType-sType)VUID-VkVideoDecodeAV1CapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_decode\_av1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_decode_av1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoDecodeAV1CapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0