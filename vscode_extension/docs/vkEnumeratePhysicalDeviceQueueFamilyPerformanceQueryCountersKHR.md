# vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(3) Manual Page

## Name

vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR - Reports properties of the performance query counters available on a queue family of a device



## [](#_c_specification)C Specification

To enumerate the performance query counters available on a queue family of a physical device, call:

```c++
// Provided by VK_KHR_performance_query
VkResult vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    uint32_t*                                   pCounterCount,
    VkPerformanceCounterKHR*                    pCounters,
    VkPerformanceCounterDescriptionKHR*         pCounterDescriptions);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the physical device whose queue family performance query counter properties will be queried.
- `queueFamilyIndex` is the index into the queue family of the physical device we want to get properties for.
- `pCounterCount` is a pointer to an integer related to the number of counters available or queried, as described below.
- `pCounters` is either `NULL` or a pointer to an array of [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html) structures.
- `pCounterDescriptions` is either `NULL` or a pointer to an array of [VkPerformanceCounterDescriptionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionKHR.html) structures.

## [](#_description)Description

If `pCounters` is `NULL` and `pCounterDescriptions` is `NULL`, then the number of counters available is returned in `pCounterCount`. Otherwise, `pCounterCount` **must** point to a variable set by the application to the number of elements in the `pCounters`, `pCounterDescriptions`, or both arrays and on return the variable is overwritten with the number of structures actually written out. If `pCounterCount` is less than the number of counters available, at most `pCounterCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available counters were returned.

Valid Usage (Implicit)

- [](#VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-physicalDevice-parameter)VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-pCounterCount-parameter)VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-pCounterCount-parameter  
  `pCounterCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-pCounters-parameter)VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-pCounters-parameter  
  If the value referenced by `pCounterCount` is not `0`, and `pCounters` is not `NULL`, `pCounters` **must** be a valid pointer to an array of `pCounterCount` [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html) structures
- [](#VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-pCounterDescriptions-parameter)VUID-vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR-pCounterDescriptions-parameter  
  If the value referenced by `pCounterCount` is not `0`, and `pCounterDescriptions` is not `NULL`, `pCounterDescriptions` **must** be a valid pointer to an array of `pCounterCount` [VkPerformanceCounterDescriptionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionKHR.html) structures

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INITIALIZATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_performance\_query](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_performance_query.html), [VkPerformanceCounterDescriptionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionKHR.html), [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0