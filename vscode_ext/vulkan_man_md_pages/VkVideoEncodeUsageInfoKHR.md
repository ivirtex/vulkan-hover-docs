# VkVideoEncodeUsageInfoKHR(3) Manual Page

## Name

VkVideoEncodeUsageInfoKHR - Structure specifying video encode usage
information



## <a href="#_c_specification" class="anchor"></a>C Specification

Additional information about the video encode use case **can** be
provided by adding a `VkVideoEncodeUsageInfoKHR` structure to the
`pNext` chain of [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html).

The `VkVideoEncodeUsageInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeUsageInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    VkVideoEncodeUsageFlagsKHR      videoUsageHints;
    VkVideoEncodeContentFlagsKHR    videoContentHints;
    VkVideoEncodeTuningModeKHR      tuningMode;
} VkVideoEncodeUsageInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `videoUsageHints` is a bitmask of
  [VkVideoEncodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagBitsKHR.html)
  specifying hints about the intended use of the video encode profile.

- `videoContentHints` is a bitmask of
  [VkVideoEncodeContentFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagBitsKHR.html)
  specifying hints about the content to be encoded using the video
  encode profile.

- `tuningMode` is a
  [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeTuningModeKHR.html) value
  specifying the tuning mode to use when encoding with the video
  profile.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeUsageInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeUsageInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeUsageInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_USAGE_INFO_KHR`

- <a href="#VUID-VkVideoEncodeUsageInfoKHR-videoUsageHints-parameter"
  id="VUID-VkVideoEncodeUsageInfoKHR-videoUsageHints-parameter"></a>
  VUID-VkVideoEncodeUsageInfoKHR-videoUsageHints-parameter  
  `videoUsageHints` **must** be a valid combination of
  [VkVideoEncodeUsageFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagBitsKHR.html)
  values

- <a href="#VUID-VkVideoEncodeUsageInfoKHR-videoContentHints-parameter"
  id="VUID-VkVideoEncodeUsageInfoKHR-videoContentHints-parameter"></a>
  VUID-VkVideoEncodeUsageInfoKHR-videoContentHints-parameter  
  `videoContentHints` **must** be a valid combination of
  [VkVideoEncodeContentFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagBitsKHR.html)
  values

- <a href="#VUID-VkVideoEncodeUsageInfoKHR-tuningMode-parameter"
  id="VUID-VkVideoEncodeUsageInfoKHR-tuningMode-parameter"></a>
  VUID-VkVideoEncodeUsageInfoKHR-tuningMode-parameter  
  If `tuningMode` is not `0`, `tuningMode` **must** be a valid
  [VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeTuningModeKHR.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagsKHR.html),
[VkVideoEncodeTuningModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeTuningModeKHR.html),
[VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeUsageInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
