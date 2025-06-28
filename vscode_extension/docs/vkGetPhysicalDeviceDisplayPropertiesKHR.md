# vkGetPhysicalDeviceDisplayPropertiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceDisplayPropertiesKHR - Query information about the available displays



## [](#_c_specification)C Specification

Various functions are provided for enumerating the available display devices present on a Vulkan physical device. To query information about the available displays, call:

```c++
// Provided by VK_KHR_display
VkResult vkGetPhysicalDeviceDisplayPropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkDisplayPropertiesKHR*                     pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is a physical device.
- `pPropertyCount` is a pointer to an integer related to the number of display devices available or queried, as described below.
- `pProperties` is either `NULL` or a pointer to an array of [VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPropertiesKHR.html) structures.

## [](#_description)Description

If `pProperties` is `NULL`, then the number of display devices available for `physicalDevice` is returned in `pPropertyCount`. Otherwise, `pPropertyCount` **must** point to a variable set by the application to the number of elements in the `pProperties` array, and on return the variable is overwritten with the number of structures actually written to `pProperties`. If the value of `pPropertyCount` is less than the number of display devices for `physicalDevice`, at most `pPropertyCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available properties were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceDisplayPropertiesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceDisplayPropertiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceDisplayPropertiesKHR-pPropertyCount-parameter)VUID-vkGetPhysicalDeviceDisplayPropertiesKHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceDisplayPropertiesKHR-pProperties-parameter)VUID-vkGetPhysicalDeviceDisplayPropertiesKHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPropertiesKHR.html) structures

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkDisplayPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPropertiesKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceDisplayPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0