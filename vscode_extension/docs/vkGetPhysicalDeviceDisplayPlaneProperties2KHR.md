# vkGetPhysicalDeviceDisplayPlaneProperties2KHR(3) Manual Page

## Name

vkGetPhysicalDeviceDisplayPlaneProperties2KHR - Query information about the available display planes.



## [](#_c_specification)C Specification

To query the properties of a device’s display planes, call:

```c++
// Provided by VK_KHR_get_display_properties2
VkResult vkGetPhysicalDeviceDisplayPlaneProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkDisplayPlaneProperties2KHR*               pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is a physical device.
- `pPropertyCount` is a pointer to an integer related to the number of display planes available or queried, as described below.
- `pProperties` is either `NULL` or a pointer to an array of `VkDisplayPlaneProperties2KHR` structures.

## [](#_description)Description

`vkGetPhysicalDeviceDisplayPlaneProperties2KHR` behaves similarly to [vkGetPhysicalDeviceDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayPlanePropertiesKHR.html), with the ability to return extended information via chained output structures.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pPropertyCount-parameter)VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pProperties-parameter)VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneProperties2KHR.html) structures

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html), [VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneProperties2KHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceDisplayPlaneProperties2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0