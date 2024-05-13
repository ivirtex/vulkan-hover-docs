# VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR - Structure describing
the video encode profile and quality level to query properties for



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    const VkVideoProfileInfoKHR*    pVideoProfile;
    uint32_t                        qualityLevel;
} VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pVideoProfile` is a pointer to a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure
  specifying the video profile to query the video encode quality level
  properties for.

- `qualityLevel` is the video encode quality level to query properties
  for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08259"
  id="VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08259"></a>
  VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08259  
  `pVideoProfile` **must** be a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-support"
  target="_blank" rel="noopener">supported video profile</a>

- <a
  href="#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08260"
  id="VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08260"></a>
  VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08260  
  `pVideoProfile->videoCodecOperation` **must** specify an encode
  operation

- <a
  href="#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-qualityLevel-08261"
  id="VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-qualityLevel-08261"></a>
  VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-qualityLevel-08261  
  `qualityLevel` **must** be less than
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxQualityLevels`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile specified in `pVideoProfile`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-sType-sType"
  id="VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_QUALITY_LEVEL_INFO_KHR`

- <a
  href="#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pNext-pNext"
  id="VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pNext-pNext"></a>
  VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-parameter"
  id="VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-parameter"></a>
  VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-parameter  
  `pVideoProfile` **must** be a valid pointer to a valid
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html),
[vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
