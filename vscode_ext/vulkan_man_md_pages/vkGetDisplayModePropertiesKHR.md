# vkGetDisplayModePropertiesKHR(3) Manual Page

## Name

vkGetDisplayModePropertiesKHR - Query the set of mode properties
supported by the display



## <a href="#_c_specification" class="anchor"></a>C Specification

Each display has one or more supported modes associated with it by
default. These built-in modes are queried by calling:

``` c
// Provided by VK_KHR_display
VkResult vkGetDisplayModePropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display,
    uint32_t*                                   pPropertyCount,
    VkDisplayModePropertiesKHR*                 pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device associated with `display`.

- `display` is the display to query.

- `pPropertyCount` is a pointer to an integer related to the number of
  display modes available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  `VkDisplayModePropertiesKHR` structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of display modes available
on the specified `display` for `physicalDevice` is returned in
`pPropertyCount`. Otherwise, `pPropertyCount` **must** point to a
variable set by the user to the number of elements in the `pProperties`
array, and on return the variable is overwritten with the number of
structures actually written to `pProperties`. If the value of
`pPropertyCount` is less than the number of display modes for
`physicalDevice`, at most `pPropertyCount` structures will be written,
and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to
indicate that not all the available display modes were returned.

Valid Usage (Implicit)

- <a href="#VUID-vkGetDisplayModePropertiesKHR-physicalDevice-parameter"
  id="VUID-vkGetDisplayModePropertiesKHR-physicalDevice-parameter"></a>
  VUID-vkGetDisplayModePropertiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetDisplayModePropertiesKHR-display-parameter"
  id="VUID-vkGetDisplayModePropertiesKHR-display-parameter"></a>
  VUID-vkGetDisplayModePropertiesKHR-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkGetDisplayModePropertiesKHR-pPropertyCount-parameter"
  id="VUID-vkGetDisplayModePropertiesKHR-pPropertyCount-parameter"></a>
  VUID-vkGetDisplayModePropertiesKHR-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkGetDisplayModePropertiesKHR-pProperties-parameter"
  id="VUID-vkGetDisplayModePropertiesKHR-pProperties-parameter"></a>
  VUID-vkGetDisplayModePropertiesKHR-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModePropertiesKHR.html)
  structures

- <a href="#VUID-vkGetDisplayModePropertiesKHR-display-parent"
  id="VUID-vkGetDisplayModePropertiesKHR-display-parent"></a>
  VUID-vkGetDisplayModePropertiesKHR-display-parent  
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

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModePropertiesKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDisplayModePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
