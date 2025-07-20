# vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM - Query the data processing engines and the operations they support for a given queue family of a physical device



## [](#_c_specification)C Specification

To query the data graph processing engines and operations they support for a specific queue family of a physical device, call:

```c++
// Provided by VK_ARM_data_graph
VkResult vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    uint32_t*                                   pQueueFamilyDataGraphPropertyCount,
    VkQueueFamilyDataGraphPropertiesARM*        pQueueFamilyDataGraphProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device to query.
- `queueFamilyIndex` is the index of the queue family being queried.
- `pQueueFamilyDataGraphPropertyCount` is a pointer to an integer related to the number of properties available or queried.
- `pQueueFamilyDataGraphProperties` is either `NULL` or a pointer to an array of [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html) structures.

## [](#_description)Description

If `pQueueFamilyDataGraphProperties` is `NULL`, then the number of properties available is returned in `pQueueFamilyDataGraphPropertyCount`. Otherwise, `pQueueFamilyDataGraphPropertyCount` **must** point to a variable set by the application to the number of elements in the `pQueueFamilyDataGraphProperties` array, and on return the variable is overwritten with the number of structures actually written to `pQueueFamilyDataGraphProperties`. If `pQueueFamilyDataGraphPropertyCount` is less than the number of properties available, at most `pQueueFamilyDataGraphPropertyCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available properties were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM-physicalDevice-parameter)VUID-vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM-pQueueFamilyDataGraphPropertyCount-parameter)VUID-vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM-pQueueFamilyDataGraphPropertyCount-parameter  
  `pQueueFamilyDataGraphPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM-pQueueFamilyDataGraphProperties-parameter)VUID-vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM-pQueueFamilyDataGraphProperties-parameter  
  If the value referenced by `pQueueFamilyDataGraphPropertyCount` is not `0`, and `pQueueFamilyDataGraphProperties` is not `NULL`, `pQueueFamilyDataGraphProperties` **must** be a valid pointer to an array of `pQueueFamilyDataGraphPropertyCount` [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html) structures

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

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkQueueFamilyDataGraphPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyDataGraphPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyDataGraphPropertiesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0