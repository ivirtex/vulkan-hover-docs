# vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR - Reports the number of passes require for a performance query pool type



## [](#_c_specification)C Specification

To query the number of passes required to query a performance query pool on a physical device, call:

```c++
// Provided by VK_KHR_performance_query
void vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkQueryPoolPerformanceCreateInfoKHR*  pPerformanceQueryCreateInfo,
    uint32_t*                                   pNumPasses);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the physical device whose queue family performance query counter properties will be queried.
- `pPerformanceQueryCreateInfo` is a pointer to a `VkQueryPoolPerformanceCreateInfoKHR` of the performance query that is to be created.
- `pNumPasses` is a pointer to an integer related to the number of passes required to query the performance query pool, as described below.

## [](#_description)Description

The `pPerformanceQueryCreateInfo` member `VkQueryPoolPerformanceCreateInfoKHR`::`queueFamilyIndex` **must** be a queue family of `physicalDevice`. The number of passes required to capture the counters specified in the `pPerformanceQueryCreateInfo` member `VkQueryPoolPerformanceCreateInfoKHR`::`pCounters` is returned in `pNumPasses`.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pPerformanceQueryCreateInfo-parameter)VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pPerformanceQueryCreateInfo-parameter  
  `pPerformanceQueryCreateInfo` **must** be a valid pointer to a valid [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pNumPasses-parameter)VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pNumPasses-parameter  
  `pNumPasses` **must** be a valid pointer to a `uint32_t` value

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0