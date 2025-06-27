# VkVideoEncodeQualityLevelInfoKHR(3) Manual Page

## Name

VkVideoEncodeQualityLevelInfoKHR - Structure specifying used video
encode quality level



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeQualityLevelInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeQualityLevelInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           qualityLevel;
} VkVideoEncodeQualityLevelInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `qualityLevel` is the used video encode quality level.

## <a href="#_description" class="anchor"></a>Description

This structure **can** be specified in the following places:

- In the `pNext` chain of
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)
  to specify the video encode quality level to use for a video session
  parameters object created for a video encode session. If no instance
  of this structure is included in the `pNext` chain of
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html),
  then the video session parameters object is created with a video
  encode quality level of zero.

- In the `pNext` chain of
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html) to
  change the video encode quality level state of the bound video
  session.

Valid Usage

- <a href="#VUID-VkVideoEncodeQualityLevelInfoKHR-qualityLevel-08311"
  id="VUID-VkVideoEncodeQualityLevelInfoKHR-qualityLevel-08311"></a>
  VUID-VkVideoEncodeQualityLevelInfoKHR-qualityLevel-08311  
  `qualityLevel` **must** be less than
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxQualityLevels`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeQualityLevelInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeQualityLevelInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeQualityLevelInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_QUALITY_LEVEL_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeQualityLevelInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
