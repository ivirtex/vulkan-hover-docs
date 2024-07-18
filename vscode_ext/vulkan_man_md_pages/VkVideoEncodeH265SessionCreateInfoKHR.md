# VkVideoEncodeH265SessionCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265SessionCreateInfoKHR - Structure specifies H.265 encode
session parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH265SessionCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265SessionCreateInfoKHR {
    VkStructureType         sType;
    const void*             pNext;
    VkBool32                useMaxLevelIdc;
    StdVideoH265LevelIdc    maxLevelIdc;
} VkVideoEncodeH265SessionCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `useMaxLevelIdc` indicates whether the value of `maxLevelIdc` should
  be used by the implementation. When it is set to `VK_FALSE`, the
  implementation ignores the value of `maxLevelIdc` and uses the value
  of
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxLevelIdc`,
  as reported by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile.

- `maxLevelIdc` is a `StdVideoH265LevelIdc` value specifying the upper
  bound on the H.265 level for the video bitstreams produced by the
  created video session, where enum constant
  `STD_VIDEO_H265_LEVEL_IDC_<major>_<minor>` identifies H.265 level
  `<major>.<minor>` as defined in section A.4 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265SessionCreateInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265SessionCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265SessionCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_CREATE_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265SessionCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
