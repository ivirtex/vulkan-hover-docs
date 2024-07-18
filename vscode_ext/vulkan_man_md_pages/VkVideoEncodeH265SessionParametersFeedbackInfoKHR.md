# VkVideoEncodeH265SessionParametersFeedbackInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265SessionParametersFeedbackInfoKHR - Structure providing
feedback about the requested H.265 video session parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH265SessionParametersFeedbackInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265SessionParametersFeedbackInfoKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           hasStdVPSOverrides;
    VkBool32           hasStdSPSOverrides;
    VkBool32           hasStdPPSOverrides;
} VkVideoEncodeH265SessionParametersFeedbackInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `hasStdVPSOverrides` indicates whether any of the parameters of the
  requested <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-vps"
  target="_blank" rel="noopener">H.265 video parameter set</a>, if one
  was requested via
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`writeStdVPS`,
  were <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">overridden</a> by the implementation.

- `hasStdSPSOverrides` indicates whether any of the parameters of the
  requested <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-sps"
  target="_blank" rel="noopener">H.265 sequence parameter set</a>, if
  one was requested via
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`writeStdSPS`,
  were <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">overridden</a> by the implementation.

- `hasStdPPSOverrides` indicates whether any of the parameters of the
  requested <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-pps"
  target="_blank" rel="noopener">H.265 picture parameter set</a>, if one
  was requested via
  [VkVideoEncodeH265SessionParametersGetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265SessionParametersGetInfoKHR.html)::`writeStdPPS`,
  were <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">overridden</a> by the implementation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkVideoEncodeH265SessionParametersFeedbackInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265SessionParametersFeedbackInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265SessionParametersFeedbackInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_SESSION_PARAMETERS_FEEDBACK_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265SessionParametersFeedbackInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
