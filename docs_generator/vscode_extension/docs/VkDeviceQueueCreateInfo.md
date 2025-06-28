# VkDeviceQueueCreateInfo(3) Manual Page

## Name

VkDeviceQueueCreateInfo - Structure specifying parameters of a newly created device queue



## [](#_c_specification)C Specification

The `VkDeviceQueueCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDeviceQueueCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkDeviceQueueCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
    uint32_t                    queueCount;
    const float*                pQueuePriorities;
} VkDeviceQueueCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask indicating behavior of the queues.
- `queueFamilyIndex` is an unsigned integer indicating the index of the queue family in which to create the queues on this device. This index corresponds to the index of an element of the `pQueueFamilyProperties` array that was returned by `vkGetPhysicalDeviceQueueFamilyProperties`.
- `queueCount` is an unsigned integer specifying the number of queues to create in the queue family indicated by `queueFamilyIndex`, and with the behavior specified by `flags`.
- `pQueuePriorities` is a pointer to an array of `queueCount` normalized floating-point values, specifying priorities of work that will be submitted to each created queue. See [Queue Priority](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-priority) for more information.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDeviceQueueCreateInfo-queueFamilyIndex-00381)VUID-VkDeviceQueueCreateInfo-queueFamilyIndex-00381  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount` returned by `vkGetPhysicalDeviceQueueFamilyProperties`
- [](#VUID-VkDeviceQueueCreateInfo-queueCount-00382)VUID-VkDeviceQueueCreateInfo-queueCount-00382  
  `queueCount` **must** be less than or equal to the `queueCount` member of the `VkQueueFamilyProperties` structure, as returned by `vkGetPhysicalDeviceQueueFamilyProperties` in the `pQueueFamilyProperties`\[queueFamilyIndex]
- [](#VUID-VkDeviceQueueCreateInfo-pQueuePriorities-00383)VUID-VkDeviceQueueCreateInfo-pQueuePriorities-00383  
  Each element of `pQueuePriorities` **must** be between `0.0` and `1.0` inclusive
- [](#VUID-VkDeviceQueueCreateInfo-flags-02861)VUID-VkDeviceQueueCreateInfo-flags-02861  
  If the [`protectedMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-protectedMemory) feature is not enabled, the `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT` bit of `flags` **must** not be set
- [](#VUID-VkDeviceQueueCreateInfo-flags-06449)VUID-VkDeviceQueueCreateInfo-flags-06449  
  If `flags` includes `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT`, `queueFamilyIndex` **must** be the index of a queue family that includes the `VK_QUEUE_PROTECTED_BIT` capability
- [](#VUID-VkDeviceQueueCreateInfo-pNext-09398)VUID-VkDeviceQueueCreateInfo-pNext-09398  
  If the `pNext` chain includes a [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html) structure then [VkPhysicalDeviceSchedulingControlsPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSchedulingControlsPropertiesARM.html)::`schedulingControlsFlags` **must** contain `VK_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_SHADER_CORE_COUNT_ARM`

Valid Usage (Implicit)

- [](#VUID-VkDeviceQueueCreateInfo-sType-sType)VUID-VkDeviceQueueCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO`
- [](#VUID-VkDeviceQueueCreateInfo-pNext-pNext)VUID-VkDeviceQueueCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDeviceQueueGlobalPriorityCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueGlobalPriorityCreateInfo.html) or [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)
- [](#VUID-VkDeviceQueueCreateInfo-sType-unique)VUID-VkDeviceQueueCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDeviceQueueCreateInfo-flags-parameter)VUID-VkDeviceQueueCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkDeviceQueueCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateFlagBits.html) values
- [](#VUID-VkDeviceQueueCreateInfo-pQueuePriorities-parameter)VUID-VkDeviceQueueCreateInfo-pQueuePriorities-parameter  
  `pQueuePriorities` **must** be a valid pointer to an array of `queueCount` `float` values
- [](#VUID-VkDeviceQueueCreateInfo-queueCount-arraylength)VUID-VkDeviceQueueCreateInfo-queueCount-arraylength  
  `queueCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html), [VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceQueueCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceQueueCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0