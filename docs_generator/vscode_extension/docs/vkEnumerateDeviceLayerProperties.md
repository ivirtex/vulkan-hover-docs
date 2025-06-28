# vkEnumerateDeviceLayerProperties(3) Manual Page

## Name

vkEnumerateDeviceLayerProperties - Returns properties of available physical device layers



## [](#_c_specification)C Specification

To enumerate device layers, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkEnumerateDeviceLayerProperties(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkLayerProperties*                          pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device that will be queried.
- `pPropertyCount` is a pointer to an integer related to the number of layer properties available or queried.
- `pProperties` is either `NULL` or a pointer to an array of [VkLayerProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerProperties.html) structures.

## [](#_description)Description

If `pProperties` is `NULL`, then the number of layer properties available is returned in `pPropertyCount`. Otherwise, `pPropertyCount` **must** point to a variable set by the application to the number of elements in the `pProperties` array, and on return the variable is overwritten with the number of structures actually written to `pProperties`. If `pPropertyCount` is less than the number of layer properties available, at most `pPropertyCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available properties were returned.

The list of layers enumerated by `vkEnumerateDeviceLayerProperties` **must** be exactly the sequence of layers enabled for the instance. The members of `VkLayerProperties` for each enumerated layer **must** be the same as the properties when the layer was enumerated by `vkEnumerateInstanceLayerProperties`.

Note

Due to platform details on Android, `vkEnumerateDeviceLayerProperties` may be called with `physicalDevice` equal to `NULL` during layer discovery. This behavior will only be observed by layer implementations, and not the underlying Vulkan driver.

Valid Usage (Implicit)

- [](#VUID-vkEnumerateDeviceLayerProperties-physicalDevice-parameter)VUID-vkEnumerateDeviceLayerProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkEnumerateDeviceLayerProperties-pPropertyCount-parameter)VUID-vkEnumerateDeviceLayerProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkEnumerateDeviceLayerProperties-pProperties-parameter)VUID-vkEnumerateDeviceLayerProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkLayerProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerProperties.html) structures

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_INCOMPLETE`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkLayerProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayerProperties.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkEnumerateDeviceLayerProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0