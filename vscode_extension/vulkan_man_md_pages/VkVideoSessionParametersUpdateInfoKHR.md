# VkVideoSessionParametersUpdateInfoKHR(3) Manual Page

## Name

VkVideoSessionParametersUpdateInfoKHR - Structure specifying video
session parameters update information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoSessionParametersUpdateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoSessionParametersUpdateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           updateSequenceCount;
} VkVideoSessionParametersUpdateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `updateSequenceCount` is the new <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-session-parameters"
  target="_blank" rel="noopener">update sequence count</a> to set for
  the video session parameters object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoSessionParametersUpdateInfoKHR-sType-sType"
  id="VUID-VkVideoSessionParametersUpdateInfoKHR-sType-sType"></a>
  VUID-VkVideoSessionParametersUpdateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_SESSION_PARAMETERS_UPDATE_INFO_KHR`

- <a href="#VUID-VkVideoSessionParametersUpdateInfoKHR-pNext-pNext"
  id="VUID-VkVideoSessionParametersUpdateInfoKHR-pNext-pNext"></a>
  VUID-VkVideoSessionParametersUpdateInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoDecodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264SessionParametersAddInfoKHR.html),
  [VkVideoDecodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265SessionParametersAddInfoKHR.html),
  [VkVideoEncodeH264SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersAddInfoKHR.html),
  or
  [VkVideoEncodeH265SessionParametersAddInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersAddInfoKHR.html)

- <a href="#VUID-VkVideoSessionParametersUpdateInfoKHR-sType-unique"
  id="VUID-VkVideoSessionParametersUpdateInfoKHR-sType-unique"></a>
  VUID-VkVideoSessionParametersUpdateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkUpdateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateVideoSessionParametersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoSessionParametersUpdateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
