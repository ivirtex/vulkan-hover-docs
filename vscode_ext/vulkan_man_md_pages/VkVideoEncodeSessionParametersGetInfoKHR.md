# VkVideoEncodeSessionParametersGetInfoKHR(3) Manual Page

## Name

VkVideoEncodeSessionParametersGetInfoKHR - Structure specifying
parameters for retrieving encoded video session parameter data



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeSessionParametersGetInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeSessionParametersGetInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkVideoSessionParametersKHR    videoSessionParameters;
} VkVideoEncodeSessionParametersGetInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `videoSessionParameters` is the
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) object
  to retrieve encoded parameter data from.

## <a href="#_description" class="anchor"></a>Description

Depending on the used video encode operation, additional codec-specific
structures **may** need to be included in the `pNext` chain of this
structure to identify the specific video session parameters to retrieve
encoded parameter data for, as described in the corresponding sections.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_PARAMETERS_GET_INFO_KHR`

- <a href="#VUID-VkVideoEncodeSessionParametersGetInfoKHR-pNext-pNext"
  id="VUID-VkVideoEncodeSessionParametersGetInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeSessionParametersGetInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersGetInfoKHR.html)
  or
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)

- <a href="#VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-unique"
  id="VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-unique"></a>
  VUID-VkVideoEncodeSessionParametersGetInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a
  href="#VUID-VkVideoEncodeSessionParametersGetInfoKHR-videoSessionParameters-parameter"
  id="VUID-VkVideoEncodeSessionParametersGetInfoKHR-videoSessionParameters-parameter"></a>
  VUID-VkVideoEncodeSessionParametersGetInfoKHR-videoSessionParameters-parameter  
  `videoSessionParameters` **must** be a valid
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html),
[vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetEncodedVideoSessionParametersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeSessionParametersGetInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
