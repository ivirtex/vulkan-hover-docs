# vkGetDisplayModeProperties2KHR(3) Manual Page

## Name

vkGetDisplayModeProperties2KHR - Query information about the available
display modes.



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the properties of a device’s built-in display modes, call:

``` c
// Provided by VK_KHR_get_display_properties2
VkResult vkGetDisplayModeProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display,
    uint32_t*                                   pPropertyCount,
    VkDisplayModeProperties2KHR*                pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device associated with `display`.

- `display` is the display to query.

- `pPropertyCount` is a pointer to an integer related to the number of
  display modes available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  `VkDisplayModeProperties2KHR` structures.

## <a href="#_description" class="anchor"></a>Description

`vkGetDisplayModeProperties2KHR` behaves similarly to
[vkGetDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayModePropertiesKHR.html),
with the ability to return extended information via chained output
structures.

Valid Usage (Implicit)

- <a href="#VUID-vkGetDisplayModeProperties2KHR-physicalDevice-parameter"
  id="VUID-vkGetDisplayModeProperties2KHR-physicalDevice-parameter"></a>
  VUID-vkGetDisplayModeProperties2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetDisplayModeProperties2KHR-display-parameter"
  id="VUID-vkGetDisplayModeProperties2KHR-display-parameter"></a>
  VUID-vkGetDisplayModeProperties2KHR-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkGetDisplayModeProperties2KHR-pPropertyCount-parameter"
  id="VUID-vkGetDisplayModeProperties2KHR-pPropertyCount-parameter"></a>
  VUID-vkGetDisplayModeProperties2KHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkGetDisplayModeProperties2KHR-pProperties-parameter"
  id="VUID-vkGetDisplayModeProperties2KHR-pProperties-parameter"></a>
  VUID-vkGetDisplayModeProperties2KHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeProperties2KHR.html)
  structures

- <a href="#VUID-vkGetDisplayModeProperties2KHR-display-parent"
  id="VUID-vkGetDisplayModeProperties2KHR-display-parent"></a>
  VUID-vkGetDisplayModeProperties2KHR-display-parent  
  `display` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_display_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_display_properties2.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeProperties2KHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDisplayModeProperties2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
