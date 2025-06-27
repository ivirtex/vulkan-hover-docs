# vkGetPhysicalDeviceToolProperties(3) Manual Page

## Name

vkGetPhysicalDeviceToolProperties - Reports properties of tools active
on the specified physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about tools providing debugging, profiling, or similar
services, active for a given physical device, can be obtained by
calling:

``` c
// Provided by VK_VERSION_1_3
VkResult vkGetPhysicalDeviceToolProperties(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pToolCount,
    VkPhysicalDeviceToolProperties*             pToolProperties);
```

or the equivalent command

``` c
// Provided by VK_EXT_tooling_info
VkResult vkGetPhysicalDeviceToolPropertiesEXT(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pToolCount,
    VkPhysicalDeviceToolProperties*             pToolProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the physical device to query for
  active tools.

- `pToolCount` is a pointer to an integer describing the number of tools
  active on `physicalDevice`.

- `pToolProperties` is either `NULL` or a pointer to an array of
  [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceToolProperties.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pToolProperties` is `NULL`, then the number of tools currently
active on `physicalDevice` is returned in `pToolCount`. Otherwise,
`pToolCount` **must** point to a variable set by the application to the
number of elements in the `pToolProperties` array, and on return the
variable is overwritten with the number of structures actually written
to `pToolProperties`. If `pToolCount` is less than the number of
currently active tools, at most `pToolCount` structures will be written.

The count and properties of active tools **may** change in response to
events outside the scope of the specification. An application **should**
assume these properties might change at any given time.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceToolProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceToolProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceToolProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceToolProperties-pToolCount-parameter"
  id="VUID-vkGetPhysicalDeviceToolProperties-pToolCount-parameter"></a>
  VUID-vkGetPhysicalDeviceToolProperties-pToolCount-parameter  
  `pToolCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceToolProperties-pToolProperties-parameter"
  id="VUID-vkGetPhysicalDeviceToolProperties-pToolProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceToolProperties-pToolProperties-parameter  
  If the value referenced by `pToolCount` is not `0`, and
  `pToolProperties` is not `NULL`, `pToolProperties` **must** be a valid
  pointer to an array of `pToolCount`
  [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceToolProperties.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_tooling_info](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_tooling_info.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceToolProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceToolProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
