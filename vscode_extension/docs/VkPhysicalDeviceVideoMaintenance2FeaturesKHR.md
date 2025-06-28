# VkPhysicalDeviceVideoMaintenance2FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceVideoMaintenance2FeaturesKHR - Structure describing the video maintenance features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVideoMaintenance2FeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_maintenance2
typedef struct VkPhysicalDeviceVideoMaintenance2FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           videoMaintenance2;
} VkPhysicalDeviceVideoMaintenance2FeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`videoMaintenance2` specifies that the implementation supports the following:
  
  - Support for issuing [video coding control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding-control) commands against video decode sessions without a bound video session parameters object.
  - The new video session creation flag `VK_VIDEO_SESSION_CREATE_INLINE_SESSION_PARAMETERS_BIT_KHR` for video decode sessions.
  - Required support for the [rate control mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control-modes) `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR` for the following video encode profiles:
    
    - [H.264 encode profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h264-profile);
    - [H.265 encode profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-h265-profile);
    - [AV1 encode profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-av1-profile).
  - Additional guarantees on Video Std parameters used with video encode profiles that the implementations support without the need to [override](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-overrides) them.

## [](#_description)Description

If the `VkPhysicalDeviceVideoMaintenance2FeaturesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceVideoMaintenance2FeaturesKHR`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVideoMaintenance2FeaturesKHR-sType-sType)VUID-VkPhysicalDeviceVideoMaintenance2FeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_MAINTENANCE_2_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_maintenance2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVideoMaintenance2FeaturesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0