# VkVideoDecodeH264CapabilitiesKHR(3) Manual Page

## Name

VkVideoDecodeH264CapabilitiesKHR - Structure describing H.264 decode
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
to query the capabilities for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
target="_blank" rel="noopener">H.264 decode profile</a>, the
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`pNext` chain
**must** include a `VkVideoDecodeH264CapabilitiesKHR` structure that
will be filled with the profile-specific capabilities.

The `VkVideoDecodeH264CapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264CapabilitiesKHR {
    VkStructureType         sType;
    void*                   pNext;
    StdVideoH264LevelIdc    maxLevelIdc;
    VkOffset2D              fieldOffsetGranularity;
} VkVideoDecodeH264CapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxLevelIdc` is a `StdVideoH264LevelIdc` value indicating the maximum
  H.264 level supported by the profile, where enum constant
  `STD_VIDEO_H264_LEVEL_IDC_<major>_<minor>` identifies H.264 level
  `<major>.<minor>` as defined in section A.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

- `fieldOffsetGranularity` is the minimum alignment for
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)::`codedOffset`
  specified for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> when using
  the picture layout
  `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH264CapabilitiesKHR-sType-sType"
  id="VUID-VkVideoDecodeH264CapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoDecodeH264CapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h264.html),
[VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH264CapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
