# VkVideoEncodeH264SessionCreateInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264SessionCreateInfoKHR - Structure specifies H.264 encode
session parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH264SessionCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264SessionCreateInfoKHR {
    VkStructureType         sType;
    const void*             pNext;
    VkBool32                useMaxLevelIdc;
    StdVideoH264LevelIdc    maxLevelIdc;
} VkVideoEncodeH264SessionCreateInfoKHR;
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
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`maxLevelIdc`,
  as reported by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile.

- `maxLevelIdc` is a `StdVideoH264LevelIdc` value specifying the upper
  bound on the H.264 level for the video bitstreams produced by the
  created video session, where enum constant
  `STD_VIDEO_H264_LEVEL_IDC_<major>_<minor>` identifies H.264 level
  `<major>.<minor>` as defined in section A.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264SessionCreateInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264SessionCreateInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264SessionCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_SESSION_CREATE_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264SessionCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
