# vkGetPhysicalDeviceCooperativeVectorPropertiesNV(3) Manual Page

## Name

vkGetPhysicalDeviceCooperativeVectorPropertiesNV - Returns properties describing what cooperative vector types are supported



## [](#_c_specification)C Specification

To enumerate the supported cooperative vector type combinations, call:

```c++
// Provided by VK_NV_cooperative_vector
VkResult vkGetPhysicalDeviceCooperativeVectorPropertiesNV(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pPropertyCount,
    VkCooperativeVectorPropertiesNV*            pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device.
- `pPropertyCount` is a pointer to an integer related to the number of cooperative vector properties available or queried.
- `pProperties` is either `NULL` or a pointer to an array of [VkCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeVectorPropertiesNV.html) structures.

## [](#_description)Description

If `pProperties` is `NULL`, then the number of cooperative vector properties available is returned in `pPropertyCount`. Otherwise, `pPropertyCount` **must** point to a variable set by the user to the number of elements in the `pProperties` array, and on return the variable is overwritten with the number of structures actually written to `pProperties`. If `pPropertyCount` is less than the number of cooperative vector properties available, at most `pPropertyCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available cooperative vector properties were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceCooperativeVectorPropertiesNV-physicalDevice-parameter)VUID-vkGetPhysicalDeviceCooperativeVectorPropertiesNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceCooperativeVectorPropertiesNV-pPropertyCount-parameter)VUID-vkGetPhysicalDeviceCooperativeVectorPropertiesNV-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceCooperativeVectorPropertiesNV-pProperties-parameter)VUID-vkGetPhysicalDeviceCooperativeVectorPropertiesNV-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeVectorPropertiesNV.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_cooperative\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_vector.html), [VkCooperativeVectorPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeVectorPropertiesNV.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceCooperativeVectorPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0