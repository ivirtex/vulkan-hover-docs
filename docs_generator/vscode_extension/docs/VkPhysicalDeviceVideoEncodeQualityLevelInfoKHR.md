# VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR - Structure describing the video encode profile and quality level to query properties for



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR {
    VkStructureType                 sType;
    const void*                     pNext;
    const VkVideoProfileInfoKHR*    pVideoProfile;
    uint32_t                        qualityLevel;
} VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pVideoProfile` is a pointer to a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure specifying the video profile to query the video encode quality level properties for.
- `qualityLevel` is the video encode quality level to query properties for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08259)VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08259  
  `pVideoProfile` **must** be a [supported video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profile-support)
- [](#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08260)VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-08260  
  `pVideoProfile->videoCodecOperation` **must** specify an encode operation
- [](#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-qualityLevel-08261)VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-qualityLevel-08261  
  `qualityLevel` **must** be less than [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxQualityLevels`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile specified in `pVideoProfile`

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-sType-sType)VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_QUALITY_LEVEL_INFO_KHR`
- [](#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pNext-pNext)VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-parameter)VUID-VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR-pVideoProfile-parameter  
  `pVideoProfile` **must** be a valid pointer to a valid [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html), [vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoEncodeQualityLevelPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVideoEncodeQualityLevelInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0