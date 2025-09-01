# vkGetPhysicalDeviceExternalTensorPropertiesARM(3) Manual Page

## Name

vkGetPhysicalDeviceExternalTensorPropertiesARM - Function for querying external tensor handle capabilities.



## [](#_c_specification)C Specification

To query the external handle types supported by tensors, call:

```c++
// Provided by VK_ARM_tensors
void vkGetPhysicalDeviceExternalTensorPropertiesARM(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalTensorInfoARM* pExternalTensorInfo,
    VkExternalTensorPropertiesARM*              pExternalTensorProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the tensor capabilities.
- `pExternalTensorInfo` is a pointer to a [VkPhysicalDeviceExternalTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalTensorInfoARM.html) structure describing the parameters that would be consumed by [vkCreateTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorARM.html).
- `pExternalTensorProperties` is a pointer to a [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html) structure in which the capabilities are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceExternalTensorPropertiesARM-physicalDevice-parameter)VUID-vkGetPhysicalDeviceExternalTensorPropertiesARM-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceExternalTensorPropertiesARM-pExternalTensorInfo-parameter)VUID-vkGetPhysicalDeviceExternalTensorPropertiesARM-pExternalTensorInfo-parameter  
  `pExternalTensorInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceExternalTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalTensorInfoARM.html) structure
- [](#VUID-vkGetPhysicalDeviceExternalTensorPropertiesARM-pExternalTensorProperties-parameter)VUID-vkGetPhysicalDeviceExternalTensorPropertiesARM-pExternalTensorProperties-parameter  
  `pExternalTensorProperties` **must** be a valid pointer to a [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html) structure

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceExternalTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalTensorInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceExternalTensorPropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0