# vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR - Reports the
number of passes require for a performance query pool type



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the number of passes required to query a performance query pool
on a physical device, call:

``` c
// Provided by VK_KHR_performance_query
void vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkQueryPoolPerformanceCreateInfoKHR*  pPerformanceQueryCreateInfo,
    uint32_t*                                   pNumPasses);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the physical device whose queue
  family performance query counter properties will be queried.

- `pPerformanceQueryCreateInfo` is a pointer to a
  `VkQueryPoolPerformanceCreateInfoKHR` of the performance query that is
  to be created.

- `pNumPasses` is a pointer to an integer related to the number of
  passes required to query the performance query pool, as described
  below.

## <a href="#_description" class="anchor"></a>Description

The `pPerformanceQueryCreateInfo` member
`VkQueryPoolPerformanceCreateInfoKHR`::`queueFamilyIndex` **must** be a
queue family of `physicalDevice`. The number of passes required to
capture the counters specified in the `pPerformanceQueryCreateInfo`
member `VkQueryPoolPerformanceCreateInfoKHR`::`pCounters` is returned in
`pNumPasses`.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pPerformanceQueryCreateInfo-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pPerformanceQueryCreateInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pPerformanceQueryCreateInfo-parameter  
  `pPerformanceQueryCreateInfo` **must** be a valid pointer to a valid
  [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pNumPasses-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pNumPasses-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR-pNumPasses-parameter  
  `pNumPasses` **must** be a valid pointer to a `uint32_t` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
