# vkGetPhysicalDeviceDisplayPlanePropertiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceDisplayPlanePropertiesKHR - Query the plane
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

Images are presented to individual planes on a display. Devices **must**
support at least one plane on each display. Planes **can** be stacked
and blended to composite multiple images on one display. Devices **may**
support only a fixed stacking order and fixed mapping between planes and
displays, or they **may** allow arbitrary application specified stacking
orders and mappings between planes and displays. To query the properties
of device display planes, call:

``` c
// Provided by VK_KHR_display
VkResult vkGetPhysicalDeviceDisplayPlanePropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkDisplayPlanePropertiesKHR*                pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is a physical device.

- `pPropertyCount` is a pointer to an integer related to the number of
  display planes available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  `VkDisplayPlanePropertiesKHR` structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of display planes available
for `physicalDevice` is returned in `pPropertyCount`. Otherwise,
`pPropertyCount` **must** point to a variable set by the user to the
number of elements in the `pProperties` array, and on return the
variable is overwritten with the number of structures actually written
to `pProperties`. If the value of `pPropertyCount` is less than the
number of display planes for `physicalDevice`, at most `pPropertyCount`
structures will be written.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-pPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-pPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceDisplayPlanePropertiesKHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlanePropertiesKHR.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlanePropertiesKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceDisplayPlanePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
