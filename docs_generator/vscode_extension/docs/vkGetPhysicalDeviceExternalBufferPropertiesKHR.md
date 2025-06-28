# vkGetPhysicalDeviceExternalBufferProperties(3) Manual Page

## Name

vkGetPhysicalDeviceExternalBufferProperties - Query external handle types supported by buffers



## [](#_c_specification)C Specification

To query the external handle types supported by buffers, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceExternalBufferProperties(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalBufferInfo*   pExternalBufferInfo,
    VkExternalBufferProperties*                 pExternalBufferProperties);
```

or the equivalent command

```c++
// Provided by VK_KHR_external_memory_capabilities
void vkGetPhysicalDeviceExternalBufferPropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalBufferInfo*   pExternalBufferInfo,
    VkExternalBufferProperties*                 pExternalBufferProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the buffer capabilities.
- `pExternalBufferInfo` is a pointer to a [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalBufferInfo.html) structure describing the parameters that would be consumed by [vkCreateBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBuffer.html).
- `pExternalBufferProperties` is a pointer to a [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html) structure in which capabilities are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceExternalBufferProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceExternalBufferProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferInfo-parameter)VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferInfo-parameter  
  `pExternalBufferInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalBufferInfo.html) structure
- [](#VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferProperties-parameter)VUID-vkGetPhysicalDeviceExternalBufferProperties-pExternalBufferProperties-parameter  
  `pExternalBufferProperties` **must** be a valid pointer to a [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceExternalBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalBufferInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceExternalBufferProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0