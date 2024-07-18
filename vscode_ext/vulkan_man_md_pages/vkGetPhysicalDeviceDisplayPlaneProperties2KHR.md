# vkGetPhysicalDeviceDisplayPlaneProperties2KHR(3) Manual Page

## Name

vkGetPhysicalDeviceDisplayPlaneProperties2KHR - Query information about
the available display planes.



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the properties of a device’s display planes, call:

``` c
// Provided by VK_KHR_get_display_properties2
VkResult vkGetPhysicalDeviceDisplayPlaneProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkDisplayPlaneProperties2KHR*               pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is a physical device.

- `pPropertyCount` is a pointer to an integer related to the number of
  display planes available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  `VkDisplayPlaneProperties2KHR` structures.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceDisplayPlaneProperties2KHR` behaves similarly to
[vkGetPhysicalDeviceDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlanePropertiesKHR.html),
with the ability to return extended information via chained output
structures.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceDisplayPlaneProperties2KHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneProperties2KHR.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_display_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_display_properties2.html),
[VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneProperties2KHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceDisplayPlaneProperties2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
