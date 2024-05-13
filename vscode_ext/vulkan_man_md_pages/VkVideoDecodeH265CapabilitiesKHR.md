# VkVideoDecodeH265CapabilitiesKHR(3) Manual Page

## Name

VkVideoDecodeH265CapabilitiesKHR - Structure describing H.265 decode
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
to query the capabilities for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-profile"
target="_blank" rel="noopener">H.265 decode profile</a>, the
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain
**must** include a `VkVideoDecodeH265CapabilitiesKHR` structure that
will be filled with the profile-specific capabilities.

The `VkVideoDecodeH265CapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265CapabilitiesKHR {
    VkStructureType         sType;
    void*                   pNext;
    StdVideoH265LevelIdc    maxLevelIdc;
} VkVideoDecodeH265CapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxLevelIdc` is a `StdVideoH265LevelIdc` value indicating the maximum
  H.265 level supported by the profile, where enum constant
  `STD_VIDEO_H265_LEVEL_IDC_<major>_<minor>` identifies H.265 level
  `<major>.<minor>` as defined in section A.4 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH265CapabilitiesKHR-sType-sType"
  id="VUID-VkVideoDecodeH265CapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoDecodeH265CapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH265CapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
