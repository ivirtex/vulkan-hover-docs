# VkVideoEncodeSessionParametersFeedbackInfoKHR(3) Manual Page

## Name

VkVideoEncodeSessionParametersFeedbackInfoKHR - Structure providing
feedback about the requested video session parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeSessionParametersFeedbackInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeSessionParametersFeedbackInfoKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hasOverrides;
} VkVideoEncodeSessionParametersFeedbackInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `hasOverrides` indicates whether any of the requested parameter data
  were <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">overridden</a> by the implementation.

## <a href="#_description" class="anchor"></a>Description

Depending on the used video encode operation, additional codec-specific
structures **can** be included in the `pNext` chain of this structure to
capture codec-specific feedback information about the requested
parameter data, as described in the corresponding sections.

Valid Usage (Implicit)

- <a
  href="#VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`

- <a
  href="#VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-pNext-pNext"
  id="VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264SessionParametersFeedbackInfoKHR.html)
  or
  [VkVideoEncodeH265SessionParametersFeedbackInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersFeedbackInfoKHR.html)

- <a
  href="#VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-unique"
  id="VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-unique"></a>
  VUID-VkVideoEncodeSessionParametersFeedbackInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetEncodedVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetEncodedVideoSessionParametersKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeSessionParametersFeedbackInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
