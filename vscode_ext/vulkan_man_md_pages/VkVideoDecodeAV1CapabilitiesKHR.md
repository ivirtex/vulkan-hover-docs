# VkVideoDecodeAV1CapabilitiesKHR(3) Manual Page

## Name

VkVideoDecodeAV1CapabilitiesKHR - Structure describing AV1 decode
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
to query the capabilities for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-av1-profile"
target="_blank" rel="noopener">AV1 decode profile</a>, the
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain
**must** include a `VkVideoDecodeAV1CapabilitiesKHR` structure that will
be filled with the profile-specific capabilities.

The `VkVideoDecodeAV1CapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_av1
typedef struct VkVideoDecodeAV1CapabilitiesKHR {
    VkStructureType     sType;
    void*               pNext;
    StdVideoAV1Level    maxLevel;
} VkVideoDecodeAV1CapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxLevel` is a `StdVideoAV1Level` value specifying the maximum AV1
  level supported by the profile, as defined in section A.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#aomedia-av1"
  target="_blank" rel="noopener">AV1 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeAV1CapabilitiesKHR-sType-sType"
  id="VUID-VkVideoDecodeAV1CapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoDecodeAV1CapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_AV1_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_av1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_av1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeAV1CapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
