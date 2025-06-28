# vkGetPhysicalDeviceToolProperties(3) Manual Page

## Name

vkGetPhysicalDeviceToolProperties - Reports properties of tools active on the specified physical device



## [](#_c_specification)C Specification

Information about tools providing debugging, profiling, or similar services, active for a given physical device, can be obtained by calling:

```c++
// Provided by VK_VERSION_1_3
VkResult vkGetPhysicalDeviceToolProperties(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pToolCount,
    VkPhysicalDeviceToolProperties*             pToolProperties);
```

or the equivalent command

```c++
// Provided by VK_EXT_tooling_info
VkResult vkGetPhysicalDeviceToolPropertiesEXT(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pToolCount,
    VkPhysicalDeviceToolProperties*             pToolProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the physical device to query for active tools.
- `pToolCount` is a pointer to an integer describing the number of tools active on `physicalDevice`.
- `pToolProperties` is either `NULL` or a pointer to an array of [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolProperties.html) structures.

## [](#_description)Description

If `pToolProperties` is `NULL`, then the number of tools currently active on `physicalDevice` is returned in `pToolCount`. Otherwise, `pToolCount` **must** point to a variable set by the application to the number of elements in the `pToolProperties` array, and on return the variable is overwritten with the number of structures actually written to `pToolProperties`. If `pToolCount` is less than the number of currently active tools, at most `pToolCount` structures will be written.

The count and properties of active tools **may** change in response to events outside the scope of the specification. An application **should** assume these properties might change at any given time.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceToolProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceToolProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceToolProperties-pToolCount-parameter)VUID-vkGetPhysicalDeviceToolProperties-pToolCount-parameter  
  `pToolCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceToolProperties-pToolProperties-parameter)VUID-vkGetPhysicalDeviceToolProperties-pToolProperties-parameter  
  If the value referenced by `pToolCount` is not `0`, and `pToolProperties` is not `NULL`, `pToolProperties` **must** be a valid pointer to an array of `pToolCount` [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolProperties.html) structures

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_tooling\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_tooling_info.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceToolProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceToolProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceToolProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0