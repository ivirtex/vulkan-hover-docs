# vkGetPhysicalDeviceQueueFamilyProperties(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyProperties - Reports properties of the queues of the specified physical device



## [](#_c_specification)C Specification

To query properties of queues available on a physical device, call:

Warning

This functionality is deprecated by [Vulkan Version 1.1](#versions-1.1). See [Deprecated Functionality](#deprecation-gpdp2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceQueueFamilyProperties(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pQueueFamilyPropertyCount,
    VkQueueFamilyProperties*                    pQueueFamilyProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the physical device whose properties will be queried.
- `pQueueFamilyPropertyCount` is a pointer to an integer related to the number of queue families available or queried, as described below.
- `pQueueFamilyProperties` is either `NULL` or a pointer to an array of [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html) structures.

## [](#_description)Description

If `pQueueFamilyProperties` is `NULL`, then the number of queue families available is returned in `pQueueFamilyPropertyCount`. Implementations **must** support at least one queue family. Otherwise, `pQueueFamilyPropertyCount` **must** point to a variable set by the application to the number of elements in the `pQueueFamilyProperties` array, and on return the variable is overwritten with the number of structures actually written to `pQueueFamilyProperties`. If `pQueueFamilyPropertyCount` is less than the number of queue families available, at most `pQueueFamilyPropertyCount` structures will be written.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceQueueFamilyProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceQueueFamilyProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyPropertyCount-parameter)VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyPropertyCount-parameter  
  `pQueueFamilyPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyProperties-parameter)VUID-vkGetPhysicalDeviceQueueFamilyProperties-pQueueFamilyProperties-parameter  
  If the value referenced by `pQueueFamilyPropertyCount` is not `0`, and `pQueueFamilyProperties` is not `NULL`, `pQueueFamilyProperties` **must** be a valid pointer to an array of `pQueueFamilyPropertyCount` [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0